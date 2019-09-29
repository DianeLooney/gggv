package ffmpeg

import (
	"flag"
	"fmt"
	"time"
	"unsafe"

	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"github.com/giorgisio/goav/swscale"
)

var fflogformat = flag.Bool("fflogformat", false, "Run AvDumpFormat when decoding a video file")

type fileDecoder struct {
	width  int
	height int

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

func (d *fileDecoder) Size() (width, height int) {
	return d.width, d.height
}

func (d *fileDecoder) Read() (frame Frame, err error) {
	width, height := d.Size()
	frame.duration = d.frameDuration()

	for d.pFormatContext.AvReadFrame(d.packet) >= 0 {
		if d.packet.StreamIndex() != d.videoStreamNum {
			continue
		}
		response := d.pCodecCtx.AvcodecSendPacket(d.packet)
		if response < 0 {
			continue
		}
		response = d.pCodecCtx.AvcodecReceiveFrame((*avcodec.Frame)(unsafe.Pointer(d.pFrame)))
		if response == avutil.AvErrorEAGAIN {
			continue
		}
		if response == avutil.AvErrorEOF {
			err = fmt.Errorf("EOF reached")
			return
		}
		if response != 0 {
			err = fmt.Errorf("Error while receiving a frame from the decoder", avutil.ErrorFromCode(response))
			return
		}
		swscale.SwsScale2(
			d.swsCtx,
			avutil.Data(d.pFrame),
			avutil.Linesize(d.pFrame),
			0,
			d.pCodecCtx.Height(),
			avutil.Data(d.pFrameRGB),
			avutil.Linesize(d.pFrameRGB),
		)

		offset := uintptr(unsafe.Pointer(avutil.Data(d.pFrameRGB)[0]))
		linesize := uintptr(avutil.Linesize(d.pFrameRGB)[0])
		rgb := make([]uint8, width*height*3)

		for y := 0; y < height; y++ {
			for i := 0; i < width*3; i++ {
				ptr := offset + uintptr(i)
				rgb[(height-1-y)*3*width+i] = *(*uint8)(unsafe.Pointer(ptr))
			}
			offset += linesize
		}

		frame.pix = rgb

		d.packet.AvFreePacket()
		return
	}
	return
}

func (d *fileDecoder) frameDuration() time.Duration {
	r := d.pCodecCtx.AvCodecGetPktTimebase()
	return (time.Duration)(r.Num()) * time.Second / time.Duration(r.Den())
}

func (d *fileDecoder) Dealloc() {
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

type Frame struct {
	pix      []uint8
	duration time.Duration
}
