package systemd

import "os/exec"

func Restart(name string) error {
	service := "olcrtc-" + name + ".service"

	cmd := exec.Command(
		"systemctl",
		"restart",
		service,
	)

	return cmd.Run()
}
