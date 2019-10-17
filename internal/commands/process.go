package commands

import (
	"github.com/coreos/go-semver/semver"
	"github.com/urfave/cli"
	"github.com/zainkai/bumpit/pkg/utils"
)

func ProcessVersionLine(line string, update UPDATE_SEMVER) string {
	vStr := GetVersion(line)
	s, _ := semver.NewVersion(vStr)

	SemverBump[update](s)

	return ReplaceVersion(line, s.String())
}

func Process(c *cli.Context, res Resouces) error {
	e, err := utils.NewEditor(res.TargetFile)
	if err != nil {
		return err
	}

	e.EditLine(func(i int, line string) string {
		if IsVersionNo(line) {
			return ProcessVersionLine(line, res.Update) + "\n"
		}
		return line + "\n"
	})

	if err := e.CommitFile(); err != nil {
		return err
	}

	return nil
}
