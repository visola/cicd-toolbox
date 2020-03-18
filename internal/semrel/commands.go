package semrel

import (
	"io/ioutil"
	"log"

	"github.com/VinnieApps/cicd-tools/internal/git"
	"github.com/VinnieApps/cicd-tools/internal/github"
	"github.com/VinnieApps/cicd-tools/internal/semver"
	"github.com/spf13/cobra"
)

// CreateSemanticReleaseCommand creates the root Semantic Release command
func CreateSemanticReleaseCommand() *cobra.Command {
	semanticReleaseCommand := &cobra.Command{
		Use:   "semantic-release",
		Short: "All commands related to Semantic Release",
		Long:  "To learn more about Semantic Versioning: https://semver.org/",
	}

	semanticReleaseCommand.PersistentFlags().StringVarP(&github.GitHubToken, "github-token", "", "", "GitHub API Token")

	semanticReleaseCommand.AddCommand(createVersionFileCommand())
	return semanticReleaseCommand
}

func createVersionFileCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "version-file {GITHUB_SLUG}",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			gitHubSlug := args[0]

			tags, tagsErr := github.FetchTags(gitHubSlug)
			if tagsErr != nil {
				log.Fatal(tagsErr)
			}

			var latestVersion semver.Version
			var latestVersionSha string
			if len(tags) == 0 {
				latestVersion = semver.Version{}
				latestVersionSha = ""
			} else {
				latestTag := tags[len(tags)-1]
				version, parseVersionErr := semver.Parse(latestTag.TagName())
				if parseVersionErr != nil {
					log.Fatal(parseVersionErr)
				}

				latestVersion = version
				latestVersionSha = latestTag.Object.SHA
			}

			commits, commitsErr := git.FetchCommits(".", latestVersionSha)
			if commitsErr != nil {
				log.Fatal(commitsErr)
			}

			nextRelease, nextReleaseErr := CalculateNextRelease(latestVersion, commits)
			if nextReleaseErr != nil {
				log.Fatal(nextReleaseErr)
			}

			ioutil.WriteFile(".version", []byte(nextRelease.Version.String()), 0744)
		},
	}
}
