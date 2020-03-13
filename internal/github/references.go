package github

import "strings"

type Reference struct {
	Object    Object `json:"object"`
	Reference string `json:"ref"`
}

func (ref Reference) TagName() string {
	if strings.HasPrefix(ref.Reference, REFERENCE_TAG_PREFIX) {
		return ref.Reference[len(REFERENCE_TAG_PREFIX):]
	}

	return ""
}
