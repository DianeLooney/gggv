package ffmpeg

import (
	"time"
)

func NewFileSampler(fname string) (a *Sampler, err error) {
	a = &Sampler{
		nextFrame: make(chan Frame, 20),
		done:      make(chan bool),
	}
	r, err := NewReadSizer(fname)
	if err != nil {
		return nil, err
	}
	a.d = r

	a.t = time.Now()

	go func() {
		for {
			a.d.Read()
			frame, err := a.d.Read()

			if err != nil {
				a.d, _ = NewReadSizer(fname)
				continue
			}

			select {
			case a.nextFrame <- frame:
			case <-a.done:
				return
			}
		}
	}()

	return a, nil
}

type Sampler struct {
	d         ReadSizer
	t         time.Time
	nextFrame chan Frame
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
	duration := f.duration
	a.t = a.t.Add(duration)

	width, height = a.d.Size()
	pix = f.pix

	return
}
