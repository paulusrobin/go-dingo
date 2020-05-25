package cmd

import (
	"github.com/paulusrobin/go-dingo/cmd/domain"
	"github.com/paulusrobin/go-dingo/cmd/project"
	"github.com/paulusrobin/go-dingo/shared"
	"github.com/urfave/cli/v2"
)

var Commands []*cli.Command

func init() {
	Commands = append(Commands, project.Init())
	Commands = append(Commands,
		&cli.Command{
			Name:        shared.CommandDomain,
			Usage:       "manage domain",
			Description: "manage domain",
			Before: func(context *cli.Context) error {
				return shared.RequirePrograms(shared.RequiredPrograms)
			},
			Subcommands: []*cli.Command{
				domain.Init(),
				domain.AddFunc(),
			},
		},
	)
}
