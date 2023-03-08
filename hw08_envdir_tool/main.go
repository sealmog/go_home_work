package main

import (
	"log"
	"os"
)

func exportEnv(mEnv Environment) error {
	for k, v := range mEnv {
		if v.NeedRemove {
			err := os.Unsetenv(v.Value)
			if err != nil {
				return err
			}
		}

		err := os.Setenv(k, v.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 6 {
		log.Fatal("Program executed without required args")
		os.Exit(22)
	}

	path := os.Args[1]
	mEnv, err := ReadDir(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(20)
	}

	err = exportEnv(mEnv)
	if err != nil {
		log.Fatal(err)
		os.Exit(10)
	}

	cmd := make([]string, 0)
	cmd = append(cmd, os.Args[2:]...)

	exitCode := RunCmd(cmd, mEnv)
	os.Exit(exitCode)
}
