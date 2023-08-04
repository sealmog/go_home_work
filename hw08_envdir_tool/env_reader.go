package main

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var ErrUnsupportedFile = errors.New("unsupported file")

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	mEnv := make(map[string]EnvValue)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fInfo, err := file.Info()
		if err != nil {
			return nil, err
		}
		if fInfo.Size() == 0 {
			mEnv[file.Name()] = EnvValue{
				Value:      "",
				NeedRemove: true,
			}
			continue
		}

		f, err := os.Open(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		line, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}

		bValue := bytes.ReplaceAll(line, []byte{0x00}, []byte("\n"))
		sValue := strings.TrimRight(string(bValue), " \t")

		mEnv[file.Name()] = EnvValue{
			Value:      sValue,
			NeedRemove: false,
		}
	}
	return mEnv, nil
}
