package gh

import (
	"context"
	"time"

	"github.com/google/go-github/v31/github"
)

// CreateReference creates a reference in GitHub.
// References should be of the format: refs/type/name, e.g.: refs/tag/v1.2.0
func CreateReference(gitHubSlug string, reference string, sha string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gitHubClient := Client()
	owner, repo := ToOwnerRepo(gitHubSlug)

	_, _, refErr := gitHubClient.Git.CreateRef(ctx, owner, repo, &github.Reference{
		Ref:    &reference,
		Object: &github.GitObject{SHA: &sha},
	})
	return refErr
}
