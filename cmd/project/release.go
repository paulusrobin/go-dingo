package project

import (
	"github.com/paulusrobin/go-dingo/app/project"
	"github.com/paulusrobin/go-dingo/shared"
	"github.com/urfave/cli/v2"
)

func Release() *cli.Command {
	return &cli.Command{
		Name:        shared.CommandRelease,
		Usage:       "release new version of dingo project",
		Description: "release new version of dingo project",
		Before: func(context *cli.Context) error {
			return shared.RequirePrograms(shared.RequiredPrograms)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     shared.CommandParamVersion,
				Usage:    "version number, example: v1.0.0",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			version := ctx.String(shared.CommandParamVersion)
			return project.Release(version).Run()
		},
	}
}
