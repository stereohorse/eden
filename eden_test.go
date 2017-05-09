package main

import (
	st "github.com/stereohorse/eden/storage"
	"testing"
)

func TestEdenRuns(t *testing.T) {
	storage, err := st.NewInMemoryStorage()
	if err != nil {
		t.Error(err)
	}

	run([]string{"get", "something"}, storage)
}
