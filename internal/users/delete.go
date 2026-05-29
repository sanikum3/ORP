package users

import (
	"fmt"
	"os"

	"github.com/sanikum3/ORP/internal/storage"
	"github.com/sanikum3/ORP/internal/systemd"
)

func Delete(name string) error {
	err := systemd.Delete(name)
	if err != nil {
		return fmt.Errorf("reload systemd: %w", err)
	}
	userDir := "/opt/olcrtc-users/" + name
	delete(storage.Users, name)
	storage.Save()
	fmt.Println("User ", name, " deleted")
	return os.RemoveAll(userDir)
}
