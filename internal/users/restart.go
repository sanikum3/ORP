package users

import (
	"github.com/sanikum3/ORP/internal/systemd"
)

func Restart(name string) error {
	return systemd.Restart(name)
}
