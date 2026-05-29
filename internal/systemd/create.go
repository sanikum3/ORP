package systemd

import (
	"os"
	"text/template"

	"github.com/sanikum3/ORP/internal/templates"
)

type ServiceData struct {
	Username string
}

func Create(username string) error {

	servicePath :=
		"/etc/systemd/system/olcrtc-" +
			username +
			".service"

	content, err := templates.FS.ReadFile("service.tmpl")
	if err != nil {
		return err
	}
	tmpl, err := template.New("service").Parse(string(content))
	file, err := os.Create(servicePath)
	if err != nil {
		return err
	}

	defer file.Close()

	data := ServiceData{
		Username: username,
	}

	return tmpl.Execute(file, data)
}
