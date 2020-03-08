package project

import (
	"github.com/paulusrobin/go-dingo/app/project"
	"github.com/paulusrobin/go-dingo/shared"
	"github.com/urfave/cli/v2"
)

func New() *cli.Command {
	return &cli.Command{
		Name:        "init",
		Usage:       "generate new dingo project",
		Description: "generate new dingo project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "namespace",
				Usage:    "namespace of your project, example: github.com/paulusrobin",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "project",
				Usage:    "name of your project, example: go-dingo",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			if err := shared.RequirePrograms(shared.RequiredPrograms); err != nil {
				return err
			}

			namespace := ctx.String(shared.Namespace)
			projectName := ctx.String(shared.Project)
			return project.New(namespace, projectName).Run()
		},
	}
}
