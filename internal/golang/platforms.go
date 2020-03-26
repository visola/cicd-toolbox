package golang

import "github.com/VinnieApps/cicd-toolbox/internal/arrayutil"

// Platform represents a platform that binaries can be compiled for
type Platform struct {
	Architecture    string
	Extension       string
	OperatingSystem string
}

// PlatformList represent a list of available platforms
type PlatformList []Platform

// Platforms represents all the platforms that packages will compile
// for by default
var Platforms = PlatformList{
	Platform{Architecture: "386", Extension: "", OperatingSystem: "darwin"},
	Platform{Architecture: "amd64", Extension: "", OperatingSystem: "darwin"},
	Platform{Architecture: "arm", Extension: "", OperatingSystem: "darwin"},
	Platform{Architecture: "arm64", Extension: "", OperatingSystem: "darwin"},

	Platform{Architecture: "386", Extension: "", OperatingSystem: "freebsd"},
	Platform{Architecture: "amd64", Extension: "", OperatingSystem: "freebsd"},

	Platform{Architecture: "385", Extension: "", OperatingSystem: "linux"},
	Platform{Architecture: "amd64", Extension: "", OperatingSystem: "linux"},
	Platform{Architecture: "arm", Extension: "", OperatingSystem: "linux"},
	Platform{Architecture: "arm64", Extension: "", OperatingSystem: "linux"},

	Platform{Architecture: "386", Extension: "", OperatingSystem: "openbsd"},
	Platform{Architecture: "amd64", Extension: "", OperatingSystem: "openbsd"},

	Platform{Architecture: "386", Extension: "exe", OperatingSystem: "windows"},
	Platform{Architecture: "amd64", Extension: "exe", OperatingSystem: "windows"},
}

// WithArchitectures filters the list of available platform to only include those with the
// specified architecture
func (platforms PlatformList) WithArchitectures(architectures ...string) PlatformList {
	if len(architectures) == 0 {
		return platforms
	}

	filteredPlatforms := make(PlatformList, 0)
	for _, platform := range platforms {
		if arrayutil.ContainsString(architectures, platform.Architecture) {
			filteredPlatforms = append(filteredPlatforms, platform)
		}
	}
	platforms = filteredPlatforms

	return filteredPlatforms
}

// WithOperatingSystems filters the list of available platform to only include those with the
// specified operating system
func (platforms PlatformList) WithOperatingSystems(operatingSystems ...string) PlatformList {
	if len(operatingSystems) == 0 {
		return platforms
	}

	filteredPlatforms := make(PlatformList, 0)
	for _, platform := range platforms {
		if arrayutil.ContainsString(operatingSystems, platform.OperatingSystem) {
			filteredPlatforms = append(filteredPlatforms, platform)
		}
	}
	platforms = filteredPlatforms

	return filteredPlatforms
}
