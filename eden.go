package main

import (
	"errors"
	"log"
	"os"

	cmd "github.com/stereohorse/eden/commands"
	st "github.com/stereohorse/eden/storage"
)

func main() {
	if err := run(os.Args[1:], createFsStorage()); err != nil {
		log.Fatal(err)
	}
}

func run(args []string, storage st.Storage) error {
	command := cmd.CommandFrom(args)
	if command == nil {
		return errors.New("bad command")
	}

	if err := command.ExecuteOn(storage); err != nil {
		return err
	}

	return nil
}

func createFsStorage() st.Storage {
	workDirPath, err := st.GetDefaultWorkDirPath()
	if err != nil {
		log.Fatal(err)
	}

	storage, err := st.NewFsStorage(workDirPath)
	if err != nil {
		log.Fatal(err)
	}

	return storage
}
