package main

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	f, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer f.Close()

	fs, err := f.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}

	if offset > fs.Size() {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 || limit >= fs.Size()-offset {
		limit = fs.Size() - offset
	}

	_, err = f.Seek(offset, 0)
	if err != nil {
		return ErrUnsupportedFile
	}

	reader := bufio.NewReader(f)
	limitReader := io.LimitReader(reader, limit)

	fo, err := os.OpenFile(toPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer fo.Close()

	bar := progressbar.DefaultBytes(
		limit,
		"coping",
	)

	_, err = io.Copy(io.MultiWriter(fo, bar), limitReader)
	if err != nil {
		return ErrUnsupportedFile
	}
	return nil
}
