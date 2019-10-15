package internal

import (
	"github.com/urfave/cli"
)

func addCommands(app *cli.App) {
	commands := []cli.Command{}
	app.Commands = commands
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
