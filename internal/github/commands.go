package github

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	// GitHubToken stores the token to be used to authenticate with GitHub
	GitHubToken string
)

// CreateGitHubCommand creates the root GitHub command where all other GitHub
// related commands will be added to.
func CreateGitHubCommand() *cobra.Command {
	gitHubCommand := &cobra.Command{
		Use:   "github",
		Short: "All commands related to GitHub",
	}

	gitHubCommand.PersistentFlags().StringVarP(&GitHubToken, "github-token", "", "", "GitHub API Token")

	gitHubCommand.AddCommand(createListCommitsCommand())
	gitHubCommand.AddCommand(createListTagsCommand())
	return gitHubCommand
}

func createListCommitsCommand() *cobra.Command {
	listCommitsCommand := &cobra.Command{
		Use:   "list-commits {GITHUB_SLUG} {AFTER_SHA}",
		Short: "Fetch all commits after the specified SHA for a repo",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			commits, err := FetchCommits(args[0], args[1])
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%d commits found:\n", len(commits))
			for _, commit := range commits {
				fmt.Printf("-- %s --\n", commit.SHA)
				fmt.Println(commit.Commit.Message)
				fmt.Println("----------------------------------------------")
				fmt.Println()
			}
		},
	}

	return listCommitsCommand
}

func createListTagsCommand() *cobra.Command {
	listTagsCommand := &cobra.Command{
		Use:  "list-tags {GITHUB_SLUG}",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tags, err := FetchTags(args[0])
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%d tags found:\n", len(tags))
			for _, tag := range tags {
				fmt.Printf("  %s: %s\n", tag.TagName(), tag.Object.SHA)
			}
		},
	}

	return listTagsCommand
}
