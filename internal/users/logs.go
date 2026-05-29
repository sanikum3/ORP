package users

import (
	"github.com/sanikum3/ORP/internal/systemd"
)

func Logs(name string) error {
	return systemd.Logs(name)
}
