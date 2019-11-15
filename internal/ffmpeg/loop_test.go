package ffmpeg_test

import (
	"errors"
	"testing"

	"github.com/dianelooney/gggv/internal/ffmpeg"
)

func TestLoopedReader(t *testing.T) {
	t.Run("It gobbles errors", func(t *testing.T) {
		i := 0
		f := func() (ffmpeg.Reader, error) {
			i++

			if i%2 == 0 {
				return nil, errors.New("whoops")
			} else {
				return &mockFaulty{}, nil
			}
		}
		l := ffmpeg.Loop(f)
		for i := 0; i < 100; i++ {
			_, err := l.Read()
			if err != nil {
				t.Fail()
			}
		}
	})
}
