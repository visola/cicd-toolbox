package gh

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/v31/github"
)

// ListTags lists all tags for a specified repository
func ListTags(gitHubSlug string) ([]*github.Reference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	owner, repo := ToOwnerRepo(gitHubSlug)

	// Check if repo exists
	if _, _, repoErr := Client().Repositories.Get(ctx, owner, repo); repoErr != nil {
		return nil, fmt.Errorf("Repository '%s' doesn't exist or you don't have access to it", gitHubSlug)
	}

	refs, resp, refsErr := Client().Git.ListRefs(ctx, owner, repo, &github.ReferenceListOptions{Type: "tags"})

	if resp.StatusCode == http.StatusNotFound {
		// If repo exists, 404 here means no tags
		return make([]*github.Reference, 0), nil
	}

	return refs, refsErr
}
