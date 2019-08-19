package main

import (
	"github.com/dianelooney/gggv/internal/ffmpeg"
)

var fixtures = []string{
	"leto-1.mp4",
}

func main() {
	d, err := ffmpeg.NewFileDecoder("leto-1.mp4")
	if err != nil {
		panic(err)
	}
	frame := d.NextFrame()

}
