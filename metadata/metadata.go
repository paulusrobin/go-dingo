package metadata

type (
	Metadata struct {
		Namespace string    `json:"namespace"`
		Project   string    `json:"project"`
		Version   string    `json:"version"`
		Domains   []Domain  `json:"domains"`
		Utilities []Utility `json:"utilities"`
	}
)
