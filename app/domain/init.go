package domain

import (
	"github.com/paulusrobin/go-dingo/app"
	"github.com/paulusrobin/go-dingo/metadata"
)

type domain struct {
	domainName string
	metadata   metadata.Metadata
}

func (a domain) Metadata() metadata.Metadata {
	return a.metadata
}

func (a domain) Run() error {
	var err error

	if err = a.metadata.Load(); err != nil {
		return err
	}

	return nil
}

func Init(domainName string) app.App {
	return &domain{domainName: domainName}
}
