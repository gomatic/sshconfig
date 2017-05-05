package main

import (
	"fmt"

	"github.com/gomatic/sshconfig"
	"github.com/urfave/cli"
)

//
func command_query(ctx *cli.Context, config sshconfig.Config) error {
	args := ctx.Args()
	q := ""
	if len(args) > 0 {
		q = args[0]
	}
	fmt.Print(config.Find(sshconfig.NewQuery(q)))
	return nil
}
