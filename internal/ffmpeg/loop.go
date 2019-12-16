package ffmpeg

type loopedReader struct {
	r Reader
	f func() (Reader, error)
}

func (r *loopedReader) Read() (frame Frame, err error) {
	if r.r == nil {
		r.r, err = r.f()
		if err != nil {
			return Frame{}, err
		}
	}

	for {
		frame, err = r.r.Read()
		if err == nil && len(frame.Pix) > 0 {
			return
		}
		r.r, err = r.f()
		if err != nil {
			return Frame{}, err
		}
	}
}

// Loop wraps the video, returning the start of the video's frame data when it is exhausted.
func Loop(f func() (Reader, error)) Reader {
	return &loopedReader{nil, f}
}
