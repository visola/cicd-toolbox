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
	if parallel <= 0 {
		parallel = 1
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
	outFileName := "main"
	if buildSpec.BaseName != "" {
		outFileName = buildSpec.BaseName
	}

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

	return executil.RunAndCaptureOutputIfError(command)
}
