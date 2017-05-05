package main

import (
	"os/user"
	"path/filepath"

	"github.com/gomatic/servicer/usage"
	"github.com/gomatic/sshconfig"
	"github.com/urfave/cli"
)

//
func reader(op func(ctx *cli.Context, config sshconfig.Config) error) func(*cli.Context) error {
	return usage.Trapper(func(ctx *cli.Context) error {
		args := ctx.Args()
		if len(args) <= 1 {
			u, err := user.Current()
			if err != nil {
				return err
			}
			args = []string{filepath.Join(u.HomeDir, ".ssh", "config")}
		}
		config, err := sshconfig.New(args[0])
		if err != nil {
			return err
		}
		return op(ctx, config)
	})
}
