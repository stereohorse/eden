package storage

import (
	"strings"
)

type InMemoryStorage struct {
	docs []Document
}

func (self *InMemoryStorage) Remember(text string) error {
	self.docs = append(self.docs, Document{
		Body: text,
	})
	return nil
}

func (self *InMemoryStorage) Recall(needle string) (hits []Hit, err error) {
	for _, doc := range self.docs {
		if strings.Contains(doc.Body, needle) {
			hits = append(hits, Hit{
				Doc: doc,
			})
		}
	}

	return
}

func (st *InMemoryStorage) Delete(doc *Document) error {
	return nil
}

func (st *InMemoryStorage) Close() error {
	return nil
}

func (st *InMemoryStorage) Init() error {
	return nil
}

func NewInMemoryStorage() (*InMemoryStorage, error) {
	return &InMemoryStorage{}, nil
}

func (self *InMemoryStorage) rememberDoc(doc *Document) {
	self.docs = append(self.docs, *doc)
}
