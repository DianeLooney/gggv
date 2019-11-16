package ffmpeg_test

import (
	"errors"
	"testing"
	"time"

	"github.com/dianelooney/gggv/internal/ffmpeg"
)

type quickFailReader struct{}

func (p *quickFailReader) Read() (ffmpeg.Frame, error) {
	return ffmpeg.Frame{Duration: time.Millisecond}, errors.New("whoops")
}

type panicReader struct {
	read bool
}

func (p *panicReader) Read() (ffmpeg.Frame, error) {
	if p.read {
		panic("already read from")
	}
	p.read = true
	return ffmpeg.Frame{Duration: time.Second}, nil
}

type quickPanic struct {
	read bool
}

func (p *quickPanic) Read() (ffmpeg.Frame, error) {
	if p.read {
		return ffmpeg.Frame{}, errors.New("whoops")
	}
	p.read = true
	return ffmpeg.Frame{Duration: time.Millisecond}, nil
}

func TestTimer(t *testing.T) {
	t.Run("It caches the return value", func(t *testing.T) {
		r := ffmpeg.NewTimer(&panicReader{})
		r.Read()
		for i := 0; i < 100; i++ {
			r.Read()
		}
	})
	t.Run("It re-reads quickly on errors", func(t *testing.T) {
		r := ffmpeg.NewTimer(&quickFailReader{})
		r.Read()
		_, err := r.Read()
		if err != nil {
			t.Fail()
		}
	})
	t.Run("Setting timescale 0 pauses readback", func(t *testing.T) {
		r := ffmpeg.NewTimer(&quickFailReader{})
		r.Timescale(0)
		r.Read()
		time.Sleep(20 * time.Millisecond)
		if _, err := r.Read(); err != nil {
			t.Fail()
		}
	})
}
