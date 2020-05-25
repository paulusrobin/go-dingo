package domain

import (
	"github.com/paulusrobin/go-dingo/app/domain"
	"github.com/paulusrobin/go-dingo/shared"
	"github.com/urfave/cli/v2"
)

func AddFunc() *cli.Command {
	return &cli.Command{
		Name:        shared.CommandAddFunc,
		Usage:       "add new func to domain",
		Description: "add new func to domain",
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
		},
		Action: func(ctx *cli.Context) error {
			domainName := ctx.String(shared.CommandParamDomainName)
			funcName := ctx.String(shared.CommandParamFuncName)
			return domain.AddFunc(domainName, funcName).Run()
		},
	}
}
