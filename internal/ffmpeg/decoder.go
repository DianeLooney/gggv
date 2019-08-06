package ffmpeg

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"github.com/giorgisio/goav/swscale"
)

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
}

func (d *Decoder) Dimensions() (width, height int) {
	return d.pCodecCtx.Width(), d.pCodecCtx.Height()
}

func (d *Decoder) Begin(fname string) {

	// Open video file
	d.pFormatContext = avformat.AvformatAllocContext()
	if avformat.AvformatOpenInput(&d.pFormatContext, fname, nil, nil) != 0 {
		fmt.Printf("Unable to open file %s\n", os.Args[1])
		os.Exit(1)
	}

	// Retrieve stream information
	if d.pFormatContext.AvformatFindStreamInfo(nil) < 0 {
		fmt.Println("Couldn't find stream information")
		os.Exit(1)
	}

	// Dump information about file onto standard error
	d.pFormatContext.AvDumpFormat(0, fname, 0)

	// Find the first video stream
	for i := 0; i < int(d.pFormatContext.NbStreams()); i++ {
		switch d.pFormatContext.Streams()[i].CodecParameters().AvCodecGetType() {
		case avformat.AVMEDIA_TYPE_VIDEO:

			// Get a pointer to the codec context for the video stream
			d.pCodecCtxOrig = d.pFormatContext.Streams()[i].Codec()
			// Find the decoder for the video stream
			pCodec := avcodec.AvcodecFindDecoder(avcodec.CodecId(d.pCodecCtxOrig.GetCodecId()))
			if pCodec == nil {
				fmt.Println("Unsupported codec!")
				os.Exit(1)
			}
			// Copy context
			d.pCodecCtx = pCodec.AvcodecAllocContext3()
			if d.pCodecCtx.AvcodecCopyContext((*avcodec.Context)(unsafe.Pointer(d.pCodecCtxOrig))) != 0 {
				fmt.Println("Couldn't copy codec context")
				os.Exit(1)
			}

			// Open codec
			if d.pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
				fmt.Println("Could not open codec")
				os.Exit(1)
			}

			// Allocate video frame
			d.pFrame = avutil.AvFrameAlloc()

			// Allocate an AVFrame structure
			if d.pFrameRGB = avutil.AvFrameAlloc(); d.pFrameRGB == nil {
				fmt.Println("Unable to allocate RGB Frame")
				os.Exit(1)
			}

			// Determine required buffer size and allocate buffer
			numBytes := uintptr(avcodec.AvpictureGetSize(avcodec.AV_PIX_FMT_RGB24, d.pCodecCtx.Width(),
				d.pCodecCtx.Height()))
			d.buffer = avutil.AvMalloc(numBytes)

			// Assign appropriate parts of buffer to image planes in pFrameRGB
			// Note that pFrameRGB is an AVFrame, but AVFrame is a superset
			// of AVPicture
			avp := (*avcodec.Picture)(unsafe.Pointer(d.pFrameRGB))
			avp.AvpictureFill((*uint8)(d.buffer), avcodec.AV_PIX_FMT_RGB24, d.pCodecCtx.Width(), d.pCodecCtx.Height())

			// initialize SWS context for software scaling
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

			// Read frames and save first five frames to disk
			d.videoStreamNum = i
			d.packet = avcodec.AvPacketAlloc()

			return
		default:
			fmt.Println("Didn't find a video stream")
			os.Exit(1)
		}
	}
}

func (d *Decoder) readFrame() (ok bool) {
	if d.packet.StreamIndex() != d.videoStreamNum {
		return false
	}

	// Decode video frame
	response := d.pCodecCtx.AvcodecSendPacket(d.packet)
	if response < 0 {
		fmt.Printf("Error while sending a packet to the decoder: %s\n", avutil.ErrorFromCode(response))
	}
	for response >= 0 {
		response = d.pCodecCtx.AvcodecReceiveFrame((*avcodec.Frame)(unsafe.Pointer(d.pFrame)))
		if response == avutil.AvErrorEAGAIN || response == avutil.AvErrorEOF {
			return false
		} else if response < 0 {
			fmt.Printf("Error while receiving a frame from the decoder: %s\n", avutil.ErrorFromCode(response))
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

func (d *Decoder) NextFrame() {
	for d.pFormatContext.AvReadFrame(d.packet) >= 0 {
		// Is this a packet from the video stream?
		ok := d.readFrame()
		if ok {
			return
		}

		// Free the packet that was allocated by av_read_frame
		d.packet.AvFreePacket()
	}
}

func (d *Decoder) Dealloc() {
	fmt.Println("Finalizing decoder")
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

func NewFileDecoder(fname string) (d *Decoder, frame *avutil.Frame) {
	d = &Decoder{}
	d.Begin(fname)

	return d, d.pFrameRGB
}
