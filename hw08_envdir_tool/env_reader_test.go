package main

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mEnv() Environment {
	var Env Environment = make(map[string]EnvValue)
	tst := []struct {
		name   string
		value  string
		remove bool
	}{
		{name: "BAR", value: "bar", remove: false},
		{name: "EMPTY", value: "", remove: false},
		{name: "FOO", value: "   foo\nwith new line", remove: false},
		{name: "HELLO", value: "\"hello\"", remove: false},
		{name: "UNSET", value: "", remove: true},
	}
	for _, ts := range tst {
		Env[ts.name] = EnvValue{
			Value:      ts.value,
			NeedRemove: ts.remove,
		}
	}
	return Env
}

func TestReadDir(t *testing.T) {
	tests := []struct {
		name     string
		dir      string
		expected bool
		err      error
	}{
		{name: "get env success", dir: "./testdata/env", expected: true, err: nil},
		{name: "get blank env", dir: "/tmp", expected: false, err: nil},
		{name: "get unread dir", dir: "/123", expected: false, err: errors.New("open /123: no such file or directory")},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := ReadDir(tc.dir)
			if err != nil {
				assert.EqualErrorf(t, err, tc.err.Error(), "error message")
				return
			}
			assert.NoError(t, err)
			eq := reflect.DeepEqual(result, mEnv())
			require.Equal(t, tc.expected, eq)
		})
	}
}
