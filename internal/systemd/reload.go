package systemd

import "os/exec"

func Reload() error {

	cmd := exec.Command(
		"systemctl",
		"daemon-reload",
	)

	return cmd.Run()
}
