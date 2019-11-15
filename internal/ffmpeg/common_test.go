package ffmpeg_test

import (
	"errors"
	"time"

	"github.com/dianelooney/gggv/internal/ffmpeg"
)

type mockSlow struct {
}

func (r *mockSlow) Read() (ffmpeg.Frame, error) {
	time.Sleep(time.Second)
	return ffmpeg.Frame{}, nil
}

type mockFast struct {
	called int
}

func (r *mockFast) Read() (ffmpeg.Frame, error) {
	r.called++
	return ffmpeg.Frame{}, nil
}

type mockErr struct{}

func (r mockErr) Read() (ffmpeg.Frame, error) {
	return ffmpeg.Frame{}, errors.New("something went wrong")
}

type mockFaulty struct {
	i int
}

func (m *mockFaulty) Read() (ffmpeg.Frame, error) {
	m.i++
	if m.i%2 == 0 {
		return ffmpeg.Frame{}, nil
	} else {
		return ffmpeg.Frame{}, errors.New("something went wrong")
	}
}
