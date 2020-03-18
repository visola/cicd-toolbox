package semrel

import (
	"testing"

	"github.com/VinnieApps/cicd-tools/internal/git"
	"github.com/stretchr/testify/assert"
)

func TestCalculateChangeInvalidMessage(t *testing.T) {
	commit := git.Commit{
		Message: "Not a valid semantic release message",
		SHA:     "1234abcd1234",
	}

	change := CalculateChange(commit)
	assert.Equal(t, commit, change.Commit)
	assert.False(t, change.Minor)
	assert.False(t, change.Major)
}

func TestCalculatePatchChange(t *testing.T) {
	commit := git.Commit{
		Message: "fix: build was broken",
		SHA:     "1234abcd1234",
	}

	change := CalculateChange(commit)
	assert.Equal(t, commit, change.Commit)
	assert.False(t, change.Minor)
	assert.False(t, change.Major)
}

func TestCalculateMinorChange(t *testing.T) {
	commit := git.Commit{
		Message: "feat: this is a new feature",
		SHA:     "1234abcd1234",
	}

	change := CalculateChange(commit)
	assert.Equal(t, commit, change.Commit)
	assert.True(t, change.Minor)
	assert.False(t, change.Major)
}

func TestCalculateMajorChange(t *testing.T) {
	commit := git.Commit{
		Message: `feat: this is a new feature

This is an important change.

BREAKING CHANGE:
Now the return value is going to be encrypted.
`,
		SHA: "1234abcd1234",
	}

	change := CalculateChange(commit)
	assert.Equal(t, commit, change.Commit)
	assert.True(t, change.Minor)
	assert.True(t, change.Major)
}
