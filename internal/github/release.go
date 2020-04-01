package github

import "fmt"

// ReleaseBody represents the body of a create release request
type ReleaseBody struct {
	Body            string `json:"body"`
	IsDraft         bool   `json:"draft"`
	Name            string `json:"name"`
	PreRelease      bool   `json:"prerelease"`
	TagName         string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish"`
}

// CreateRelease creates a release for the repository as specified
func CreateRelease(gitHubSlug string, releaseBody ReleaseBody) error {
	postRequest, createRequestErr := createGitHubPOSTRequest(fmt.Sprintf("/repos/%s/releases", gitHubSlug), releaseBody)
	if createRequestErr != nil {
		return createRequestErr
	}

	_, executeErr := executeRequest(postRequest)
	return executeErr
}
