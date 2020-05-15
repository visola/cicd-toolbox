package gh

import (
	"strings"

	"github.com/google/go-github/v31/github"
)

// ReferenceTagPrefix References that have this prefix are tags
const ReferenceTagPrefix = "refs/tags/"

// TagName returns the tag name if the reference is a tag, otherwise, returns empty
func TagName(ref *github.Reference) string {
	if strings.HasPrefix(*ref.Ref, ReferenceTagPrefix) {
		return (*ref.Ref)[len(ReferenceTagPrefix):]
	}

	return ""
}

// ToOwnerRepo transforms a GitHub slug into owner and repo strings
// "VinnieApps/cicd-toolbox" -> "VinnieApps", "cicd-toolbox"
func ToOwnerRepo(gitHubSlug string) (string, string) {
	split := strings.Split(gitHubSlug, "/")
	return split[0], split[1]
}
