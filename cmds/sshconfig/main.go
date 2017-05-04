package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/gomatic/sshconfig"
	"github.com/urfave/cli"
)

//
var tag, version = "0.1", "0"

//
func main() {
	app := cli.NewApp()
	app.Name = "sshconfig"
	app.Usage = "sshconfig."
	app.Version = tag + "." + version
	app.EnableBashCompletion = true

	app.Action = func(ctx *cli.Context) (err error) {
		args := ctx.Args()
		if len(args) != 1 {
			u, err := user.Current()
			if err != nil {
				return err
			}
			args = []string{filepath.Join(u.HomeDir, ".ssh", "config")}
		}
		config := sshconfig.MustNew(args[0])
		fmt.Print(config)
		return nil
	}

	app.Run(os.Args)
}
