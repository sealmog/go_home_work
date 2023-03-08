package main

import (
	"errors"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	_ = env
	if len(cmd) > 3 {
		command := cmd[0]
		args := cmd[1:]
		cm := exec.Command(command, args...)

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
	return 22
}
