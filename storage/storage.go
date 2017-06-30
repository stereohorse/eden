package storage

import (
	"fmt"
)

type Storage interface {
	Remember(string) error
	Recall(needle string) ([]Hit, error)
	Delete(*Document) error
	Close() error
	Init() error
}

type Hit struct {
	Doc Document
}

type Document struct {
	Body string
}

func (d *Document) String() string {
	return fmt.Sprintf("Document(body=%s)", d.Body)
}
