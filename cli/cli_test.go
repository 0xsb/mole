package cli_test

import (
	"testing"

	"github.com/davrodpin/mole/cli"
)

func TestHandleArgs(t *testing.T) {

	tests := []struct {
		args     []string
		expected string
	}{
		{
			[]string{"./mole", "-version"},
			"version",
		},
		{
			[]string{"./mole", "-help"},
			"help",
		},
		{
			[]string{"./mole", "-remote", ":443", "-server", "example1"},
			"start",
		},
		{
			[]string{"./mole", "-alias", "xyz", "-remote", ":443", "-server", "example1"},
			"new-alias",
		},
		{
			[]string{"./mole", "-alias", "xyz", "-delete"},
			"rm-alias",
		},
		{
			[]string{"./mole", "-start", "example1-alias"},
			"start-from-alias",
		},
	}

	var c *cli.App

	for _, test := range tests {
		c = cli.New(test.args)
		c.Parse()
		if test.expected != c.Command {
			t.Errorf("test failed. Expected: %s, value: %s", test.expected, c.Command)
		}
	}
}
