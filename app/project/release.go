package project

import (
	"github.com/paulusrobin/go-dingo/app"
	"github.com/paulusrobin/go-dingo/metadata"
)

type release struct {
	version  string
	metadata metadata.Metadata
}

func (r release) Metadata() metadata.Metadata {
	return r.metadata
}

func (r release) Run() error {
	var err error

	if err = r.metadata.Load(); err != nil {
		return err
	}

	return nil
}

func Release(version string) app.App {
	return &release{version: version}
}
