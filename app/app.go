package app

import "github.com/paulusrobin/go-dingo/metadata"

type App interface {
	Run() error
	Metadata() metadata.Metadata
}
