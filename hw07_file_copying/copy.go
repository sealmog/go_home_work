package main

import (
	"errors"
	"os"
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

	buf := make([]byte, limit)
	_, err = f.Seek(offset, 0)
	if err != nil {
		return ErrUnsupportedFile
	}

	_, err = f.Read(buf)
	if err != nil {
		return ErrUnsupportedFile
	}

	err = os.WriteFile(toPath, buf, 0o600)
	if err != nil {
		return ErrUnsupportedFile
	}
	return nil
}
