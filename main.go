package main

import (
	"fmt"
	"os"

	"github.com/thecasualcoder/pg-ping/cmd"
	"github.com/urfave/cli"
)

var version string

const defaultVersion = "dev"

func main() {
	var app = cli.NewApp()
	if version == "" {
		version = defaultVersion
	}
	app.Version = version

	cli.HelpFlag = cli.BoolFlag{Name: "help"}

	if err := cmd.Execute(app); err != nil {
		fmt.Fprintf(os.Stderr, "pg-ping failed: %v\n", err)
		os.Exit(1)
	}
}
