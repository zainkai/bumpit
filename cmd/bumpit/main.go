package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/zainkai/bumpit/pkg/utils"
)

func isVersionNo(line string) bool {
	versionStr := regexp.MustCompile(`(^|\s)\"version\": .*`)
	return versionStr.MatchString(line)
}

func getVersion(line string) string {
	v := regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)?`)
	return v.FindString(line)
}

func main() {
	e, err := utils.NewEditor("./package.json")
	if err != nil {
		log.Fatalf("Could not create file editor", err)
	}

	e.EditLine(func(i int, line string) string {
		if isVersionNo(line) {
			return "FOUND IT\n"
		}
		return fmt.Sprintf("%d:  %s\n", i, line)
	})

	if err := e.CommitFile(); err != nil {
		log.Fatalf("Could not commit changes to file", err)
	}
}
