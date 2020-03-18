package semrel

import (
	"regexp"
	"strings"
)

var firstLineFormat = regexp.MustCompile("^(feat|feature|fix|docs|style|refactor|perf|test|chore|revert)(?:\\((.+)\\))?\\: (.*)$")

// ParsedCommitMessage represents a commit message that was parsed from
// a message in the format followed by the Semantic Release format. This
// format was adopted from the Angular.js community and you can read more
// here: https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#-git-commit-guidelines
type ParsedCommitMessage struct {
	Body    string
	Footer  string
	Scope   string
	Subject string
	Type    string
}

// ParseCommitMessage parses a commit message
func ParseCommitMessage(commitMessage string) *ParsedCommitMessage {
	// <type>(<scope>): <subject>
	// <BLANK LINE>
	// <body>
	// <BLANK LINE>
	// <footer>
	lines := strings.Split(commitMessage, "\n")

	firstLine := lines[0]

	if !firstLineFormat.MatchString(firstLine) {
		return nil
	}

	parts := firstLineFormat.FindStringSubmatch(firstLine)
	body := make([]string, 0)
	footer := make([]string, 0)

	if len(lines) > 1 {
		isFooter := false
		for i := 1; i < len(lines); i++ {
			line := lines[i]
			if line == "" {
				if len(body) == 0 {
					continue
				}
				if len(body) > 0 && !isFooter {
					isFooter = true
					continue
				}
			}

			if isFooter {
				footer = append(footer, line)
			} else {
				body = append(body, line)
			}
		}
	}

	return &ParsedCommitMessage{
		Body:    strings.Join(body, "\n"),
		Footer:  strings.Join(footer, "\n"),
		Scope:   parts[2],
		Subject: parts[3],
		Type:    parts[1],
	}
}
