package build

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/VinnieApps/cicd-toolbox/internal/build/nodejs"
)

var dirsToIgnore = []string{
	".cache/",
	".git/",
	"node_modules/",
}

var introspectors = []Introspector{
	&nodejs.Introspector{},
}

type Introspector interface {
	Matches(files []string) bool
}

func introspect(dirPath string) {
	absPath, absErr := filepath.Abs(dirPath)
	if absErr != nil {
		panic(absErr)
	}
	fmt.Printf("Introspecting directory: %s\n", absPath)
	files := make([]string, 0)
	filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, ignoredPath := range dirsToIgnore {
			if strings.Contains(path, ignoredPath) {
				return nil
			}
		}

		files = append(files, path)
		return nil
	})

	for _, introspector := range introspectors {
		match := introspector.Matches(files)
		if match {
			fmt.Printf("Matched introspector: %s", introspector)
		}
	}
}
