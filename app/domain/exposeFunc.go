package domain

import (
	"github.com/paulusrobin/go-dingo/app"
	"github.com/paulusrobin/go-dingo/metadata"
)

type exposeFunc struct {
	domainName, funcName, exposeTarget string
	metadata                           metadata.Metadata
}

func (e exposeFunc) Metadata() metadata.Metadata {
	return e.metadata
}

func (e exposeFunc) Run() error {
	var err error

	if err = e.metadata.Load(); err != nil {
		return err
	}

	return nil
}

func ExposeFunc(domainName, funcName, exposeTarget string) app.App {
	return &exposeFunc{domainName: domainName, funcName: funcName, exposeTarget: exposeTarget}
}
