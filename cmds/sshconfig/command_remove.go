package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gomatic/sshconfig"
	"github.com/urfave/cli"
)

//
func command_remove(ctx *cli.Context, config sshconfig.Config) error {
	fmt.Print(config)
	log.Println("removed", time.Now())
	return nil
}
