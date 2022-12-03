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

	if IsDigit(s[0:1]) {
		return "", ErrInvalidString
	}

	var prev int32
	var b strings.Builder

	for _, r := range s {
		if IsDigit(string(r)) {
			if IsDigit(string(prev)) {
				return "", ErrInvalidString
			}

			if int(r-'0') == 0 {
				str := b.String()
				str = str[:len(str)-len(string(prev))]
				b.Reset()
				b.WriteString(str)
			}

			for i := int(r - '0'); i > 1; i-- {
				b.WriteRune(prev)
			}
		} else {
			b.WriteRune(r)
		}
		prev = r
	}
	return b.String(), nil
}

func IsDigit(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
