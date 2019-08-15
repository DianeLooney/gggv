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
	case "set":
		return execSet(sc, d)
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
		return execAddSourceFFVideo(sc, d)
	case "program":
		return execAddProgram(sc, d)
	}
	return fmt.Errorf("unrecognized add subcommand '%s'", sc.Text())
}

func execAddSourceFFVideo(sc *bufio.Scanner, d *daemon.D) (err error) {
	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	name := sc.Text()

	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	path := sc.Text()

	d.AddSourceFFVideo(name, path)

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

func execSet(sc *bufio.Scanner, d *daemon.D) (err error) {
	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	switch sc.Text() {
	case "uniform":
		return execSetUniform(sc, d)
	}
	return fmt.Errorf("unrecognized set subcommand '%s'", sc.Text())
}
func execSetUniform(sc *bufio.Scanner, d *daemon.D) (err error) {
	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	typ := sc.Text()

	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	name := sc.Text()

	if !sc.Scan() {
		return io.ErrUnexpectedEOF
	}
	vRaw := sc.Text()
	var value interface{}
	switch typ {
	case "float":
		floatV, err := strconv.ParseFloat(vRaw, 32)
		if err != nil {
			return err
		}
		value = float32(floatV)
	}

	var layers []string
	for sc.Scan() {
		layers = append(layers, sc.Text())
	}

	d.SetUniform(name, typ, value, layers)
	fmt.Println(name, typ, value, layers)
	return nil
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
