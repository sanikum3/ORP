package storage

import (
	"encoding/json"
	"os"

	"github.com/sanikum3/ORP/internal/models"
)

var Users = map[string]models.User{}

const File = "/opt/orp/data/users.json"

func Save() error {

	err := os.MkdirAll("/opt/orp/data", 0755)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(
		Users,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		File,
		data,
		0644,
	)
}

func Load() error {

	data, err := os.ReadFile(File)
	if err != nil {
		return nil
	}

	return json.Unmarshal(data, &Users)
}
