package main

import (
	"github.com/VinnieApps/cicd-tools/internal/github"
	"github.com/spf13/cobra"
)

func main() {
	rootCommand := &cobra.Command{
		Use:   "cicd",
		Short: "Toolbox for your CI/CD processes",
		Long:  "An unopinionated toolbox for all Continuous Integration and Deliver needs.",
	}

	rootCommand.AddCommand(github.CreateGitHubCommand())
	rootCommand.Execute()
}
