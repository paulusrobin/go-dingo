package shared

import (
	"github.com/pkg/errors"
	"os"
)

func Mkdir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return errors.Wrapf(err, "failed to mkdir %s", path)
	}
	return nil
}
