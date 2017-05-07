package main

import (
	cmd "github.com/stereohorse/eden/commands"
	st "github.com/stereohorse/eden/storage"
	"os"
)

func main() {
	command := cmd.CommandFrom(os.Args)
	if command == nil {
		os.Exit(1)
	}

	storage, err := st.GetStorage()
	if err != nil {
		os.Exit(1)
	}

	if err := command.ExecuteOn(storage); err != nil {
		os.Exit(1)
	}
}
