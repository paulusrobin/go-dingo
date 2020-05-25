package domain

import (
	"github.com/paulusrobin/go-dingo/app/domain"
	"github.com/paulusrobin/go-dingo/shared"
	"github.com/urfave/cli/v2"
)

func ExposeFunc() *cli.Command {
	return &cli.Command{
		Name:        shared.CommandExposeFunc,
		Usage:       "expose a domain function to infrastructure (http, messaging, worker)",
		Description: "expose a domain function to infrastructure (http, messaging, worker)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     shared.CommandParamDomainName,
				Usage:    "name of your domain",
				Required: true,
			},
			&cli.StringFlag{
				Name:     shared.CommandParamFuncName,
				Usage:    "name of your function",
				Required: true,
			},
			&cli.StringFlag{
				Name:     shared.CommandParamExposeTarget,
				Usage:    "expose target (http, messaging, worker)",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			domainName := ctx.String(shared.CommandParamDomainName)
			funcName := ctx.String(shared.CommandParamFuncName)
			exposeTarget := ctx.String(shared.CommandParamExposeTarget)
			return domain.ExposeFunc(domainName, funcName, exposeTarget).Run()
		},
	}
}
