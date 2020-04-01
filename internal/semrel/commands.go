package semrel

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/VinnieApps/cicd-toolbox/internal/git"
	"github.com/VinnieApps/cicd-toolbox/internal/github"
	"github.com/VinnieApps/cicd-toolbox/internal/semver"
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

	semanticReleaseCommand.AddCommand(createChangeLogCommand())
	semanticReleaseCommand.AddCommand(createPublishReleaseCommand())
	semanticReleaseCommand.AddCommand(createVersionFileCommand())
	return semanticReleaseCommand
}

func createChangeLogCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "change-log {GITHUB_SLUG}",
		Short: "Prints the change log containing information about what changed since the previous release",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			nextRelease, nextReleaseErr := calculateNextRelease(args[0])
			if nextReleaseErr != nil {
				log.Fatal(nextReleaseErr)
			}

			fmt.Println(nextRelease.ChangeLog())
		},
	}
}

func createPublishReleaseCommand() *cobra.Command {
	var assetsToUpload []string
	publishReleaseCommand := &cobra.Command{
		Use:   "publish-release {GITHUB_SLUG}",
		Short: "Calculates the next version and publish the release to GitHub",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			gitHubSlug := args[0]

			if github.GitHubToken == "" {
				log.Fatal("GitHub token is required to publish a release.")
			}

			if len(assetsToUpload) == 0 {
				log.Fatal("No assets to upload.")
			}

			filesToUpload, parseAssetErr := parseAssetPaths(assetsToUpload)
			if parseAssetErr != nil {
				log.Fatalf("Error while parsing assets to upload: %v", parseAssetErr)
			}

			nextRelease, nextReleaseErr := calculateNextRelease(args[0])
			if nextReleaseErr != nil {
				log.Fatal(nextReleaseErr)
			}

			if len(nextRelease.Changes) == 0 {
				log.Printf("Nothing to release.")
				os.Exit(1)
			}

			log.Printf("New version is %s\n", nextRelease.Version.String())

			latestCommit := nextRelease.Changes[0].Commit
			tagName := fmt.Sprintf("v%s", nextRelease.Version.String())
			reference := fmt.Sprintf("refs/tags/%s", tagName)
			log.Printf("Creating reference %s to commit -> %s (%s)\n", reference, latestCommit.Message, latestCommit.ShortSHA())

			if refErr := github.CreateReference(gitHubSlug, reference, latestCommit.SHA); refErr != nil {
				log.Fatalf("Error while creating reference: %v", refErr.Error())
			}

			log.Println("Creating release...")
			releaseResponse, relErr := github.CreateRelease(gitHubSlug, github.ReleaseRequestBody{
				Body:            nextRelease.ChangeLog(),
				IsDraft:         false,
				Name:            tagName,
				PreRelease:      false,
				TagName:         tagName,
				TargetCommitish: latestCommit.SHA,
			})

			if relErr != nil {
				log.Fatalf("Error while creating release: %v", relErr)
			}

			log.Print("Uploading files to release:")
			for _, file := range filesToUpload {
				log.Printf("  %s\n", file)
			}

			uploadErr := github.UploadAssetsToRelease(releaseResponse.UploadURL, filesToUpload)
			if uploadErr != nil {
				log.Fatalf("Error while uploading asset to release: %v", uploadErr)
			}
		},
	}

	publishReleaseCommand.Flags().StringArrayVarP(&assetsToUpload, "upload", "", []string{}, "Assets to upload. File or directory. If directory, all files in the directory will be uploaded individually (non-recursively).")

	return publishReleaseCommand
}

func createVersionFileCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version-file {GITHUB_SLUG}",
		Short: "Generate a version file containing the next version",
		Long: `Generates a version file containing the next version calculated from the
latest release from the GitHub repo specified by the slug and comparing
the commit messages in the current (local) git repository.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			nextRelease, nextReleaseErr := calculateNextRelease(args[0])
			if nextReleaseErr != nil {
				log.Fatal(nextReleaseErr)
			}

			ioutil.WriteFile(".version", []byte(nextRelease.Version.String()), 0744)
		},
	}
}

func calculateNextRelease(gitHubSlug string) (Release, error) {
	tags, tagsErr := github.FetchTags(gitHubSlug)
	if tagsErr != nil {
		log.Fatal("Error while fetching tags", tagsErr)
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
			log.Fatal("Error while parsing version.", parseVersionErr)
		}

		latestVersion = version
		latestVersionSha = latestTag.Object.SHA
	}

	commits, commitsErr := git.FetchCommits(".", latestVersionSha)
	if commitsErr != nil {
		log.Fatal("Error while fetching commits.", commitsErr)
	}

	return CalculateNextRelease(latestVersion, commits)
}

func parseAssetPaths(assetsToUpload []string) ([]string, error) {
	filesToUpload := make([]string, 0)
	for _, asset := range assetsToUpload {
		stat, statErr := os.Stat(asset)
		if statErr != nil {
			return nil, statErr
		}

		if stat.IsDir() {
			entries, listErr := ioutil.ReadDir(asset)
			if listErr != nil {
				return nil, listErr
			}

			for _, dirEntry := range entries {
				entryPath := filepath.Join(asset, dirEntry.Name())
				entryStat, entryStatErr := os.Stat(entryPath)
				if entryStatErr != nil {
					return nil, entryStatErr
				}
				if entryStat.IsDir() {
					continue
				}
				filesToUpload = append(filesToUpload, entryPath)
			}
		} else {
			filesToUpload = append(filesToUpload, asset)
		}
	}

	return filesToUpload, nil
}
