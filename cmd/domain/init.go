package domain

import (
	"github.com/paulusrobin/go-dingo/app/domain"
	"github.com/paulusrobin/go-dingo/shared"
	"github.com/urfave/cli/v2"
)

func Init() *cli.Command {
	return &cli.Command{
		Name:        shared.CommandInit,
		Usage:       "initialize a new domain",
		Description: "initialize a new domain",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     shared.CommandParamDomainName,
				Usage:    "name of your domain",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			domainName := ctx.String(shared.CommandParamDomainName)
			return domain.Init(domainName).Run()
		},
	}
}
