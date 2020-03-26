package executil

import (
	"fmt"
	"os/exec"
)

// ExecError represents an error during executing an external command.
type ExecError struct {
	Command        *exec.Cmd
	Source         error
	StandardError  string
	StandardOutput string
}

func (err *ExecError) Error() string {
	message := fmt.Sprintf("Error while executing command:\n  '%s'", err.Command.String())
	message = fmt.Sprintf("%s\nError: %s", message, err.Source.Error())

	if err.StandardOutput != "" {
		message = fmt.Sprintf("%s\n--- Standard Output ---\n%s\n-----------------------", message, err.StandardOutput)
	}

	if err.StandardError != "" {
		message = fmt.Sprintf("%s\n--- Standard Error ---\n%s\n-----------------------", message, err.StandardError)
	}

	return message
}
