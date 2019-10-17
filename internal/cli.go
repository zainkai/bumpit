package internal

import (
	"github.com/urfave/cli"
	"github.com/zainkai/bumpit/internal/commands"
)

const DefaultFile string = "./package.json"

func addCommands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "Patch",
			Aliases: []string{"p", "patch"},
			Usage:   "Update patch version",
			Action: func(c *cli.Context) error {
				return commands.Process(c, commands.Resouces{
					c.GlobalString("file"),
					commands.PATCH,
				})
			},
		}, {
			Name:    "Minor",
			Aliases: []string{"mi", "minor"},
			Usage:   "Update minor version",
			Action: func(c *cli.Context) error {
				return commands.Process(c, commands.Resouces{
					c.GlobalString("file"),
					commands.MINOR,
				})
			},
		}, {
			Name:    "Major",
			Aliases: []string{"ma", "major"},
			Usage:   "Update major version",
			Action: func(c *cli.Context) error {
				return commands.Process(c, commands.Resouces{
					c.GlobalString("file"),
					commands.MAJOR,
				})
			},
		},
	}
}

func addFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f, path, p",
			Value: DefaultFile,
			Usage: "target file",
		},
		cli.BoolFlag{
			Name:  "commit, c",
			Usage: "commits version to git",
		},
	}
}

func addInfo(app *cli.App) {
	app.Name = "bumpit"
	app.Description = "Auto update Major/Minor/Patch of any JavaScript project!"
	app.Author = "Kevin Turkington (Zainkai)"
	app.Version = "1.0.0"
}

// InitApp Initializes CLI application
func InitApp() *cli.App {
	app := cli.NewApp()
	addInfo(app)
	addCommands(app)
	addFlags(app)
	addAfter(app)

	return app
}
