package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	Version = "unknown version"
)

func main() {
	app := cli.NewApp()
	app.Name = "ddl2plantuml"
	app.Description = "ddl2plantuml is a tool to generate plantuml diagram from database ddl."
	app.Version = Version

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
