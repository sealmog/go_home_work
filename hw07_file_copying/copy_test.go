package main

import (
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
		err      error
	}{
		{
			name: "offset 0, limit 0", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 0, limit: 0, expected: 6617, err: nil,
		},
		{
			name: "offset 0, limit 10", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 0, limit: 10, expected: 10, err: nil,
		},
		{
			name: "offset 0, limit 1000", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 0, limit: 1000, expected: 1000, err: nil,
		},
		{
			name: "offset 0, limit 10000", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 0, limit: 10000, expected: 6617, err: nil,
		},
		{
			name: "offset 100, limit 1000", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 100, limit: 1000, expected: 1000, err: nil,
		},
		{
			name: "offset 6000, limit 1000", from: "testdata/input.txt", to: "/tmp/output.txt",
			offset: 6000, limit: 1000, expected: 617, err: nil,
		},

		{
			name: "/dev/random offset 0, limit 0", from: "/dev/random", to: "/tmp/output.txt",
			offset: 0, limit: 0, expected: 0, err: nil,
		},

		{
			name: "/dev/random offset 0, limit 10", from: "/dev/random", to: "/tmp/output.txt",
			offset: 0, limit: 10, expected: 0, err: nil,
		},

		{
			name: "/dev/urandom offset 0, limit 0", from: "/dev/urandom", to: "/tmp/output.txt",
			offset: 0, limit: 0, expected: 0, err: nil,
		},

		{
			name: "/dev/urandom offset 0, limit 10", from: "/dev/urandom", to: "/tmp/output.txt",
			offset: 0, limit: 10, expected: 0, err: nil,
		},

		{
			name: "/dev/zero offset 0, limit 0", from: "/dev/zero", to: "/tmp/output.txt",
			offset: 0, limit: 0, expected: 0, err: nil,
		},

		{
			name: "/dev/zero offset 0, limit 10", from: "/dev/zero", to: "/tmp/output.txt",
			offset: 0, limit: 10, expected: 0, err: nil,
		},

		{
			name: "/dev/null offset 0, limit 0", from: "/dev/null", to: "/tmp/output.txt",
			offset: 0, limit: 0, expected: 0, err: nil,
		},

		{
			name: "/dev/null offset 0, limit 10", from: "/dev/null", to: "/tmp/output.txt",
			offset: 0, limit: 10, expected: 0, err: nil,
		},

		{
			name: "/dev/random offset 10, limit 10", from: "/dev/random", to: "/tmp/output.txt",
			offset: 10, limit: 10, expected: 0, err: ErrOffsetExceedsFileSize,
		},

		{
			name: "/dev/urandom offset 10, limit 10", from: "/dev/urandom", to: "/tmp/output.txt",
			offset: 10, limit: 10, expected: 0, err: ErrOffsetExceedsFileSize,
		},

		{
			name: "/dev/zero offset 10, limit 10", from: "/dev/zero", to: "/tmp/output.txt",
			offset: 10, limit: 10, expected: 0, err: ErrOffsetExceedsFileSize,
		},

		{
			name: "/dev/null offset 10, limit 10", from: "/dev/null", to: "/tmp/output.txt",
			offset: 10, limit: 10, expected: 0, err: ErrOffsetExceedsFileSize,
		},

		{
			name: "open non existing file", from: "testdata/nofile", to: "/tmp/output.txt",
			offset: 0, limit: 0, expected: 0, err: ErrUnsupportedFile,
		},

		{
			name: "write to /root dir", from: "testdata/input.txt", to: "/root/output.txt",
			offset: 0, limit: 0, expected: 0, err: ErrUnsupportedFile,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := Copy(tc.from, tc.to, tc.offset, tc.limit)
			require.ErrorIs(t, err, tc.err)

			if err != nil {
				return
			}

			f, err := os.Open(tc.to)
			check(err)
			fs, err := f.Stat()
			check(err)
			defer f.Close()

			result := fs.Size()
			require.Equal(t, tc.expected, result)
		})
	}
}
