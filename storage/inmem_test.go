package storage

import (
	"testing"
)

func TestInMemoryStorage(t *testing.T) {
	storage, err := NewInMemoryStorage()
	if err != nil {
		t.Error(err)
	}

	err = storage.Remember("hi")
	if err != nil {
		t.Error(err)
	}
	err = storage.Remember("hello")
	if err != nil {
		t.Error(err)
	}

	hits, err := storage.Recall("hi")
	if err != nil {
		t.Error(err)
	}
	if hits == nil {
		t.Error("no hits")
	}
	if len(hits) != 1 {
		t.Error("should be exactly 1 hit")
	}
}
