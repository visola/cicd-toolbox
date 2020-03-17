package semrel

import (
	"regexp"

	"github.com/VinnieApps/cicd-tools/internal/github"
)

var breakingChangePattern = regexp.MustCompile("(BREAKING CHANGE)?")

// Change represents a change that will go out
type Change struct {
	Commit        github.Commit
	Major         bool
	Minor         bool
	ParsedMessage ParsedCommitMessage
}

// CalculateChange calculates a changed out of a commit
func CalculateChange(commit github.Commit) Change {
	parsedMessage := ParseCommitMessage(commit.Commit.Message)

	return Change{
		Commit:        commit,
		Major:         breakingChangePattern.MatchString(commit.Commit.Message),
		Minor:         parsedMessage.Type == "feat" || parsedMessage.Type == "feature",
		ParsedMessage: *parsedMessage,
	}
}
