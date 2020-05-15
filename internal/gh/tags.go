package gh

import (
	"context"
	"time"

	"github.com/google/go-github/v31/github"
)

// ListTags lists all tags for a specified repository
func ListTags(gitHubSlug string) ([]*github.Reference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	owner, repo := ToOwnerRepo(gitHubSlug)
	refs, _, refsErr := Client().Git.ListRefs(ctx, owner, repo, &github.ReferenceListOptions{Type: "tags"})
	return refs, refsErr
}
