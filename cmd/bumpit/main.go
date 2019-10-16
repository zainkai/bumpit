package main

import (
	"log"
	"os"

	"github.com/zainkai/bumpit/internal"
)

func main() {
	app := internal.InitApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
