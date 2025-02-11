package shell

import (
	"testing"
)

func TestCmd(t *testing.T) {
	out, err := Cmd("ls", "-alh")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("$> ls -alh\n%s", out.String())
}
