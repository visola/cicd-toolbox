package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyHeaderReturnsNoLink(t *testing.T) {
	links := ParseLinks("")
	assert.Empty(t, links)
}

func TestParseLinks(t *testing.T) {
	linksHeader := `<https://api.github.com/user/repos?page=3&per_page=100>; rel="next",
	<https://api.github.com/user/repos?page=50&per_page=100>; rel="last"`

	links := ParseLinks(linksHeader)
	assert.Equal(t, 2, len(links))

	assert.Equal(t, "https://api.github.com/user/repos?page=3&per_page=100", links[0].URL)
	assert.Equal(t, "next", links[0].Relation)

	assert.Equal(t, "https://api.github.com/user/repos?page=50&per_page=100", links[1].URL)
	assert.Equal(t, "last", links[1].Relation)
}
