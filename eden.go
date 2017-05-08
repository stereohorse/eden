package main

import (
	"log"
	"os"

	cmd "github.com/stereohorse/eden/commands"
	st "github.com/stereohorse/eden/storage"
)

func main() {
	command := cmd.CommandFrom(os.Args[1:])
	if command == nil {
		log.Fatal("bad command")
	}

	storage := st.GetStorage()

	if err := command.ExecuteOn(storage); err != nil {
		log.Fatal("unable to execute command")
	}
}
