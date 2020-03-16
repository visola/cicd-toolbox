package semver

import (
	"fmt"
	"regexp"
	"strconv"
)

var semVerFormat = regexp.MustCompile("v?(\\d+)\\.(\\d+)\\.(\\d+)")

// Version represents a semantic version
type Version struct {
	Major int
	Minor int
	Patch int
}

// Parse parses a string to a Version
func Parse(toParse string) (Version, error) {
	version := Version{}

	if !semVerFormat.MatchString(toParse) {
		return version, fmt.Errorf("Invalid SemVer format: %s", toParse)
	}

	parts := semVerFormat.FindStringSubmatch(toParse)
	version.Major, _ = strconv.Atoi(parts[1])
	version.Minor, _ = strconv.Atoi(parts[2])
	version.Patch, _ = strconv.Atoi(parts[3])

	return version, nil
}
