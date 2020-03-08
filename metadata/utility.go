package metadata

type (
	Utility struct {
		Name                 string                         `json:"name"`
		Import               string                         `json:"import"`
		ImportAs             string                         `json:"import_as"`
		ProviderFunc         string                         `json:"provider_func"`
		Config               []*UtilityConfig               `json:"config"`
		ModDependencies      []string                       `json:"mod_dependencies"`
		ExternalDependencies []Dependency                   `json:"external_dependencies"`
		Install              func(metadata *Metadata) error `json:"-"`
		Uninstall            func(metadata *Metadata) error `json:"-"`
	}

	Dependency struct {
		Import   string `json:"import"`
		ImportAs string `json:"import_as"`
	}

	UtilityConfig struct {
		Key       string `json:"key"`
		Type      string `json:"type"`
		Value     string `json:"value"`
		StructTag string `json:"struct_tag"`
		EnvKey    string `json:"env_key"`
	}
)
