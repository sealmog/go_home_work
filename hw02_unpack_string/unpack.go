package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	ru := []rune(s)
	if IsDigit(ru[0]) {
		return "", ErrInvalidString
	}

	var prev int32
	var b strings.Builder

	for _, r := range s {
		if IsDigit(r) {
			if IsDigit(prev) {
				return "", ErrInvalidString
			}

			if int(r-'0') == 0 {
				str := b.String()
				str = str[:len(str)-len(string(prev))]
				b.Reset()
				b.WriteString(str)
				continue
			}
			str := strings.Repeat(string(prev), int(r-'0')-1)
			b.WriteString(str)
		} else {
			b.WriteRune(r)
		}
		prev = r
	}
	return b.String(), nil
}

func IsDigit(r rune) bool {
	if r < '0' || r > '9' {
		return false
	}
	return true
}
