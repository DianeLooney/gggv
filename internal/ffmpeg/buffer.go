package ffmpeg

type bufferedReader struct {
	reader Reader
	buffer chan interface{}
}

func (b *bufferedReader) Read() (Frame, error) {
	f := <-b.buffer

	switch v := f.(type) {
	case Frame:
		return v, nil
	case error:
		return Frame{}, v
	}

	panic("Should not occur")
}

// Buffer returns a Reader wrapping r that performs decodes ahead of time to smooth out framerates
func Buffer(r Reader) Reader {
	b := &bufferedReader{
		r,
		make(chan interface{}, 20),
	}

	go func() {
		for {
			frame, err := b.reader.Read()
			if err != nil {
				b.buffer <- err
			} else {
				b.buffer <- frame
			}
		}
	}()

	return b
}
