package golang

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/VinnieApps/cicd-toolbox/internal/executil"
)

// BuildSpecification represents a spec to build a go binary
type BuildSpecification struct {
	Architecture    string
	BaseName        string
	Extension       string
	FileToBuild     string
	LinkerVariables []Variable
	OperatingSystem string
}

// Variable represents a variable with a name and value
type Variable struct {
	Name  string
	Value string
}

// Build builds the executable using the build spec.
func (buildSpec *BuildSpecification) Build() error {
	buildDir := fmt.Sprintf("build/binaries/%s_%s", buildSpec.Architecture, buildSpec.OperatingSystem)

	os.MkdirAll(buildDir, 0744)

	commandArgs := []string{"build"}
	outFileName := "main"
	if buildSpec.BaseName != "" {
		outFileName = buildSpec.BaseName
	}

	if buildSpec.Extension != "" {
		outFileName = fmt.Sprintf("%s.%s", outFileName, buildSpec.Extension)
	}
	commandArgs = append(commandArgs, "-o", filepath.Join(buildDir, outFileName))

	commandArgs = append(commandArgs, buildSpec.FileToBuild)

	command := exec.Command("go", commandArgs...)
	env := []string{
		fmt.Sprintf("GOARCH=%s", buildSpec.Architecture),
		fmt.Sprintf("GOOS=%s", buildSpec.OperatingSystem),
	}

	command.Env = append(os.Environ(), env...)

	return executil.RunAndCaptureOutputIfError(command)
}
