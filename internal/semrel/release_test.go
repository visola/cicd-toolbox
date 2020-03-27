package semrel

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/VinnieApps/cicd-toolbox/internal/git"
	"github.com/VinnieApps/cicd-toolbox/internal/semver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChangeLog(t *testing.T) {
	messagesByType := map[string][]string{
		"feat":     []string{"this is a new feature", "another feature"},
		"feature":  []string{"one more feature", "i love features"},
		"chore":    []string{"cleaning up", "fix build"},
		"refactor": []string{"moving some code around"},
	}

	commits := make([]git.Commit, 0)
	for commitType, messages := range messagesByType {
		for _, message := range messages {
			commits = append(commits, git.Commit{
				Message: fmt.Sprintf("%s: %s", commitType, message),
				SHA:     "1234abcd4321",
			})
		}
	}
	release, err := CalculateNextRelease(semver.Version{}, commits)
	require.Nil(t, err)

	expectedParts := []string{
		"", "", "", "", "", "",
		fmt.Sprintf("## 0.1.0 (%s)", time.Now().Format("2006-01-02")),
		"#### Refactor", "* moving some code around (1234abcd)",
		"#### Feature", "* this is a new feature (1234abcd)", "* another feature (1234abcd)", "* one more feature (1234abcd)", "* i love features (1234abcd)",
		"#### Chore", "* cleaning up (1234abcd)", "* fix build (1234abcd)",
	}
	sort.Strings(expectedParts)

	parts := strings.Split(release.ChangeLog(), "\n")
	sort.Strings(parts)

	assert.Equal(t, expectedParts, parts)
}

func TestChangesByType(t *testing.T) {
	messagesByType := map[string][]string{
		"feat":     []string{"this is a new feature", "another feature"},
		"feature":  []string{"one more feature", "i love features"},
		"chore":    []string{"cleaning up", "fix build"},
		"refactor": []string{"moving some code around"},
	}

	commits := make([]git.Commit, 0)
	for commitType, messages := range messagesByType {
		for _, message := range messages {
			commits = append(commits, git.Commit{
				Message: fmt.Sprintf("%s: %s", commitType, message),
				SHA:     "1234abcd4321",
			})
		}
	}
	release, err := CalculateNextRelease(semver.Version{}, commits)
	require.Nil(t, err)

	changesByType := release.ChangesByType()
	assert.Equal(t, 3, len(changesByType))

	expectedTypes := []string{"chore", "feature", "refactor"}
	expectedMessagesByType := map[string][]string{
		"chore":    messagesByType["chore"],
		"feature":  append(messagesByType["feat"], messagesByType["feature"]...),
		"refactor": messagesByType["refactor"],
	}

	for _, expectedType := range expectedTypes {
		changes, exists := changesByType[expectedType]
		assert.True(t, exists, "Expected type to exist in result: %s", expectedType)

		messages := make([]string, len(changes))
		for i, change := range changes {
			messages[i] = change.ParsedMessage.Subject
		}

		sort.Strings(messages)

		expectedMessages := expectedMessagesByType[expectedType]
		sort.Strings(expectedMessages)

		assert.Equal(t, expectedMessages, messages)
	}
}
