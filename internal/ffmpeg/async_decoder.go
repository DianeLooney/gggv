package ffmpeg

import (
	"time"
)

func NewFileSampler(fname string) (a *Sampler, err error) {
	a = &Sampler{
		d:         &Decoder{},
		nextFrame: make(chan frame, 20),
		done:      make(chan bool),
	}
	err = a.d.Begin(fname)
	if err != nil {
		return nil, err
	}

	a.t = time.Now()
	a.d.width = a.d.pCodecCtx.Width()
	a.d.height = a.d.pCodecCtx.Height()

	go func() {
		defer a.d.Dealloc()

		for {
			rgb, duration := a.d.NextFrame()

			if rgb == nil {
				a.d.pCodecCtx.AvcodecFlushBuffers()
				a.d.Begin(fname)
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

type Sampler struct {
	d         *Decoder
	t         time.Time
	nextFrame chan frame
	done      chan bool
}

func (a *Sampler) Done() {
	a.done <- true
	close(a.nextFrame)
}

func (a *Sampler) Ready() bool {
	return a.t.Before(time.Now())
}

func (a *Sampler) SkipSample() {
	a.t = time.Now()
}

func (a *Sampler) Sample() (width, height int, pix []uint8) {
	f := <-a.nextFrame
	duration := time.Duration(f.duration) * a.d.frameDuration()
	a.t = a.t.Add(duration)

	width, height = a.d.Dimensions()
	pix = f.pix

	return
}
