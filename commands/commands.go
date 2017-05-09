package commands

import (
	st "github.com/stereohorse/eden/storage"
)

type Command struct {
	handle handleCmd
	args   []string
}

func (self Command) ExecuteOn(storage st.Storage) error {
	return self.handle(self.args, storage)
}

func CommandFrom(cmdParts []string) (cmd *Command) {
	if cmdParts == nil || len(cmdParts) == 0 {
		return
	}

	switch cmdParts[0] {
	case "put":
		cmd = &Command{
			handle: putIntoStorage,
		}
	case "get":
		cmd = &Command{
			handle: getFromStorage,
		}
	default:
		return
	}

	if len(cmdParts) > 1 {
		cmd.args = cmdParts[1:]
	}

	return
}

type handleCmd func(args []string, storage st.Storage) error
