package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithArchitectures(t *testing.T) {
	testPlatforms := PlatformList{
		Platform{Architecture: "386", Extension: "", OperatingSystem: "one"},
		Platform{Architecture: "amd64", Extension: "", OperatingSystem: "one"},
		Platform{Architecture: "arm64", Extension: "arm", OperatingSystem: "one"},

		Platform{Architecture: "386", Extension: "", OperatingSystem: "two"},
		Platform{Architecture: "amd64", Extension: "", OperatingSystem: "two"},
		Platform{Architecture: "arm64", Extension: "", OperatingSystem: "two"},
	}

	armPlatforms := testPlatforms.WithArchitectures("arm64")

	require.Equal(t, 2, len(armPlatforms))

	first := armPlatforms[0]
	assert.Equal(t, "arm64", first.Architecture)
	assert.Equal(t, "arm", first.Extension)
	assert.Equal(t, "one", first.OperatingSystem)

	second := armPlatforms[1]
	assert.Equal(t, "arm64", second.Architecture)
	assert.Equal(t, "", second.Extension)
	assert.Equal(t, "two", second.OperatingSystem)
}

func TestWithOperatingSystems(t *testing.T) {
	testPlatforms := PlatformList{
		Platform{Architecture: "386", Extension: "", OperatingSystem: "one"},
		Platform{Architecture: "amd64", Extension: "", OperatingSystem: "one"},
		Platform{Architecture: "arm64", Extension: "arm", OperatingSystem: "one"},

		Platform{Architecture: "386", Extension: "exe", OperatingSystem: "two"},
		Platform{Architecture: "arm64", Extension: "", OperatingSystem: "two"},
	}

	twoPlatforms := testPlatforms.WithOperatingSystems("two")

	require.Equal(t, 2, len(twoPlatforms))

	first := twoPlatforms[0]
	assert.Equal(t, "386", first.Architecture)
	assert.Equal(t, "exe", first.Extension)
	assert.Equal(t, "two", first.OperatingSystem)

	second := twoPlatforms[1]
	assert.Equal(t, "arm64", second.Architecture)
	assert.Equal(t, "", second.Extension)
	assert.Equal(t, "two", second.OperatingSystem)
}
