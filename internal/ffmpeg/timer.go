package ffmpeg

import "time"

type timer struct {
	r    Reader
	prev Frame
	next time.Time
}

func (t *timer) Read() (Frame, error) {
	if t.next.After(time.Now()) {
		return t.prev, nil
	}

	f, err := t.r.Read()
	if err != nil {
		return f, err
	}
	t.prev = f
	t.next = t.next.Add(f.Duration)
	return f, nil
}

func NewTimer(r Reader) Reader {
	return &timer{r, Frame{}, time.Now()}
}
