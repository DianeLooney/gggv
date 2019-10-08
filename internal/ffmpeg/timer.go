package ffmpeg

import "time"

type TimescaleReader interface {
	Reader
	Timescale(f float64)
}

type timer struct {
	r         Reader
	prev      Frame
	next      time.Time
	timescale float64
}

func (t *timer) Timescale(ts float64) {
	t.timescale = ts
}

func (t *timer) Read() (Frame, error) {
	if t.timescale == 0 {
		return t.prev, nil
	}

	if time.Now().Add(10 * time.Second).Before(t.next) {
		t.next = time.Now()
	}

	if t.next.After(time.Now()) {
		return t.prev, nil
	}

	f, err := t.r.Read()
	if err != nil {
		return f, err
	}
	t.prev = f
	frameDuration := time.Duration(float64(f.Duration) / t.timescale)
	t.next = t.next.Add(time.Duration(frameDuration))
	return f, nil
}

func NewTimer(r Reader) TimescaleReader {
	return &timer{
		r,
		Frame{},
		time.Now(),
		1,
	}
}
