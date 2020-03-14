package github

import "strings"

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
