package project

import (
	"github.com/paulusrobin/go-dingo/app/project"
	"github.com/paulusrobin/go-dingo/shared"
	"github.com/urfave/cli/v2"
)

func Build() *cli.Command {
	return &cli.Command{
		Name:        shared.CommandBuild,
		Usage:       "build and run dingo project",
		Description: "build and run dingo project",
		Before: func(context *cli.Context) error {
			return shared.RequirePrograms(shared.RequiredPrograms)
		},
		Action: func(ctx *cli.Context) error {
			return project.Build().Run()
		},
	}
}
