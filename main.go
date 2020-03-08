package main

import (
	"github.com/paulusrobin/go-dingo/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{Commands: cmd.Commands}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
