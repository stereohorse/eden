package commands

import (
	"fmt"
	st "github.com/stereohorse/eden/storage"
	u "github.com/stereohorse/eden/utils"
)

func getFromStorage(args []string, storage st.Storage) error {
	ui := NewConsoleUI(storage)

	if err := ui.Run(); err != nil {
		return u.NewError("unable to run UI", err)
	}

	searchResult := ui.GetSearchResult()
	if searchResult != nil {
		fmt.Println(searchResult.Body)
	}

	return nil
}
