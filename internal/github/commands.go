package github

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	githubToken string
)

func CreateGitHubCommand() *cobra.Command {
	gitHubCommand := &cobra.Command{
		Use:   "github",
		Short: "All commands related to GitHub",
	}

	gitHubCommand.PersistentFlags().StringVarP(&githubToken, "github-token", "t", "", "GitHub API Token")

	gitHubCommand.AddCommand(createListTagsCommand())
	return gitHubCommand
}

func createListTagsCommand() *cobra.Command {
	listTagsCommand := &cobra.Command{
		Use:  "list-tags {GITHUB_SLUG}",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tags, err := FetchTags(args[0])
			if err != nil {
				panic(err)
			}

			fmt.Printf("%d tags found:\n", len(tags))
			for _, tag := range tags {
				fmt.Printf("  %s: %s\n", tag.TagName(), tag.Object.SHA)
			}
		},
	}

	return listTagsCommand
}
