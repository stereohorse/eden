package commands

import (
	"testing"
)

func TestParseCommand(t *testing.T) {
	command := CommandFrom(nil)
	if command != nil {
		t.Error("should return nil if input is nil")
	}

	command = CommandFrom([]string{"arg0"})
	if command != nil {
		t.Error("should return nil if command is unknown")
	}

	command = CommandFrom([]string{"put", "something"})
	if command == nil {
		t.Error("should select right command, expected 'put', received nil")
	}

	if len(command.args) != 1 {
		t.Error("should omit command in args list")
	}

	if command.args[0] != "something" {
		t.Error("should put valid arg")
	}
}
