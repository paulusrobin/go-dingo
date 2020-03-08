package file_system

import "github.com/gobuffalo/packr/v2"

type (
	FileSystem interface {
		Find(filename string) ([]byte, error)
	}

	fileSystem struct {
		box *packr.Box
	}
)

func (f *fileSystem) Find(filename string) ([]byte, error) {
	buf, err := f.box.Find(filename)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func NewFileSystem(name, path string) FileSystem {
	return &fileSystem{packr.New(name, path)}
}
