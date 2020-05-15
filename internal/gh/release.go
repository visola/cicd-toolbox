package gh

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/google/go-github/v31/github"
)

// CreateRelease creates a release for the repository as specified
func CreateRelease(gitHubSlug, body, tagName, sha string) (*github.RepositoryRelease, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	owner, repo := ToOwnerRepo(gitHubSlug)
	release := &github.RepositoryRelease{
		Body:            &body,
		Name:            &tagName,
		TagName:         &tagName,
		TargetCommitish: &sha,
	}
	release, _, err := Client().Repositories.CreateRelease(ctx, owner, repo, release)
	return release, err
}

// UploadAssetsToRelease uploads the specified assets to the specified release in the specified repository
func UploadAssetsToRelease(gitHubSlug string, releaseResponse *github.RepositoryRelease, assetsToUpload []string) error {
	owner, repo := ToOwnerRepo(gitHubSlug)

	for _, asset := range assetsToUpload {
		file, openErr := os.OpenFile(asset, os.O_RDONLY, 0644)
		if openErr != nil {
			return openErr
		}
		defer file.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		opts := &github.UploadOptions{
			Name: filepath.Base(file.Name()),
		}

		_, _, uploadErr := Client().Repositories.UploadReleaseAsset(ctx, owner, repo, *releaseResponse.ID, opts, file)
		if uploadErr != nil {
			return uploadErr
		}
	}

	return nil
}
