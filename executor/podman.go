package executor

import (
	"os"
	"os/exec"
)

func createContainer(name string) error {
	workspace := `C:\Users\sahil\projects\monci\workspace`
	volume := workspace + `:/workspace`

	cmd := exec.Command("podman", "create", "--name", name, "-v", volume, "alpine", "sleep", "infinity")
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func startContainer(name string) error {
	cmd := exec.Command("podman", "start", name)
	return cmd.Run()
}

func stopContainer(name string) error {
	cmd := exec.Command("podman", "rm", name)
	return cmd.Run()
}

func execCommand(name string, command string) error {
	cmd := exec.Command(
		"podman",
		"exec",
		name,
		"sh",
		"-c",
		command,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
