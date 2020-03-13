package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchReferencesMatching(gitHubSlug string, matching string) ([]Reference, error) {
	allReferences := make([]Reference, 0)

	url := fmt.Sprintf("%s/repos/%s/git/matching-refs/%s", GITHUB_API_V3_BASE, gitHubSlug, matching)
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

func FetchTags(gitHubSlug string) ([]Reference, error) {
	return FetchReferencesMatching(gitHubSlug, "tags")
}
