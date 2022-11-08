package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	if _, err := strconv.Atoi(s[0:1]); err == nil {
		return "", ErrInvalidString
	}

	var result []string

	//nolint:nestif
	for pos, char := range s {
		if unicode.IsDigit(char) {
			prev := s[pos-1 : pos]
			if d, err := strconv.Atoi(string(char)); err == nil {
				if _, err := strconv.Atoi(prev); err == nil {
					return "", ErrInvalidString
				}

				if d == 0 {
					result = result[:len(result)-1]
				}

				for i := 1; i < d; i++ {
					result = append(result, prev)
				}
			}
		} else {
			result = append(result, string(char))
		}
	}
	return strings.Join(result, ""), nil
}
