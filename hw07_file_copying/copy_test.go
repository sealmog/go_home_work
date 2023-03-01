package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name     string
		from     string
		to       string
		offset   int64
		limit    int64
		expected int64
	}{
		{
			name: "offset 0, limit 0", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 0, limit: 0, expected: 6617,
		},
		{
			name: "offset 0, limit 10", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 0, limit: 10, expected: 10,
		},
		{
			name: "offset 0, limit 1000", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 0, limit: 1000, expected: 1000,
		},
		{
			name: "offset 0, limit 10000", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 0, limit: 10000, expected: 6617,
		},
		{
			name: "offset 100, limit 1000", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 100, limit: 1000, expected: 1000,
		},
		{
			name: "offset 6000, limit 1000", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 6000, limit: 1000, expected: 617,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println()
			err := Copy(tc.from, tc.to, tc.offset, tc.limit)
			require.NoError(t, err)

			f, err := os.Open(tc.to)
			check(err)
			fs, err := f.Stat()
			check(err)

			result := fs.Size()
			require.Equal(t, tc.expected, result)
		})
	}
}
