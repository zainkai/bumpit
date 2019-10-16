package internal

import (
	"github.com/urfave/cli"
	"github.com/zainkai/bumpit/internal/commands"
)

func addCommands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "Patch",
			Aliases: []string{"p", "patch"},
			Usage:   "Update patch version",
			Action: func(c *cli.Context) error {
				return commands.Process(c, commands.PATCH)
			},
		}, {
			Name:    "Minor",
			Aliases: []string{"mi", "minor"},
			Usage:   "Update minor version",
			Action: func(c *cli.Context) error {
				return commands.Process(c, commands.MINOR)
			},
		}, {
			Name:    "Major",
			Aliases: []string{"ma", "major"},
			Usage:   "Update minor version",
			Action: func(c *cli.Context) error {
				return commands.Process(c, commands.MAJOR)
			},
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

	return app
}
