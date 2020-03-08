package shared

import (
	"github.com/iancoleman/strcase"
	"strings"
)

func GetStrippedPackageName(packageName string) string {
	splits := strings.Split(packageName, "/")
	response := splits[len(splits)-1]
	response = ToSnakeCase(response)
	return response
}

func ToKebabCase(data string) string {
	return strcase.ToKebab(data)
}

func ToSnakeCase(data string) string {
	return strcase.ToSnake(data)
}