package systemd

import (
	"os"
	"text/template"
)

type ServiceData struct {
	Username string
}

func Create(username string) error {

	servicePath :=
		"/etc/systemd/system/olcrtc-" +
			username +
			".service"

	tmpl, err := template.ParseFiles(
		"templates/service.tmpl",
	)
	if err != nil {
		return err
	}

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
