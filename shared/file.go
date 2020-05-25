package shared

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

var (
	Structures = map[string][]string{
		Applications:         {FileInjection, FileHolder},
		Interfaces:           {FileInjection, FileHolder},
		Infrastructures:      {FileInjection, FileHolder},
		ExternalDependencies: {FileInjection},
		Shared:               {FileInjection, FileHolder},
		Config:               {FileConfig},
	}
)

func Mkdir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return errors.Wrapf(err, "failed to mkdir %s", path)
	}
	return nil
}

func WriteFile(path string, data []byte) error {
	if err := ioutil.WriteFile(path, data, 0755); err != nil {
		return errors.Wrapf(err, "failed to write %s", path)
	}
	return nil
}
