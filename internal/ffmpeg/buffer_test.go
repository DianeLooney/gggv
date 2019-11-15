package ffmpeg_test

import (
	"errors"
	"testing"
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

func TestBufferedReader(t *testing.T) {
	t.Run("It waits for the underlying reader", func(t *testing.T) {
		m := mockSlow{}
		b := ffmpeg.Buffer(&m)

		fail := false
		go func() {
			b.Read()
			fail = true
		}()
		time.Sleep(time.Second / 10)
		if fail {
			t.Fail()
		}
	})

	t.Run("It allows the underlying reader to queue up data", func(t *testing.T) {
		m := mockFast{}
		ffmpeg.Buffer(&m)
		time.Sleep(time.Second / 10)
		if m.called < 10 {
			t.Fail()
		}
	})

	t.Run("It handles frames", func(t *testing.T) {
		m := mockFast{}
		b := ffmpeg.Buffer(&m)
		_, err := b.Read()
		if err != nil {
			t.Fail()
		}
	})

	t.Run("It handles errors", func(t *testing.T) {
		m := mockErr{}
		b := ffmpeg.Buffer(m)
		_, err := b.Read()
		if err == nil {
			t.Fail()
		}
	})
}
