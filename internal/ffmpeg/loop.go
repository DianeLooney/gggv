package ffmpeg

type loopedReader struct {
	r Reader
	f func() (Reader, error)
}

func (r *loopedReader) Read() (frame Frame, err error) {
	for {
		frame, err = r.read()
		if err == nil {
			return
		}
	}
}

func (r *loopedReader) read() (frame Frame, err error) {
	if r.r == nil {
		r.r, err = r.f()
		return
	}

	frame, err = r.r.Read()
	if err != nil {
		r.r = nil
	}
	return
}

// Loop wraps the video, returning the start of the video's frame data when it is exhausted.
func Loop(f func() (Reader, error)) Reader {
	return &loopedReader{nil, f}
}
