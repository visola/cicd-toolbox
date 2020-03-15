package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GitHubAPIV3BaseURL Base URL for the GitHub API V3
const GitHubAPIV3BaseURL = "https://api.github.com"

// FetchCommits fetches all commits before a specified SHA for a repository
func FetchCommits(gitHubSlug string, beforeSha string) ([]Commit, error) {
	allCommits := make([]Commit, 0)

	url := fmt.Sprintf("%s/repos/%s/commits", GitHubAPIV3BaseURL, gitHubSlug)
FetchLoop:
	for true {
		request, requestErr := createGitHubRequest(url)
		if requestErr != nil {
			return allCommits, requestErr
		}

		client := &http.Client{}
		response, responseErr := client.Do(request)
		if responseErr != nil {
			return allCommits, responseErr
		}

		if response.StatusCode != http.StatusOK {
			return allCommits, generateErrorFrom(url, response)
		}

		bodyData, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			return allCommits, readErr
		}

		var commits []Commit
		unmarshalErr := json.Unmarshal(bodyData, &commits)
		if unmarshalErr != nil {
			return allCommits, unmarshalErr
		}

		for _, commit := range commits {
			if commit.SHA == beforeSha {
				break FetchLoop
			}
			allCommits = append(allCommits, commit)
		}

		linkHeader := response.Header.Get("link")
		nextLink := FindNextLink(ParseLinks(linkHeader))
		if nextLink == nil {
			break
		}

		url = nextLink.URL
	}

	return allCommits, nil
}

// FetchTags fetches all tags from a GitHub slug
func FetchTags(gitHubSlug string) ([]Reference, error) {
	allReferences := make([]Reference, 0)

	url := fmt.Sprintf("%s/repos/%s/git/refs/tags", GitHubAPIV3BaseURL, gitHubSlug)
	request, requestErr := createGitHubRequest(url)
	if requestErr != nil {
		return allReferences, requestErr
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

func createGitHubRequest(url string) (*http.Request, error) {
	request, requestErr := http.NewRequest(http.MethodGet, url, nil)
	if requestErr != nil {
		return request, requestErr
	}

	if githubToken != "" {
		request.Header.Add("Authorization", fmt.Sprintf("token %s", githubToken))
	}

	return request, nil
}
