package main

import (
	"errors"
	"log"
	"os"
	"path"

	cmd "github.com/stereohorse/eden/commands"
	st "github.com/stereohorse/eden/storage"
	ut "github.com/stereohorse/eden/utils"
)

func main() {
	if err := run(os.Args[1:], createFsStorage()); err != nil {
		log.Fatal(err)
	}
}

func run(args []string, storage st.Storage) error {
	logFile, err := setupLogging()
	if err != nil {
		return err
	}

	defer func() {
		if e := logFile.Close(); e != nil {
			ut.NewError("unable to close log file", err)
		}
	}()

	command := cmd.CommandFrom(args)
	if command == nil {
		return errors.New("bad command")
	}

	if err = command.ExecuteOn(storage); err != nil {
		return err
	}

	return nil
}

func setupLogging() (*os.File, error) {
	wp, err := st.GetDefaultWorkDirPath()
	if err != nil {
		return nil, ut.NewError("unable to get work dir", err)
	}

	f, err := os.OpenFile(path.Join(wp, "logs"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, ut.NewError("unable to open log file", err)
	}

	log.SetOutput(f)

	return f, nil
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
