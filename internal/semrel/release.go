package semrel

import (
	"fmt"
	"strings"
	"time"

	"github.com/VinnieApps/cicd-toolbox/internal/semver"
)

// Release represents a group of changes that will be released
type Release struct {
	Changes []Change
	Version semver.Version
}

// ChangeLog returns a string representing the change log for the
// release, in markdown format
func (release Release) ChangeLog() string {
	now := time.Now()
	changeLog := fmt.Sprintf("## %s (%s)", release.Version, now.Format("2006-01-02"))

	for commitType, changes := range release.ChangesByType() {
		changeLog = changeLog + fmt.Sprintf("\n\n#### %s%s\n", strings.ToUpper(commitType[0:1]), commitType[1:])

		for _, change := range changes {
			changeLog = changeLog + fmt.Sprintf("* %s (%s)\n", change.ParsedMessage.Subject, change.Commit.ShortSHA())
		}
	}

	return changeLog
}

// ChangesByType group the changes for the release by their type
func (release Release) ChangesByType() map[string][]Change {
	changesByType := make(map[string][]Change)
	for _, change := range release.Changes {
		if change.ParsedMessage.Subject == "" {
			continue
		}

		commitType := change.ParsedMessage.Type
		if commitType == "feat" {
			commitType = "feature"
		}

		messages, exist := changesByType[commitType]
		if !exist {
			messages = make([]Change, 0)
		}
		changesByType[commitType] = append(messages, change)
	}

	return changesByType
}
