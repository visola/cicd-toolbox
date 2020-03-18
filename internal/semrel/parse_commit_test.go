package semrel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseIgnoresWrongFormat(t *testing.T) {
	commitMessage := "FEAT-1031: This is not a correctly formatted semantic release message"
	parsedCommit := ParseCommitMessage(commitMessage)

	require.Nil(t, parsedCommit)
}

func TestParseOneLineCommitWithScope(t *testing.T) {
	commitMessage := "feat(build): some important thing"
	parsedCommit := ParseCommitMessage(commitMessage)

	require.NotNil(t, parsedCommit)
	assert.Equal(t, "feat", parsedCommit.Type)
	assert.Equal(t, "some important thing", parsedCommit.Subject)
	assert.Equal(t, "build", parsedCommit.Scope)
}

func TestParseOneLineCommitWithoutScope(t *testing.T) {
	commitMessage := "feat: some important thing"
	parsedCommit := ParseCommitMessage(commitMessage)

	require.NotNil(t, parsedCommit)
	assert.Equal(t, "feat", parsedCommit.Type)
	assert.Equal(t, "some important thing", parsedCommit.Subject)
	assert.Empty(t, parsedCommit.Scope)
}

func TestParseCommitWithBody(t *testing.T) {
	commitMessage := `feature(build): some important thing

This is an important change because there was a bug
that needed to be fixed. Without it customers get mad at us.
`

	body := `This is an important change because there was a bug
that needed to be fixed. Without it customers get mad at us.`

	parsedCommit := ParseCommitMessage(commitMessage)

	require.NotNil(t, parsedCommit)
	assert.Equal(t, "feature", parsedCommit.Type)
	assert.Equal(t, "some important thing", parsedCommit.Subject)
	assert.Equal(t, "build", parsedCommit.Scope)
	assert.Equal(t, body, parsedCommit.Body)
	assert.Empty(t, parsedCommit.Footer)
}

func TestParseCommitWithBodyAndFooter(t *testing.T) {
	commitMessage := `feat(build): some important thing

This is an important change because there was a bug
that needed to be fixed. Without it customers get mad at us.

BREAKING CHANGE: Method signature changed.

Before you called it like: importantFunction(argument1)

After this change: importantFunction(argument1, anotherArgument)`

	body := `This is an important change because there was a bug
that needed to be fixed. Without it customers get mad at us.`

	footer := `BREAKING CHANGE: Method signature changed.

Before you called it like: importantFunction(argument1)

After this change: importantFunction(argument1, anotherArgument)`

	parsedCommit := ParseCommitMessage(commitMessage)

	require.NotNil(t, parsedCommit)
	assert.Equal(t, "feat", parsedCommit.Type)
	assert.Equal(t, "some important thing", parsedCommit.Subject)
	assert.Equal(t, "build", parsedCommit.Scope)
	assert.Equal(t, body, parsedCommit.Body)
	assert.Equal(t, footer, parsedCommit.Footer)
}
