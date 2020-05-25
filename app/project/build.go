package project

import (
	"github.com/paulusrobin/go-dingo/app"
	"github.com/paulusrobin/go-dingo/metadata"
)

type build struct {
	metadata metadata.Metadata
}

func (b build) Metadata() metadata.Metadata {
	return b.metadata
}

func (b build) Run() error {
	var err error

	if err = b.metadata.Load(); err != nil {
		return err
	}

	return nil
}

func Build() app.App {
	return &build{}
}
