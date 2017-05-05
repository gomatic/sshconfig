package main

import (
	"os"

	"github.com/gomatic/usage"
	"github.com/urfave/cli"
)

//
var tag, version = "0.1", "0"

var settings struct {
	DryRun bool
}

//
func main() {
	app := cli.NewApp()
	app.Name = "sshconfig"
	app.Usage = "sshconfig."
	app.Version = tag + "." + version
	app.EnableBashCompletion = true
	app.OnUsageError = usage.Error

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "dry-run, dryrun, n",
			Usage:       "Do not make any changes",
			Destination: &settings.DryRun,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "update",
			Aliases:   []string{"u"},
			Usage:     "Update host attribute.",
			ArgsUsage: "host key value",

			Action: reader(command_update),
		},
		{
			Name:      "query",
			Aliases:   []string{"q", "show"},
			Usage:     "Query the config.",
			ArgsUsage: "query",

			Action: reader(command_query),
		},
		{
			Name:      "remove",
			Usage:     "Remove host attribute key.",
			ArgsUsage: "query",

			Action: reader(command_remove),
		},
	}

	app.Run(os.Args)
}
