package metadata

type (
	Domain struct {
		Name           string         `json:"name"`
		Application    Application    `json:"application"`
		Interface      Interface      `json:"interface"`
		Infrastructure Infrastructure `json:"infrastructure"`
	}
)
