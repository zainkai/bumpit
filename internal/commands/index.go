package commands

import (
	"github.com/coreos/go-semver/semver"
	"github.com/urfave/cli"
	"github.com/zainkai/bumpit/pkg/utils"
)

func processVersionLine(line string, update UPDATE_SEMVER) string {
	vStr := GetVersion(line)
	s, _ := semver.NewVersion(vStr)

	SemverBump[update](s)

	return ReplaceVersion(line, s.String())
}

func Process(c *cli.Context, update UPDATE_SEMVER) error {
	e, err := utils.NewEditor("./package.json")
	if err != nil {
		return err
	}

	e.EditLine(func(i int, line string) string {
		if IsVersionNo(line) {
			return processVersionLine(line, update) + "\n"
		}
		return line + "\n"
	})

	if err := e.CommitFile(); err != nil {
		return err
	}

	return nil
}
