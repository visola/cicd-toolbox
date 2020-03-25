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

	goCommand.AddCommand(createBuildCommand())
	goCommand.AddCommand(createListPackagesCommand())
	goCommand.AddCommand(createRunTestsCommand())
	return goCommand
}

func createBuildCommand() *cobra.Command {
	var architectures, linkerFlags, operatingSystems []string
	var baseName string

	buildCommand := &cobra.Command{
		Use:   "build {MAIN_FILE}",
		Short: "Build a go binary in all supported platforms",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			platforms := Platforms.WithArchitectures(architectures...).
				WithOperatingSystems(operatingSystems...)

			for _, platform := range platforms {
				fmt.Printf("Compiling binary for '%s' '%s'\n", platform.Architecture, platform.OperatingSystem)
			}

			buildSpec := BuildSpecification{
				BaseName:        baseName,
				FileToBuild:     args[0],
				LinkerArguments: linkerFlags,
				Platforms:       platforms,
			}

			if buildErr := buildSpec.Build(); buildErr != nil {
				log.Fatal(buildErr)
			}
		},
	}

	buildCommand.Flags().StringVarP(&baseName, "base-name", "n", "main", "Base name for the binary")

	buildCommand.Flags().StringArrayVarP(&architectures, "arch", "a", []string{}, "Architecture to compile for")
	buildCommand.Flags().StringArrayVarP(&linkerFlags, "ldflags", "", []string{}, "Flags to pass for the linker")
	buildCommand.Flags().StringArrayVarP(&operatingSystems, "os", "", []string{}, "Operating System to compile for")

	return buildCommand
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
