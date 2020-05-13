package executil

import (
	"bytes"
	"os/exec"
)

// RunAndCaptureOutputIfError runs the command, and if there's an error
// capture standard out and error and generate an error with all the info
func RunAndCaptureOutputIfError(command *exec.Cmd) (string, string, error) {
	var standardOutput bytes.Buffer
	var standardError bytes.Buffer

	command.Stderr = &standardError
	command.Stdout = &standardOutput

	runErr := command.Run()

	stdErr := string(standardOutput.Bytes())
	stdOut := string(standardOutput.Bytes())
	if runErr != nil {
		return stdOut, stdErr, &ExecError{
			Command:        command,
			Source:         runErr,
			StandardError:  stdErr,
			StandardOutput: stdOut,
		}
	}

	return stdOut, stdErr, nil
}
