package ffmpeg_test

import (
	"testing"
	"time"

	"github.com/dianelooney/gggv/internal/ffmpeg"
)

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
