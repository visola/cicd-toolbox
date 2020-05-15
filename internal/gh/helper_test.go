package gh

import (
	"testing"

	"github.com/google/go-github/v31/github"
	"github.com/stretchr/testify/assert"
)

func TestParsesTagNameCorrectly(t *testing.T) {
	refString := "refs/tags/v1.9.5"
	url := "https://api.github.com/repos/kubernetes/kubernetes/git/commits/30e1cfa5513f14be4386ff61e82526fd3f0ab974"
	sha := "30e1cfa5513f14be4386ff61e82526fd3f0ab974"
	typeString := "commit"

	ref := &github.Reference{
		Ref: &refString,
		Object: &github.GitObject{
			URL:  &url,
			SHA:  &sha,
			Type: &typeString,
		},
	}

	assert.Equal(t, "v1.9.5", TagName(ref))
}
