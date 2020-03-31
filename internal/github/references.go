package github

import (
	"fmt"
	"strings"
)

// ReferenceTagPrefix References that have this prefix are tags
const ReferenceTagPrefix = "refs/tags/"

// Reference represents a reference to an object in the GitHub Git Database
// https://developer.github.com/v3/git/
type Reference struct {
	Object    Object `json:"object"`
	Reference string `json:"ref"`
}

// TagName returns the tag name if the reference is a tag, otherwise, returns empty
func (ref Reference) TagName() string {
	if strings.HasPrefix(ref.Reference, ReferenceTagPrefix) {
		return ref.Reference[len(ReferenceTagPrefix):]
	}

	return ""
}

type createReferenceBody struct {
	Reference string `json:"ref"`
	SHA       string `json:"sha"`
}

// CreateReference creates a reference in GitHub. Name should be of the format: refs/type/name
// e.g.: refs/tag/v1.2.0
func CreateReference(gitHubSlug, name, sha string) error {
	ref := createReferenceBody{
		Reference: name,
		SHA:       sha,
	}

	createRefRequest, createRefErr := createGitHubPOSTRequest(fmt.Sprintf("/repos/%s/git/refs", gitHubSlug), ref)
	if createRefErr != nil {
		return createRefErr
	}

	_, err := executeRequest(createRefRequest)
	return err
}
