package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) < 4 {
		return 22
	}

	command := cmd[0]
	args := cmd[1:]
	cm := exec.Command(command, args...)

	sEnv := make([]string, 0)
	for name, v := range env {
		sEnv = append(sEnv, fmt.Sprintf("%v=%v", name, v.Value))
	}
	cm.Env = append(cm.Environ(), sEnv...)

	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr
	err := cm.Run()
	if err != nil {
		var exitError *exec.ExitError

		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
		return 1
	}
	return
}
