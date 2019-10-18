package git

import (
	"os/exec"

	"github.com/urfave/cli"
)

const DefaultCommitMsg string = "bump version"

func CommitToGit(c *cli.Context) error {
	if !c.GlobalBool("commit") {
		return nil
	}

	path := c.GlobalString("file")

	gitAddCmd := exec.Command("git", "add", path)
	if err := gitAddCmd.Run(); err != nil {
		return err
	}

	gitCommitCmd := exec.Command("git", "commit", "-m", DefaultCommitMsg)
	if err := gitCommitCmd.Run(); err != nil {
		return err
	}

	return nil
}
