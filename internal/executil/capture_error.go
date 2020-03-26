package executil

import (
	"bytes"
	"os/exec"
)

// RunAndCaptureOutputIfError runs the command, and if there's an error
// capture standard out and error and generate an error with all the info
func RunAndCaptureOutputIfError(command *exec.Cmd) error {
	var standardOutput bytes.Buffer
	var standardError bytes.Buffer

	command.Stderr = &standardError
	command.Stdout = &standardOutput

	runErr := command.Run()
	if runErr != nil {
		return &ExecError{
			Command:        command,
			Source:         runErr,
			StandardError:  string(standardError.Bytes()),
			StandardOutput: string(standardOutput.Bytes()),
		}
	}

	return nil
}
