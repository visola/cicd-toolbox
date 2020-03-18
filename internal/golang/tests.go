package golang

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// RunTestsWithCoverage runs tests for all packages and generate coverage data.
// This method is guarantee to generate coverage counts even when the package has no
// tests.
func RunTestsWithCoverage(packages []*build.Package) error {
	os.MkdirAll("build/coverage", 0744)
	filesToCollect := make([]string, 0)
	for i, pkg := range packages {
		fmt.Printf("Running tests for %s\n", pkg.ImportPath)

		tempCoverOuptutFile := fmt.Sprintf("build/coverage/temp_%d.out", i)
		filesToCollect = append(filesToCollect, tempCoverOuptutFile)

		packageTestFile := filepath.Join(pkg.Dir, fmt.Sprintf("%s_test.go", pkg.Name))
		if _, err := os.Stat(packageTestFile); os.IsNotExist(err) {
			ioutil.WriteFile(packageTestFile, []byte("package "+pkg.Name), 0644)
			defer os.Remove(packageTestFile)
		}

		cmd := exec.Command("go", "test", "-cover", "-coverprofile="+tempCoverOuptutFile, pkg.ImportPath)
		runErr := cmd.Run()
		if runErr != nil {
			log.Printf("Error running test for: %s\n", pkg.ImportPath)
			output, _ := cmd.CombinedOutput()
			log.Println(string(output))
			log.Fatal(runErr)
		}
	}

	allData := "mode: set\n"
	for _, file := range filesToCollect {
		data, readErr := ioutil.ReadFile(file)
		if readErr != nil {
			log.Fatal(readErr)
		}

		allData += strings.Join(strings.Split(string(data), "\n")[1:], "\n")
	}
	ioutil.WriteFile("build/coverage/collected.out", []byte(allData), 0644)

	return nil
}
