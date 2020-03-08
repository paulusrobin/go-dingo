package shared

import (
	file_system "github.com/paulusrobin/go-dingo/util/file-system"
	"html/template"
	"os"
)

type (
	RenderDto struct {
		PackageName         string
		StrippedPackageName string
	}
)

var (
	FileSystem file_system.FileSystem
)

func init() {
	FileSystem = file_system.NewFileSystem("template", "template")
}

func RenderFromTemplate(templatePath, targetPath string, data RenderDto) error {
	file, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	templateFile, err := FileSystem.Find(templatePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New("template").Funcs(template.FuncMap{}).Parse(string(templateFile))
	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, data); err != nil {
		return err
	}
	return nil
}
