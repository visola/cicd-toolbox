package github

import (
	"regexp"
	"strings"
)

var linkFormat = regexp.MustCompile("<(.+)>; rel=\"(.+)\"")

type Link struct {
	Relation string
	URL      string
}

func ParseLinks(headerValue string) []Link {
	links := make([]Link, 0)

	if headerValue == "" {
		return links
	}

	linkTexts := strings.Split(headerValue, ",")
	for _, linkText := range linkTexts {
		parts := linkFormat.FindStringSubmatch(linkText)
		links = append(links, Link{URL: parts[0], Relation: parts[1]})
	}

	return links
}
