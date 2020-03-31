package golang

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/VinnieApps/cicd-toolbox/internal/executil"
)

// RunTestsWithCoverage runs tests for all packages and generate coverage data.
// This method is guarantee to generate coverage counts even when the package has no
// tests.
func RunTestsWithCoverage(packages []*build.Package) error {
	os.MkdirAll("build/coverage", 0744)
	filesToCollect := make([]string, 0)
	for i, pkg := range packages {
		tempCoverOuptutFile := fmt.Sprintf("build/coverage/temp_%d.out", i)
		filesToCollect = append(filesToCollect, tempCoverOuptutFile)

		packageTestFile := filepath.Join(pkg.Dir, fmt.Sprintf("%s_test.go", pkg.Name))
		if _, err := os.Stat(packageTestFile); os.IsNotExist(err) {
			ioutil.WriteFile(packageTestFile, []byte("package "+pkg.Name), 0644)
			defer os.Remove(packageTestFile)
		}
	}

	cmd := exec.Command("go", "test", "-cover", "-coverprofile=build/coverage/all.out", "./...")
	return executil.RunAndCaptureOutputIfError(cmd)
}
