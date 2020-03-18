package semrel

import (
	"github.com/VinnieApps/cicd-tools/internal/git"
	"github.com/VinnieApps/cicd-tools/internal/semver"
)

// Release represents a group of changes that will be released
type Release struct {
	Changes []Change
	Version semver.Version
}

// CalculateNextRelease will calculate what will go out in the next release,
// and if it is a major, minor or patch change.
func CalculateNextRelease(latestVersion semver.Version, commits []git.Commit) (Release, error) {
	release := Release{}

	release.Version = semver.Version{
		Major: latestVersion.Major,
		Minor: latestVersion.Minor,
		Patch: latestVersion.Patch,
	}

	major := false
	minor := false

	release.Changes = make([]Change, len(commits))
	for i, commit := range commits {
		change := CalculateChange(commit)
		major = major || change.Major
		minor = minor || change.Minor

		release.Changes[i] = change
	}

	if len(release.Changes) > 0 {
		release.Version.Patch++
	}

	if minor {
		release.Version.Minor++
		release.Version.Patch = 0
	}

	if major {
		release.Version.Major++
		release.Version.Minor = 0
		release.Version.Patch = 0
	}

	return release, nil
}
