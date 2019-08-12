package gvdl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"

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
	case "reload":
		return execReload(sc, d)
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
	case "layer":
		return execAddLayer(sc, d)
	case "program":
		return execAddProgram(sc, d)
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

	d.AddSource(name, path)

	return
}

func execAddProgram(sc *bufio.Scanner, d *daemon.D) (err error) {
	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	name := sc.Text()

	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	pathV := sc.Text()

	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	pathF := sc.Text()

	d.AddProgram(name, pathV, pathF)
	return
}

func execAddLayer(sc *bufio.Scanner, d *daemon.D) (err error) {
	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	name := sc.Text()

	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	source := sc.Text()

	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	program := sc.Text()

	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	depthA := sc.Text()
	depthI, err := strconv.ParseFloat(depthA, 32)
	if err != nil {
		return err
	}
	d.AddLayer(name, source, program, float32(depthI))
	return
}

func execReload(sc *bufio.Scanner, d *daemon.D) (err error) {
	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	name := sc.Text()

	switch name {
	case "programs":
		d.ReloadPrograms()
		return
	}

	return fmt.Errorf("unrecognized reload subcommand '%s'", name)
}
