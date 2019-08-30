package logs

import (
	"fmt"
	"log"
)

type logger = func(s ...interface{})

var Log, Debug, Warn, Error, Fatal logger

func DefaultLogger(s ...interface{}) {
	fmt.Println(s...)
}

func init() {
	Log = DefaultLogger
	Warn = DefaultLogger
	Debug = DefaultLogger
	Error = DefaultLogger
	Fatal = func(s ...interface{}) {
		log.Fatalln(s...)
	}
}
