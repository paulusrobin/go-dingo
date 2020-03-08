package metadata

type (
	Application struct {
		Repository Repository `json:"repository"`
		Service    Service    `json:"service"`
	}

	Repository struct {
		Initialize bool   `json:"initialize"`
		Functions  []Func `json:"functions"`
	}

	Service struct {
		Initialize bool   `json:"initialize"`
		Functions  []Func `json:"functions"`
	}
)
