package shared

import "github.com/iancoleman/strcase"

func ToKebabCase(data string) string {
	return strcase.ToKebab(data)
}
