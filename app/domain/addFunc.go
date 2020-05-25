package domain

import (
	"github.com/paulusrobin/go-dingo/app"
	"github.com/paulusrobin/go-dingo/metadata"
)

type addFunc struct {
	domainName, funcName string
	metadata             metadata.Metadata
}

func (a addFunc) Metadata() metadata.Metadata {
	return a.metadata
}

func (a addFunc) Run() error {
	var err error

	if err = a.metadata.Load(); err != nil {
		return err
	}

	return nil
}

func AddFunc(domainName, funcName string) app.App {
	return &addFunc{domainName: domainName, funcName: funcName}
}
