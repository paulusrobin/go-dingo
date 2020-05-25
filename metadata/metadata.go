package metadata

import (
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/paulusrobin/go-dingo/shared"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

type (
	Metadata struct {
		Namespace string    `json:"namespace"`
		Project   string    `json:"project"`
		Version   string    `json:"version"`
		Domains   []Domain  `json:"domains"`
		Utilities []Utility `json:"utilities"`
	}
)

func (m Metadata) Save(path string) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal metadata")
	}

	path = path + "/.metadata.json"
	if err := shared.WriteFile(path, data); err != nil {
		return err
	}
	return nil
}

func (m *Metadata) Load(path string) error {
	path = path + "/.metadata.json"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "could not read metadata.json")
	}

	if err := json.Unmarshal(data, m); err != nil {
		return errors.Wrap(err, "failed to unmarshal metadata")
	}
	return nil
}

func NewMetadata(namespace, project string) (Metadata, error) {
	if len(namespace) == 0 {
		return Metadata{}, errors.New(`namespace can't be blank`)
	}
	if len(project) == 0 {
		return Metadata{}, errors.New(`project can't be blank'`)
	}

	namespace = strings.ToLower(strings.Replace(namespace, " ", "", -1))
	project = strings.ToLower(strcase.ToKebab(project))
	return Metadata{
		Version:   "v0.0.0",
		Namespace: namespace,
		Project:   project,
		Domains:   make([]Domain, 0),
		Utilities: make([]Utility, 0),
	}, nil
}
