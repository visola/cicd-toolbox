package semrel

import (
	"fmt"
	"regexp"

	"github.com/VinnieApps/cicd-toolbox/internal/git"
)

var breakingChangePattern = regexp.MustCompile("BREAKING CHANGE?")

// Change represents a change that will go out
type Change struct {
	Commit        git.Commit
	Major         bool
	Minor         bool
	ParsedMessage ParsedCommitMessage
}

func (c Change) String() string {
	return fmt.Sprintf("Change %s is Major: %t, is Minor: %t, type: '%s'", c.Commit.SHA, c.Major, c.Minor, c.ParsedMessage.Type)
}

// CalculateChange calculates a changed out of a commit
func CalculateChange(commit git.Commit) Change {
	parsedMessage := ParseCommitMessage(commit.Message)

	if parsedMessage == nil {
		return Change{Commit: commit}
	}

	return Change{
		Commit:        commit,
		Major:         breakingChangePattern.MatchString(commit.Message),
		Minor:         parsedMessage.Type == "feat" || parsedMessage.Type == "feature",
		ParsedMessage: *parsedMessage,
	}
}
