package ffmpeg

import (
	"fmt"
)

type bufferedReader struct {
	reader Reader
	buffer chan Frame
}

func (b *bufferedReader) Read() (Frame, error) {
	f := <-b.buffer
	return f, nil
}

// Buffer returns a Reader wrapping r that performs decodes ahead of time to smooth out framerates
func Buffer(r Reader) Reader {
	b := &bufferedReader{
		r,
		make(chan Frame, 20),
	}

	go func() {
		for {
			frame, err := b.reader.Read()
			if err != nil {
				fmt.Println(err)
				continue
			}
			b.buffer <- frame
		}
	}()

	return b
}
