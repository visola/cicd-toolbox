package gh

import (
	"context"
	"net/http"

	"github.com/google/go-github/v31/github"
	"golang.org/x/oauth2"
)

var (
	// GitHubToken stores the token to be used to authenticate with GitHub
	GitHubToken string

	gitHubClient *github.Client
)

// Client get or create the GitHub client
func Client() *github.Client {
	if gitHubClient == nil {
		var oauthClient *http.Client
		if GitHubToken != "" {
			tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: GitHubToken})
			oauthClient = oauth2.NewClient(context.Background(), tokenSource)
		}
		gitHubClient = github.NewClient(oauthClient)
	}

	return gitHubClient
}
