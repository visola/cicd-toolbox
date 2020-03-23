package main

import (
	"fmt"

	"github.com/VinnieApps/cicd-toolbox/internal/github"
	"github.com/VinnieApps/cicd-toolbox/internal/golang"
	"github.com/VinnieApps/cicd-toolbox/internal/semrel"
	"github.com/VinnieApps/cicd-toolbox/internal/semver"
	"github.com/spf13/cobra"
)

var version string

func main() {
	rootCommand := &cobra.Command{
		Use:   "cicd",
		Short: "Toolbox for your CI/CD processes",
		Long:  "An unopinionated toolbox for all Continuous Integration and Deliver needs.",
	}

	rootCommand.AddCommand(github.CreateGitHubCommand())
	rootCommand.AddCommand(golang.CreateGoCommand())
	rootCommand.AddCommand(semrel.CreateSemanticReleaseCommand())
	rootCommand.AddCommand(semver.CreateSemVerCommand())
	rootCommand.AddCommand(createVersionCommand())
	rootCommand.Execute()
}

func createVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show the version of the toolbox",
		Run: func(cmd *cobra.Command, args []string) {
			if version == "" {
				fmt.Println("No version defined.")
				return
			}

			fmt.Println(version)
		},
	}
}
