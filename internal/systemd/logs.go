package systemd

import (
	"os"
	"os/exec"
)

func Logs(name string) error {

	service := "olcrtc-" + name + ".service"

	cmd := exec.Command("journalctl", "-u", service, "-n", "100")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
