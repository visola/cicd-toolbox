package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortSHA(t *testing.T) {
	commit := Commit{
		Message: "Some Message",
		SHA:     "92901dd395bc8854dab596d2a4cb3bc19757930e",
	}

	assert.Equal(t, "92901dd3", commit.ShortSHA())
}
