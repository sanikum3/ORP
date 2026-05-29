package systemd

import "os/exec"

func Enable(name string) error {
	service := "olcrtc-" + name + ".service"

	cmd := exec.Command(
		"systemctl",
		"enable",
		service,
	)

	return cmd.Run()
}
