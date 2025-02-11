package shell

import (
	"bytes"
	"os/exec"
)

// Cmd executes a specific linux command, and returns the output.
func Cmd(command string, args ...string) (*bytes.Buffer, error) {
	cmd := exec.Command(command, args...)
	out := new(bytes.Buffer)
	cmd.Stdout = out
	err := cmd.Run()
	return out, err
}
