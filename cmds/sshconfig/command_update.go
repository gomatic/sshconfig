package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gomatic/sshconfig"
	"github.com/urfave/cli"
)

//
func command_update(ctx *cli.Context, config sshconfig.Config) error {
	fmt.Print(config)
	log.Println("update", time.Now())
	return nil
}
