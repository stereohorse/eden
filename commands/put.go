package commands

import (
	st "github.com/stereohorse/eden/storage"
	"strings"
)

func putIntoStorage(args []string, storage *st.Storage) error {
	return storage.Remember(strings.Join(args, " "))
}
