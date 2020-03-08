package cmd

import (
	"github.com/paulusrobin/go-dingo/cmd/project"
	"github.com/urfave/cli/v2"
)

var Commands []*cli.Command

func init() {
	Commands = append(Commands, project.New())
}
