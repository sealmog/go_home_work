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
	lenRu := len(ru)
	var b strings.Builder

	if IsDigit(ru[0]) {
		return "", ErrInvalidString
	}

	for i, r := range ru {
		if i+1 < lenRu {
			if ru[i+1] == 48 {
				continue
			}
		}
		if IsDigit(r) {
			if IsDigit(ru[i-1]) {
				return "", ErrInvalidString
			}
			if ru[i] == 48 {
				continue
			}

			str := strings.Repeat(string(ru[i-1]), int(r-'0')-1)
			b.WriteString(str)
			continue
		}
		b.WriteRune(r)
	}
	return b.String(), nil
}

func IsDigit(r rune) bool {
	if r < '0' || r > '9' {
		return false
	}
	return true
}
