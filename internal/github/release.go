package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ReleaseRequestBody represents the body of a create release request
type ReleaseRequestBody struct {
	Body            string `json:"body"`
	IsDraft         bool   `json:"draft"`
	Name            string `json:"name"`
	PreRelease      bool   `json:"prerelease"`
	TagName         string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish"`
}

type ReleaseResponseBody struct {
	AssetsURL       string `json:"assets_url"`
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	TagName         string `json:"tag_name"`
	TarballURL      string `json:"tarball_url"`
	TargetCommitish string `json:"target_commitish"`
	UploadURL       string `json:"upload_url"`
	ZipballURL      string `json:"zipball_url"`
}

// CreateRelease creates a release for the repository as specified
func CreateRelease(gitHubSlug string, releaseBody ReleaseRequestBody) (ReleaseResponseBody, error) {
	var releaseResponse ReleaseResponseBody

	postRequest, createRequestErr := createGitHubPOSTRequest(fmt.Sprintf("/repos/%s/releases", gitHubSlug), releaseBody)
	if createRequestErr != nil {
		return releaseResponse, createRequestErr
	}

	responseBody, executeErr := executeRequest(postRequest)
	if executeErr != nil {
		return releaseResponse, executeErr
	}

	bodyData, readErr := ioutil.ReadAll(responseBody.Body)
	if readErr != nil {
		return releaseResponse, readErr
	}

	unmarshalErr := json.Unmarshal(bodyData, &releaseResponse)
	return releaseResponse, unmarshalErr
}

// UploadAssetsToRelease uploads the specified assets to the release asset URL
func UploadAssetsToRelease(uploadURL string, assetsToUpload []string) error {
	indexOfBrackets := strings.LastIndex(uploadURL, "{")
	uploadURL = uploadURL[:indexOfBrackets]

	for _, asset := range assetsToUpload {
		file, openErr := os.Open(asset)
		if openErr != nil {
			return openErr
		}

		uploadRequest, createRequestErr := createGitHubUploadRequest(uploadURL, file)
		if createRequestErr != nil {
			return createRequestErr
		}

		_, executeRequestErr := executeRequest(uploadRequest)
		if executeRequestErr != nil {
			return executeRequestErr
		}
	}

	return nil
}
