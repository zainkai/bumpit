package commands

import (
	"regexp"

	"github.com/coreos/go-semver/semver"
)

// IsVersionNo checks if a string contains npm like version string
func IsVersionNo(line string) bool {
	versionStr := regexp.MustCompile(`(^|\s)\"version\": .*`)
	return versionStr.MatchString(line)
}

// GetVersion retrieves the version number from a string
func GetVersion(line string) string {
	v := regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)?`)
	return v.FindString(line)
}

// ReplaceVersion replaces the verion of within a string
// input:  "version": "1.1.1" ", "2.0.0"
// output: "version": "2.00"
func ReplaceVersion(line, version string) string {
	v := regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)?`)
	return v.ReplaceAllLiteralString(line, version)
}

// SemverBump a hashmap containing update functions for semver
var SemverBump = map[UPDATE_SEMVER]func(*semver.Version){
	PATCH: func(v *semver.Version) { v.BumpPatch() },
	MINOR: func(v *semver.Version) { v.BumpMinor() },
	MAJOR: func(v *semver.Version) { v.BumpMajor() },
}
