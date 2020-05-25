package shared

const (
	Namespace            = "namespace"
	Project              = "project"
	Applications         = "applications"
	Interfaces           = "interfaces"
	Infrastructures      = "infrastructures"
	ExternalDependencies = "external-dependencies"
	Shared               = "shared"
	Config               = "shared/config"

	FileConfig    = "config.go"
	FileInjection = "di.go"
	FileHolder    = "holder.go"
)

var RequiredPrograms = []string{"go"}
