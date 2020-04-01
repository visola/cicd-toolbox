package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// GitHubAPIV3BaseURL Base URL for the GitHub API V3
const GitHubAPIV3BaseURL = "https://api.github.com"

// FetchCommits fetches all commits after a specified SHA for a repository
func FetchCommits(gitHubSlug string, afterSha string) ([]Commit, error) {
	allCommits := make([]Commit, 0)

	url := fmt.Sprintf("/repos/%s/commits", gitHubSlug)
FetchLoop:
	for true {
		request, requestErr := createGitHubRequest(url)
		if requestErr != nil {
			return allCommits, requestErr
		}

		response, responseErr := executeRequest(request)
		if responseErr != nil {
			return allCommits, responseErr
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
			if commit.SHA == afterSha {
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

	url := fmt.Sprintf("/repos/%s/git/refs/tags", gitHubSlug)
	request, requestErr := createGitHubRequest(url)
	if requestErr != nil {
		return allReferences, requestErr
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return allReferences, err
	}

	if response.StatusCode == http.StatusNotFound {
		return allReferences, nil
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
	request, requestErr := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", GitHubAPIV3BaseURL, url), nil)
	if requestErr != nil {
		return request, requestErr
	}

	if GitHubToken != "" {
		request.Header.Add("Authorization", fmt.Sprintf("token %s", GitHubToken))
	}

	return request, nil
}

func createGitHubPOSTRequest(url string, body interface{}) (*http.Request, error) {
	bodyJSON, marshalError := json.Marshal(body)
	if marshalError != nil {
		return nil, marshalError
	}

	bodyReader := ioutil.NopCloser(bytes.NewReader(bodyJSON))
	request, requestErr := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", GitHubAPIV3BaseURL, url), bodyReader)
	if requestErr != nil {
		return request, requestErr
	}

	if GitHubToken != "" {
		request.Header.Add("Authorization", fmt.Sprintf("token %s", GitHubToken))
	}

	return request, nil
}

func createGitHubUploadRequest(assetURL string, file *os.File) (*http.Request, error) {
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	uploadURL, _ := url.Parse(assetURL)

	queryString := uploadURL.Query()
	queryString.Add("name", filepath.Base(file.Name()))
	uploadURL.RawQuery = queryString.Encode()

	request, requestErr := http.NewRequest(http.MethodPost, uploadURL.String(), file)
	if requestErr != nil {
		return request, requestErr
	}

	mimeType := mime.TypeByExtension(filepath.Ext(file.Name()))
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	request.ContentLength = stat.Size()
	request.Header.Set("Content-Type", mimeType)

	if GitHubToken != "" {
		request.Header.Add("Authorization", fmt.Sprintf("token %s", GitHubToken))
	}

	return request, nil
}

func executeGitHubRequest(url string, body interface{}) error {
	request, requestErr := createGitHubRequest(url)
	if requestErr != nil {
		return requestErr
	}

	response, responseErr := executeRequest(request)
	if responseErr != nil {
		return responseErr
	}

	bodyData, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return readErr
	}

	return json.Unmarshal(bodyData, &body)
}

func executeRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return nil, responseErr
	}

	if response.StatusCode >= http.StatusBadRequest {
		return nil, generateErrorFrom(request.URL.String(), response)
	}

	return response, nil
}
