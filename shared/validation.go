package shared

import (
	"github.com/pkg/errors"
	"os/exec"
)

func RequirePrograms(programs []string) error {
	for _, program := range programs {
		if err := exec.Command("which", program).Run(); err != nil {
			return errors.Wrapf(err, "%s is not installed, please install %s first", program, program)
		}
	}
	return nil
}
