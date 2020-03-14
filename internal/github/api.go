package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GitHubAPIV3BaseURL Base URL for the GitHub API V3
const GitHubAPIV3BaseURL = "https://api.github.com"

// FetchTags fetches all tags from a GitHub slug
func FetchTags(gitHubSlug string) ([]Reference, error) {
	allReferences := make([]Reference, 0)

	url := fmt.Sprintf("%s/repos/%s/git/refs/tags", GitHubAPIV3BaseURL, gitHubSlug)
	request, requestErr := http.NewRequest(http.MethodGet, url, nil)
	if requestErr != nil {
		return allReferences, requestErr
	}

	if githubToken != "" {
		request.Header.Add("Authorization", fmt.Sprintf("token %s", githubToken))
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return allReferences, err
	}

	if response.StatusCode != http.StatusOK {
		return allReferences, generateErrorFrom(url, response)
	}

	bodyData, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return allReferences, readErr
	}

	var references []Reference
	unmarshalErr := json.Unmarshal(bodyData, &references)
	if unmarshalErr != nil {
		return allReferences, unmarshalErr
	}

	allReferences = append(allReferences, references...)

	return allReferences, nil
}
