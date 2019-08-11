package fps

import "time"

var sec = time.Now()
var lastSec int
var count int

func Next() {
	t := time.Now()

	if sec.Add(time.Second).Before(t) {
		lastSec = count
		count = 0
		for sec.Add(time.Second).Before(t) {
			sec = sec.Add(time.Second)
		}
	}

	instantaneous = t.Sub(prevFrame)
	prevFrame = t

	count++
}

func LastSec() int {
	return lastSec
}

var prevFrame = time.Now()
var instantaneous time.Duration

func FrameDuration() time.Duration {
	return instantaneous
}
