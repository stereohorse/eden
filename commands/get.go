package commands

import (
	st "github.com/stereohorse/eden/storage"
	"strings"
)

func getFromStorage(args []string, storage *st.Storage) error {
	return storage.Recall(strings.Join(args, " "))
}
