package build

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// CreateCommands creates the root Build command
func CreateCommands() []*cobra.Command {
	return []*cobra.Command{
		createPackageCommand(),
	}
}

func createPackageCommand() *cobra.Command {
	var additionalFiles []string
	var baseName string

	packageCommand := &cobra.Command{
		Use:   "package {dirs-children}",
		Short: "Create ZIP package containing multiple files",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("One argument is required")
				cmd.Help()
				os.Exit(-1)
			}

			rootDir := args[0]
			dirs, listErr := getChildrenDirectories(rootDir)
			if listErr != nil {
				panic(listErr)
			}

			outputDir := "build/packages"
			if err := os.MkdirAll(outputDir, 0744); err != nil {
				panic(err)
			}

			for _, dir := range dirs {
				outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_%s.zip", baseName, dir.Name()))
				fmt.Printf("Creating package %s\n", outputFile)
				CreatePackage(outputFile, filepath.Join(rootDir, dir.Name()), additionalFiles...)
			}
		},
	}

	currentDir, cdErr := os.Getwd()
	if cdErr != nil {
		panic(cdErr)
	}

	packageCommand.Flags().StringVarP(&baseName, "base-name", "n", filepath.Base(currentDir), "Base name for the packages. Default to working dir name.")
	packageCommand.Flags().StringArrayVarP(&additionalFiles, "add-file", "f", []string{}, "A file to add to all packages created")

	return packageCommand
}

func getChildrenDirectories(dir string) ([]os.FileInfo, error) {
	childrenDir := make([]os.FileInfo, 0)
	children, listErr := ioutil.ReadDir(dir)
	if listErr != nil {
		return nil, listErr
	}
	for _, child := range children {
		if !child.IsDir() {
			continue
		}
		childrenDir = append(childrenDir, child)
	}
	return childrenDir, nil
}
