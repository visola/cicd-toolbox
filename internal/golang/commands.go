package golang

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// CreateGoCommand creates the root Go command
func CreateGoCommand() *cobra.Command {
	goCommand := &cobra.Command{
		Use:   "golang",
		Short: "All commands related to the language Go",
	}

	goCommand.AddCommand(createListPackagesCommand())
	goCommand.AddCommand(createRunTestsCommand())
	return goCommand
}

func createListPackagesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list-packages",
		Short: "List packages in a directory",
		Run: func(cmd *cobra.Command, args []string) {
			packages, readErr := ListPackages(".")
			if readErr != nil {
				log.Fatal(readErr)
			}

			for _, pkg := range packages {
				fmt.Printf("%s: %s -> %s\n", pkg.Name, pkg.Dir, pkg.ImportPath)
			}
		},
	}
}

func createRunTestsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "run-tests",
		Short: "Run all the tests and collect coverage",
		Run: func(cmd *cobra.Command, args []string) {
			packages, readErr := ListPackages(".")
			if readErr != nil {
				log.Fatal(readErr)
			}

			RunTestsWithCoverage(packages)
		},
	}
}
