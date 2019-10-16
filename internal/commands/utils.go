package commands

import "regexp"

func IsVersionNo(line string) bool {
	versionStr := regexp.MustCompile(`(^|\s)\"version\": .*`)
	return versionStr.MatchString(line)
}

func GetVersion(line string) string {
	v := regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)?`)
	return v.FindString(line)
}

func ReplaceVersion(line, version string) string {
	v := regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)?`)
	return v.ReplaceAllLiteralString(line, version)
}
