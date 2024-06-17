package pkg

import (
	"io"
	"os/exec"
)

type CommandOutput struct {
	ReturnCode int
	Stdout     string
	Stderr     string
}

func RunCommand(command string) (*CommandOutput, error) {

	cmd := exec.Command(command)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	byteStderr, _ := io.ReadAll(stderr)
	byteStdout, _ := io.ReadAll(stdout)

	strStderr := string(byteStderr)
	strStdout := string(byteStdout)

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return &CommandOutput{
				ReturnCode: exitError.ExitCode(),
				Stdout:     strStdout,
				Stderr:     strStderr,
			}, nil
		} else {
			return nil, err
		}
	}

	return &CommandOutput{
		ReturnCode: 0,
		Stdout:     strStdout,
		Stderr:     strStderr,
	}, nil
}
