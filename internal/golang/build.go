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
	BaseName        string
	FileToBuild     string
	LinkerArguments []string
	Platforms       PlatformList
}

// Build builds the binaries using the build spec
func (buildSpec *BuildSpecification) Build() error {
	outFileName := "main"
	if buildSpec.BaseName != "" {
		outFileName = buildSpec.BaseName
	}

	for _, platform := range buildSpec.Platforms {
		buildDir := fmt.Sprintf("build/binaries/%s_%s", platform.Architecture, platform.OperatingSystem)
		os.MkdirAll(buildDir, 0744)

		commandArgs := []string{"build", "-a", "-installsuffix", "cgo"}

		for _, linkerArg := range buildSpec.LinkerArguments {
			commandArgs = append(commandArgs, "-ldflags", linkerArg)
		}

		if platform.Extension != "" {
			outFileName = fmt.Sprintf("%s.%s", outFileName, platform.Extension)
		}

		commandArgs = append(commandArgs, "-o", filepath.Join(buildDir, outFileName))
		commandArgs = append(commandArgs, buildSpec.FileToBuild)

		command := exec.Command("go", commandArgs...)

		command.Env = append(os.Environ(), []string{
			fmt.Sprintf("GOARCH=%s", platform.Architecture),
			fmt.Sprintf("GOOS=%s", platform.OperatingSystem),
		}...)

		runErr := executil.RunAndCaptureOutputIfError(command)
		if runErr != nil {
			return runErr
		}
	}

	return nil
}
