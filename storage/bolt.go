package storage

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/boltdb/bolt"
	u "github.com/stereohorse/eden/utils"
)

var (
	termsBucket       = []byte("terms")
	docsBucket        = []byte("docs")
	termsToDocsBucket = []byte("termsToDocs")

	termSplitRegex = regexp.MustCompile("\\s+")
)

type Term struct {
	text    string
	theType int
	index   int64
}

const (
	StringType = iota
	IntType
	FloatType
	BoolType
)

func NewTerm(text string, index int64) *Term {
	theType := StringType

	if _, err := strconv.ParseInt(text, 10, 64); err == nil {
		theType = IntType
	} else if _, err = strconv.ParseFloat(text, 64); err == nil {
		theType = FloatType
	} else if _, err = strconv.ParseBool(text); err == nil {
		theType = BoolType
	}

	return &Term{
		text:    text,
		index:   index,
		theType: theType,
	}
}

type BoltStorage struct {
	db *bolt.DB

	termsBucket       *bolt.Bucket
	docsBucket        *bolt.Bucket
	termsToDocsBucket *bolt.Bucket
}

func (bs *BoltStorage) Remember(text string) (err error) {
	err = bs.db.Update(func(tx *bolt.Tx) error {
		return nil
	})

	return err
}

func (bs *BoltStorage) TextToTerms(text string) (terms []*Term) {
	parts := termSplitRegex.Split(text, -1)
	if len(parts) == 0 {
		return nil
	}

	for i, part := range parts {
		part = bs.NormalizePart(part)
		term := NewTerm(part, int64(i))
		if bs.FilterPart(term) {
			terms = append(terms, term)
		}
	}

	if len(terms) == 0 {
		return nil
	}

	return terms
}

func (bs *BoltStorage) FilterPart(term *Term) bool {
	if term.theType == StringType {
		return len(term.text) > 1
	}

	return true
}

func (bs *BoltStorage) NormalizePart(part string) string {
	return strings.ToLower(part)
}

func (bs *BoltStorage) Recall(needle string) ([]Hit, error) {
	return nil, nil
}

func (bs *BoltStorage) Delete(doc *Document) error {
	return nil
}

func (bs *BoltStorage) Close() error {
	if err := bs.db.Close(); err != nil {
		return u.NewError("unable to close db", err)
	}

	return nil
}

func (bs *BoltStorage) Init() error {
	err := bs.db.Update(func(tx *bolt.Tx) error {
		termsBucket, err := initBucket(tx, termsBucket)
		if err != nil {
			return err
		}
		bs.termsBucket = termsBucket

		docsBucket, err := initBucket(tx, docsBucket)
		if err != nil {
			return err
		}
		bs.docsBucket = docsBucket

		termsToDocsBucket, err := initBucket(tx, termsToDocsBucket)
		if err != nil {
			return err
		}
		bs.termsToDocsBucket = termsToDocsBucket

		return nil
	})

	return err
}

func initBucket(tx *bolt.Tx, name []byte) (*bolt.Bucket, error) {
	b, err := tx.CreateBucketIfNotExists(name)
	if err != nil {
		return nil, fmt.Errorf("unable to create bucket [%s]: %s", name, err)
	}

	return b, nil
}

func NewBoltStorage(workDir string) (*BoltStorage, error) {
	// TODO: add timeouts
	db, err := bolt.Open(path.Join(workDir, "storage.bolt"), 0600, nil)
	if err != nil {
		return nil, err
	}

	return &BoltStorage{
		db: db,
	}, nil
}
