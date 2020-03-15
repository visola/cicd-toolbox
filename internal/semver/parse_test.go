package semver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	validVersions := []string{
		"0.0.0",
		"0.2.13",
		"1.0.0",
		"431.123.9421",
	}

	expectedVersions := [][]int{
		[]int{0, 0, 0},
		[]int{0, 2, 13},
		[]int{1, 0, 0},
		[]int{431, 123, 9421},
	}

	for _, prefix := range []string{"", "v"} {
		for index, validVersion := range validVersions {
			version, err := Parse(prefix + validVersion)
			require.Nil(t, err, "Should parse "+validVersion+" correctly")

			expectedVersion := expectedVersions[index]
			assert.Equal(t, version.Major, expectedVersion[0])
			assert.Equal(t, version.Minor, expectedVersion[1])
			assert.Equal(t, version.Patch, expectedVersion[2])
		}
	}
}
