package ffmpeg

import (
	"runtime"
	"time"
	"unsafe"

	"github.com/dianelooney/gggv/internal/errors"
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"github.com/giorgisio/goav/swscale"
)

// A Reader reads a frame, and can give the size of the frame
type Reader interface {
	Read() (frame Frame, err error)
}

// NewReader returns a new ReadSizer from the specified source
//
// source should be recognizable by ffmpeg's avformat_open_input
func NewReader(source string) (r Reader, err error) {
	d := reader{}
	d.pFormatContext = avformat.AvformatAllocContext()
	if code := avformat.AvformatOpenInput(&d.pFormatContext, source, nil, nil); code != 0 {
		return nil, errors.FFDecoderOpenInput(source, avutil.ErrorFromCode(code))
	}

	if code := d.pFormatContext.AvformatFindStreamInfo(nil); code < 0 {
		return nil, errors.FFDecoderStreamInfo(source, avutil.ErrorFromCode(code))
	}

	if *fflogformat {
		d.pFormatContext.AvDumpFormat(0, source, 0)
	}

	var stream *avformat.Stream
	for i, s := range d.pFormatContext.Streams() {
		if s.CodecParameters().AvCodecGetType() == avformat.AVMEDIA_TYPE_VIDEO {
			d.videoStreamNum = i
			stream = s
			break
		}
	}
	if stream == nil {
		return nil, errors.FFDecoderMissingStream(source)
	}
	d.pCodecCtxOrig = stream.Codec()

	codecID := d.pCodecCtxOrig.GetCodecId()
	pCodec := avcodec.AvcodecFindDecoder(avcodec.CodecId(codecID))
	if pCodec == nil {
		return nil, errors.FFDecoderUnsupportedCodec(source, codecID)
	}

	d.pCodecCtx = pCodec.AvcodecAllocContext3()
	if code := d.pCodecCtx.AvcodecCopyContext((*avcodec.Context)(unsafe.Pointer(d.pCodecCtxOrig))); code != 0 {
		return nil, errors.FFDecoderCopyCodecCtx(source, avutil.ErrorFromCode(code))
	}

	if code := d.pCodecCtx.AvcodecOpen2(pCodec, nil); code < 0 {
		return nil, errors.FFDecoderOpenCodec(source, avutil.ErrorFromCode(code))
	}

	if d.pFrame = avutil.AvFrameAlloc(); d.pFrame == nil {
		return nil, errors.FFDecoderCopyCodecCtx(source)
	}

	if d.pFrameRGB = avutil.AvFrameAlloc(); d.pFrameRGB == nil {
		return nil, errors.FFDecoderCopyCodecCtx(source)
	}

	numBytes := uintptr(avcodec.AvpictureGetSize(
		avcodec.AV_PIX_FMT_RGB24,
		d.pCodecCtx.Width(),
		d.pCodecCtx.Height(),
	))
	d.buffer = avutil.AvMalloc(numBytes)

	avp := (*avcodec.Picture)(unsafe.Pointer(d.pFrameRGB))
	avp.AvpictureFill((*uint8)(d.buffer), avcodec.AV_PIX_FMT_RGB24, d.pCodecCtx.Width(), d.pCodecCtx.Height())

	// To Do: When given an image with 0 width and height this fails?
	d.swsCtx = swscale.SwsGetcontext(
		d.pCodecCtx.Width(),
		d.pCodecCtx.Height(),
		(swscale.PixelFormat)(d.pCodecCtx.PixFmt()),
		d.pCodecCtx.Width(),
		d.pCodecCtx.Height(),
		avcodec.AV_PIX_FMT_RGB24,
		avcodec.SWS_BILINEAR,
		nil,
		nil,
		nil,
	)
	d.width = d.pCodecCtx.Width()
	d.height = d.pCodecCtx.Height()

	d.packet = avcodec.AvPacketAlloc()

	r = &d
	runtime.SetFinalizer(r, (*reader).Dealloc)

	return
}

// Frame contains the data regarding a frame of video
type Frame struct {
	Pix      []uint8
	Duration time.Duration
	Width    int
	Height   int
}
