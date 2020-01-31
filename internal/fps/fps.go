package fps

import (
	"flag"
	"fmt"
	"time"
)

var sec = time.Now()
var lastSec int
var count int

var fps = flag.Bool("fps", false, "Log fps to the console.")

func init() {
	go func() {
		for range time.NewTicker(time.Second).C {
			if *fps {
				fmt.Printf("[fps]\t%v\t%s\n", lastSec, instantaneous)
			}
		}
	}()
}

func DrawStart() {
	start = time.Now()
}

func DrawDone() {
	t := time.Now()

	if sec.Add(time.Second).Before(t) {
		lastSec = count
		count = 0
		for sec.Add(time.Second).Before(t) {
			sec = sec.Add(time.Second)
		}
	}

	instantaneous = t.Sub(start)

	count++
}

func LastSec() int {
	return lastSec
}

var start = time.Now()
var instantaneous time.Duration

func FrameDuration() time.Duration {
	return instantaneous
}
