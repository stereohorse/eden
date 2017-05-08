package commands

import (
	st "github.com/stereohorse/eden/storage"
)

type handleCmd func(args []string, storage *st.Storage) error

type command struct {
	handle handleCmd
	args   []string
}

func (self command) ExecuteOn(storage *st.Storage) error {
	return self.handle(self.args, storage)
}

func CommandFrom(cmdParts []string) (cmd *command) {
	if cmdParts == nil || len(cmdParts) == 0 {
		return
	}

	switch cmdParts[0] {
	case "put":
		cmd = &command{
			handle: putIntoStorage,
		}
	case "get":
		cmd = &command{
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
