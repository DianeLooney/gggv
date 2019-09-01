package ffmpeg

import (
	"flag"
	"time"
	"unsafe"

	"github.com/dianelooney/gggv/internal/errors"
	"github.com/dianelooney/gggv/internal/logs"
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"github.com/giorgisio/goav/swscale"
)

var fflogformat = flag.Bool("fflogformat", false, "Run AvDumpFormat when decoding a video file")

type Decoder struct {
	buffer         unsafe.Pointer
	pCodecCtxOrig  *avformat.CodecContext
	pCodecCtx      *avcodec.Context
	pFormatContext *avformat.Context
	pFrame         *avutil.Frame
	pFrameRGB      *avutil.Frame
	packet         *avcodec.Packet
	videoStreamNum int
	swsCtx         *swscale.Context

	width  int
	height int
}

func (d *Decoder) Dimensions() (width, height int) {
	return d.width, d.height
}

func (d *Decoder) Begin(fname string) error {
	d.pFormatContext = avformat.AvformatAllocContext()
	if code := avformat.AvformatOpenInput(&d.pFormatContext, fname, nil, nil); code != 0 {
		return errors.FFDecoderOpenInput(fname, avutil.ErrorFromCode(code))
	}

	if code := d.pFormatContext.AvformatFindStreamInfo(nil); code < 0 {
		return errors.FFDecoderStreamInfo(fname, avutil.ErrorFromCode(code))
	}

	if *fflogformat {
		d.pFormatContext.AvDumpFormat(0, fname, 0)
	}

	for i := 0; i < int(d.pFormatContext.NbStreams()); i++ {
		switch d.pFormatContext.Streams()[i].CodecParameters().AvCodecGetType() {
		case avformat.AVMEDIA_TYPE_VIDEO:
			d.pCodecCtxOrig = d.pFormatContext.Streams()[i].Codec()

			codecID := d.pCodecCtxOrig.GetCodecId()
			pCodec := avcodec.AvcodecFindDecoder(avcodec.CodecId(codecID))
			if pCodec == nil {
				return errors.FFDecoderUnsupportedCodec(fname, codecID)
			}

			d.pCodecCtx = pCodec.AvcodecAllocContext3()
			if code := d.pCodecCtx.AvcodecCopyContext((*avcodec.Context)(unsafe.Pointer(d.pCodecCtxOrig))); code != 0 {
				return errors.FFDecoderCopyCodecCtx(fname, avutil.ErrorFromCode(code))
			}

			if code := d.pCodecCtx.AvcodecOpen2(pCodec, nil); code < 0 {
				return errors.FFDecoderOpenCodec(fname, avutil.ErrorFromCode(code))
			}

			if d.pFrame = avutil.AvFrameAlloc(); d.pFrame == nil {
				return errors.FFDecoderCopyCodecCtx(fname)
			}

			if d.pFrameRGB = avutil.AvFrameAlloc(); d.pFrameRGB == nil {
				return errors.FFDecoderCopyCodecCtx(fname)
			}

			numBytes := uintptr(avcodec.AvpictureGetSize(
				avcodec.AV_PIX_FMT_RGB24,
				d.pCodecCtx.Width(),
				d.pCodecCtx.Height(),
			))
			d.buffer = avutil.AvMalloc(numBytes)

			avp := (*avcodec.Picture)(unsafe.Pointer(d.pFrameRGB))
			avp.AvpictureFill((*uint8)(d.buffer), avcodec.AV_PIX_FMT_RGB24, d.pCodecCtx.Width(), d.pCodecCtx.Height())

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

			d.videoStreamNum = i
			d.packet = avcodec.AvPacketAlloc()

			return nil
		}
	}
	return errors.FFDecoderMissingStream(fname)
}

func (d *Decoder) readFrame() (ok bool) {
	if d.packet.StreamIndex() != d.videoStreamNum {
		return false
	}

	// Decode video frame
	response := d.pCodecCtx.AvcodecSendPacket(d.packet)
	if response < 0 {
		logs.Error("Error while sending a packet to the decoder", avutil.ErrorFromCode(response))
	}
	for response >= 0 {
		response = d.pCodecCtx.AvcodecReceiveFrame((*avcodec.Frame)(unsafe.Pointer(d.pFrame)))
		if response == avutil.AvErrorEAGAIN || response == avutil.AvErrorEOF {
			return false
		} else if response < 0 {
			logs.Error("Error while receiving a frame from the decoder", avutil.ErrorFromCode(response))
			return false
		}

		// Convert the image from its native format to RGB
		swscale.SwsScale2(
			d.swsCtx,
			avutil.Data(d.pFrame),
			avutil.Linesize(d.pFrame),
			0,
			d.pCodecCtx.Height(),
			avutil.Data(d.pFrameRGB),
			avutil.Linesize(d.pFrameRGB),
		)
		return true
	}

	return false
}

func (d *Decoder) frameDuration() time.Duration {
	r := d.pCodecCtx.AvCodecGetPktTimebase()
	return (time.Duration)(r.Num()) * time.Second / time.Duration(r.Den())
}

func (d *Decoder) NextFrame() (rgb []uint8, duration int) {
	for d.pFormatContext.AvReadFrame(d.packet) >= 0 {
		ok := d.readFrame()
		if !ok {
			continue
		}
		width, height := d.Dimensions()
		offset := uintptr(unsafe.Pointer(avutil.Data(d.pFrameRGB)[0]))
		linesize := uintptr(avutil.Linesize(d.pFrameRGB)[0])
		for y := 0; y < height; y++ {
			buf := make([]uint8, width*3)
			for i := 0; i < width*3; i++ {
				ptr := offset + uintptr(i)
				buf[i] = *(*uint8)(unsafe.Pointer(ptr))
			}
			rgb = append(rgb, buf...)

			offset += linesize
		}

		if len(rgb) == 0 {
			continue
		}
		duration = d.packet.Duration()
		d.packet.AvFreePacket()
		return
	}
	return
}

func (d *Decoder) Dealloc() {
	d.packet.AvFreePacket()
	// Free the RGB image
	avutil.AvFree(d.buffer)
	avutil.AvFrameFree(d.pFrameRGB)

	// Free the YUV frame
	avutil.AvFrameFree(d.pFrame)

	// Close the codecs
	d.pCodecCtx.AvcodecClose()
	(*avcodec.Context)(unsafe.Pointer(d.pCodecCtxOrig)).AvcodecClose()

	// Close the video file
	d.pFormatContext.AvformatCloseInput()
}

type frame struct {
	pix      []uint8
	duration int
}
