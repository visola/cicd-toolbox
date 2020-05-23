package golang

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/VinnieApps/cicd-toolbox/internal/executil"
)

const binariesDir = "build/binaries"

// BuildSpecification represents a spec to build a go binary
type BuildSpecification struct {
	BaseName        string
	FileToBuild     string
	LinkerArguments []string
	Platforms       PlatformList
}

// Build builds the binaries using the build spec
func (buildSpec *BuildSpecification) Build() []error {
	parallel := runtime.NumCPU() - 1
	if parallel < 0 {
		parallel = 0
	}

	var errorLock sync.Mutex
	var wg sync.WaitGroup
	errors := make([]error, 0)
	semaphore := make(chan int, parallel)
	for _, platform := range buildSpec.Platforms {
		wg.Add(1)
		go func(platform Platform) {
			defer wg.Done()
			semaphore <- 1
			if err := compileForPlatform(*buildSpec, platform); err != nil {
				errorLock.Lock()
				defer errorLock.Unlock()
				errors = append(errors, err)
			}
			<-semaphore
		}(platform)
	}
	wg.Wait()

	return errors
}

func compileForPlatform(buildSpec BuildSpecification, platform Platform) error {
	outputDir := filepath.Join(binariesDir, fmt.Sprintf("%s_%s", platform.OperatingSystem, platform.Architecture))
	os.MkdirAll(outputDir, 0744)

	outFileName := "main"
	if buildSpec.BaseName != "" {
		outFileName = buildSpec.BaseName
	}

	if platform.Extension != "" {
		outFileName = fmt.Sprintf("%s.%s", outFileName, platform.Extension)
	}

	commandArgs := []string{"build", "-installsuffix", "cgo"}

	for _, linkerArg := range buildSpec.LinkerArguments {
		commandArgs = append(commandArgs, "-ldflags", linkerArg)
	}

	commandArgs = append(commandArgs, "-o", filepath.Join(outputDir, outFileName))
	commandArgs = append(commandArgs, buildSpec.FileToBuild)

	command := exec.Command("go", commandArgs...)

	command.Env = append(os.Environ(), []string{
		fmt.Sprintf("GOARCH=%s", platform.Architecture),
		fmt.Sprintf("GOOS=%s", platform.OperatingSystem),
	}...)

	_, _, err := executil.RunAndCaptureOutputIfError(command)
	return err
}
