package semrel

import (
	"testing"

	"github.com/VinnieApps/cicd-toolbox/internal/git"
	"github.com/VinnieApps/cicd-toolbox/internal/semver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculateNextReleaseReturnsSameIfNoCommits(t *testing.T) {
	release, err := CalculateNextRelease(semver.Version{}, make([]git.Commit, 0))

	require.Nil(t, err)
	assert.Equal(t, "0.0.0", release.Version.String())
}

func TestCalculateNextReleaseAlwaysReturnNextPatch(t *testing.T) {
	commits := []git.Commit{
		git.Commit{
			Message: "chore: this is a test",
			SHA:     "1234abcd1234",
		},
		git.Commit{
			Message: "Merge PR Bla Bla Bla",
			SHA:     "1234abcd1234",
		},
		git.Commit{
			Message: "refactor: move code from here to there",
			SHA:     "1234abcd1234",
		},
	}
	release, err := CalculateNextRelease(semver.Version{}, commits)

	require.Nil(t, err)
	assert.Equal(t, "0.0.1", release.Version.String())
}

func TestCalculateNextReleaseReturnNextMinorIfNewFeature(t *testing.T) {
	commits := []git.Commit{
		git.Commit{
			Message: "feature: this is a test",
			SHA:     "1234abcd1234",
		},
		git.Commit{
			Message: "Merge PR Bla Bla Bla",
			SHA:     "1234abcd1234",
		},
		git.Commit{
			Message: "feat: move code from here to there",
			SHA:     "1234abcd1234",
		},
	}
	release, err := CalculateNextRelease(semver.Version{}, commits)

	require.Nil(t, err)
	assert.Equal(t, "0.1.0", release.Version.String())
}

func TestCalculateNextReleaseReturnNextMajorWhenBreakingChange(t *testing.T) {
	commits := []git.Commit{
		git.Commit{
			Message: `feat: this is a new feature

This is an important change.

BREAKING CHANGE:
Now the return value is going to be encrypted.
`,
			SHA: "1234abcd1234",
		},
		git.Commit{
			Message: "Merge PR Bla Bla Bla",
			SHA:     "1234abcd1234",
		},
		git.Commit{
			Message: "feat: move code from here to there",
			SHA:     "1234abcd1234",
		},
	}
	release, err := CalculateNextRelease(semver.Version{}, commits)

	require.Nil(t, err)
	assert.Equal(t, "1.0.0", release.Version.String())
}
