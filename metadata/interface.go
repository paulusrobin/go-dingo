package metadata

type (
	Interface struct {
		Initialize bool   `json:"initialize"`
		Functions  []Func `json:"functions"`
	}
)
