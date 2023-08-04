package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	runs := []struct {
		name     string
		cmd      []string
		expected int
	}{
		{name: "run success", cmd: []string{"/bin/bash", "./testdata/echo.sh", "arg1=1", "arg2=2"}, expected: 0},
		{name: "invalid argument", cmd: []string{"/bin/bash", "./testdata/echo.sh", "arg1=1"}, expected: 22},
		{name: "operation not permitted", cmd: []string{"/root", "./testdata/echo.sh", "arg1=1", "arg2=2"}, expected: 1},
	}

	var Env Environment = make(map[string]EnvValue)

	for _, tc := range runs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := RunCmd(tc.cmd, Env)
			require.Equal(t, tc.expected, result)
		})
	}
}
