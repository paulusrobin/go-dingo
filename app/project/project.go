package project

import (
	"github.com/paulusrobin/go-dingo/app"
	"github.com/paulusrobin/go-dingo/metadata"
	"github.com/paulusrobin/go-dingo/shared"
)

type project struct {
	namespace, project string
	metadata           metadata.Metadata
}

func (p project) Metadata() metadata.Metadata {
	return p.metadata
}

func (p project) Run() error {
	var err error

	// - Initialize Metadata
	p.metadata, err = metadata.NewMetadata(p.namespace, p.project)
	if err != nil {
		return err
	}

	// - Generate Project Folder
	if err = shared.Mkdir(p.project); err != nil {
		return err
	}

	// - Generate Project File
	// TODO: Generate Project File

	// - Generate .metadata.json file
	return p.metadata.Save(p.project)
}

func NewProject(namespace, projectName string) app.App {
	return &project{namespace: namespace, project: projectName}
}
