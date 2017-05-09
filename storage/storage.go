package storage

import (
	"fmt"
)

type Storage interface {
	Remember(string) error
	Recall(needle string) ([]Hit, error)
}

type Hit struct {
	Doc Document
}

type Document struct {
	Body string
}

func (self *Document) String() string {
	return fmt.Sprintf("Document(body=%s)", self.Body)
}
