package github

import (
	"regexp"
	"strings"
)

var linkFormat = regexp.MustCompile("<(.+)>; rel=\"(.+)\"")

// Link represents a link that came in a `Link` header from a GitHub API call
type Link struct {
	Relation string
	URL      string
}

// ParseLinks parses all links from a `Link` header that came from a GitHub API call
func ParseLinks(headerValue string) []Link {
	links := make([]Link, 0)

	if headerValue == "" {
		return links
	}

	linkTexts := strings.Split(headerValue, ",")
	for _, linkText := range linkTexts {
		parts := linkFormat.FindStringSubmatch(linkText)
		links = append(links, Link{URL: parts[1], Relation: parts[2]})
	}

	return links
}