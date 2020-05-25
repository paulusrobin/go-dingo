package project

import (
	"github.com/paulusrobin/go-dingo/app/project"
	"github.com/paulusrobin/go-dingo/shared"
	"github.com/urfave/cli/v2"
)

func Init() *cli.Command {
	return &cli.Command{
		Name:        shared.CommandInit,
		Usage:       "generate new dingo project",
		Description: "generate new dingo project",
		Before: func(context *cli.Context) error {
			return shared.RequirePrograms(shared.RequiredPrograms)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     shared.CommandParamNamespace,
				Usage:    "namespace of your project, example: github.com/paulusrobin",
				Required: true,
			},
			&cli.StringFlag{
				Name:     shared.CommandParamProject,
				Usage:    "name of your project, example: go-dingo",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			namespace := ctx.String(shared.CommandParamNamespace)
			projectName := ctx.String(shared.CommandParamProject)
			return project.Init(namespace, projectName).Run()
		},
	}
}
