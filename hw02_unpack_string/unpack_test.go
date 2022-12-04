package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "ASCII", input: "a4bc2d5e", expected: "aaaabccddddde"},
		{name: "ASCII", input: "abccd", expected: "abccd"},
		{name: "ASCII", input: "", expected: ""},
		{name: "ASCII", input: "aaa0b", expected: "aab"},

		{name: "nonASCII", input: "а1б2в3", expected: "аббввв"},
		{name: "nonASCII", input: "a0", expected: ""},
		{name: "nonASCII", input: "a1", expected: "a"},
		{name: "nonASCII", input: "я9", expected: "яяяяяяяяя"},
		{name: "nonASCII, 3 byte", input: "สวัสดี", expected: "สวัสดี"},
		{name: "nonASCII: 3 byte", input: "สวัส4ดี", expected: "สวัสสสสดี"},
		{name: "nonASCII: emoji ", input: "🙃0", expected: ""},
		{name: "nonASCII: emoji", input: "🙂9", expected: "🙂🙂🙂🙂🙂🙂🙂🙂🙂"},

		{name: "NonArabic", input: "১২৩", expected: "১২৩"},
		{name: "NonArabic", input: "১2২৩0", expected: "১১২"},
		{name: "NonArabic", input: "੩4", expected: "੩੩੩੩"},

		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "0abc", "0абв", "🙃10", "১১44"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
