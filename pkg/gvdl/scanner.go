package gvdl

import (
	"errors"
	"io"
	"regexp"
	"strconv"
)

var regWhitespace = regexp.MustCompile(`^(\s*)`)
var regComment = regexp.MustCompile(`^#([^$]*)$`)
var regStringSng = regexp.MustCompile(`^('(?:[^\]|\[^'])*')`)
var regStringDbl = regexp.MustCompile(`^("(?:[^\]|\[^"])*")`)
var regWord = regexp.MustCompile(`^([a-zA-Z_][a-zA-Z_0-9]*)`)
var regNum = regexp.MustCompile(`^([-+]?[0-9]*\.?[0-9]+)`)
var regSemi = regexp.MustCompile(`^(;)`)

func splitRegexp(data []byte, r *regexp.Regexp) (advance int, token []byte, err error) {
	matches := r.FindSubmatch(data)
	if len(matches) == 0 {
		return 0, nil, errors.New("Scan error")
	}

	advance = len(matches[0])
	token = matches[1]

	return
}

func SplitFn(data []byte, atEOF bool) (advance int, token []byte, err error) {
	whitespaceSkipped, _, _ := splitRegexp(data, regWhitespace)
	defer func() {
		advance += whitespaceSkipped
	}()

	data = data[whitespaceSkipped:]

	if len(data) == 0 {
		return 0, nil, io.EOF
	}

	switch data[0] {
	case '#':
		return splitRegexp(data, regComment)
	case '"':
		advance, token, err = splitRegexp(data, regStringDbl)
		if err != nil {
			return advance, nil, err
		}
		var str string
		str, err = strconv.Unquote(string(token))
		token = []byte(str)
		return
	case ';':
		return splitRegexp(data, regSemi)
	}

	if data[0] == '_' || (data[0] >= 'a' && data[0] <= 'z') || (data[0] >= 'A' && data[0] <= 'Z') {
		return splitRegexp(data, regWord)
	}

	if (data[0] >= '0' && data[0] <= '9') || data[0] == '+' || data[0] == '-' {
		return splitRegexp(data, regNum)
	}

	return 0, nil, nil
}
