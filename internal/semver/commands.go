package semver

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// CreateSemVerCommand creates the root Semantic Version command
func CreateSemVerCommand() *cobra.Command {
	semVerCommand := &cobra.Command{
		Use:   "semver",
		Short: "All commands related to Semantic Versioning",
		Long:  "To learn more about Semantic Versioning: https://semver.org/",
	}

	semVerCommand.AddCommand(createParseSemVerCommand())
	return semVerCommand
}

func createParseSemVerCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "parse {SEMVER_STRING}",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			version, err := Parse(args[0])

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Major: %d, Minor: %d, Patch: %d\n", version.Major, version.Minor, version.Patch)
		},
	}
}
