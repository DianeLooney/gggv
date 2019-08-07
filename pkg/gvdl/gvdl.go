package gvdl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/dianelooney/gvd/pkg/daemon"
)

func Exec(data []byte, d *daemon.D) (err error) {
	sc := bufio.NewScanner(bytes.NewBuffer(data))
	sc.Split(SplitFn)
	return exec(sc, d)
}

func exec(sc *bufio.Scanner, d *daemon.D) (err error) {
	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	switch sc.Text() {
	case "add":
		return execAdd(sc, d)
	}
	return fmt.Errorf("unrecognized command '%s'", sc.Text())
}

func execAdd(sc *bufio.Scanner, d *daemon.D) (err error) {
	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	switch sc.Text() {
	case "source":
		return execAddSource(sc, d)
	}
	return fmt.Errorf("unrecognized add subcommand '%s'", sc.Text())
}

func execAddSource(sc *bufio.Scanner, d *daemon.D) (err error) {
	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	name := sc.Text()
	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	path := sc.Text()
	return d.AddSource(name, path)
}
