package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsesTagNameCorrectly(t *testing.T) {
	ref := Reference{
		Reference: "refs/tags/v1.9.5",
		Object: Object{
			URL:  "https://api.github.com/repos/kubernetes/kubernetes/git/commits/30e1cfa5513f14be4386ff61e82526fd3f0ab974",
			SHA:  "30e1cfa5513f14be4386ff61e82526fd3f0ab974",
			Type: "commit",
		},
	}

	assert.Equal(t, "v1.9.5", ref.TagName())
}
