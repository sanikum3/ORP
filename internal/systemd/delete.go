package systemd

import (
	"fmt"
	"os"
	"os/exec"
)

func Delete(name string) error {

	service := "olcrtc-" +
		name +
		".service"

	cmd := exec.Command(
		"systemctl",
		"stop",
		service,
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println("stop error:", err)
	}

	cmd = exec.Command(
		"systemctl",
		"disable",
		service,
	)

	err = cmd.Run()
	if err != nil {
		fmt.Println("disable error:", err)
	}

	servicePath :=
		"/etc/systemd/system/" +
			service

	err = os.Remove(servicePath)
	if err != nil {
		return err
	}

	cmd = exec.Command(
		"systemctl",
		"daemon-reload",
	)

	return cmd.Run()
}
