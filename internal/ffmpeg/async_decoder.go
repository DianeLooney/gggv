package ffmpeg

import "time"

func NewAsyncFileDecoder(fname string) (a *AsyncDecoder, err error) {
	a = &AsyncDecoder{
		d:         &Decoder{},
		nextFrame: make(chan frame, 20),
		done:      make(chan bool),
	}
	err = a.Begin(fname)
	if err != nil {
		return nil, err
	}

	a.d.width = a.d.pCodecCtx.Width()
	a.d.height = a.d.pCodecCtx.Height()

	go func() {
		defer a.d.Dealloc()

		for {
			rgb, duration := a.d.NextFrame()

			if rgb == nil {
				a.d.pCodecCtx.AvcodecFlushBuffers()
				a.Begin(fname)
				continue
			}

			select {
			case a.nextFrame <- frame{rgb, duration}:
			case <-a.done:
				return
			}
		}
	}()

	return a, nil
}

type AsyncDecoder struct {
	d         *Decoder
	nextFrame chan frame
	done      chan bool
}

func (a *AsyncDecoder) Dimensions() (width, height int) {
	return a.d.Dimensions()
}

func (a *AsyncDecoder) Begin(fname string) (err error) {
	return a.d.Begin(fname)
}

func (a *AsyncDecoder) NextFrame() (rgb []uint8, duration time.Duration) {
	f := <-a.nextFrame
	return f.pix, time.Duration(f.duration) * a.d.frameDuration()
}

func (a *AsyncDecoder) Dealloc() {
	a.done <- true
	close(a.nextFrame)
}
