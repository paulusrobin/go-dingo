package metadata

import (
	"encoding/json"
	"fmt"
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

func (m Metadata) Render(path string) error {
	for packageName, files := range shared.Structures {
		if err := shared.Mkdir(path + "/" + packageName); err != nil {
			return err
		}

		for _, file := range files {
			templatePath := fmt.Sprintf("%s.tmpl", file)
			pathToRender := fmt.Sprintf("%s/%s/%s", path, packageName, file)
			if err := shared.RenderFromTemplate(templatePath, pathToRender, shared.RenderDto{
				PackageName:         packageName,
				StrippedPackageName: shared.GetStrippedPackageName(packageName),
			}); err != nil {
				return err
			}
		}
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
		Version:   "0.0.0-1",
		Namespace: namespace,
		Project:   project,
		Domains:   make([]Domain, 0),
		Utilities: make([]Utility, 0),
	}, nil
}
