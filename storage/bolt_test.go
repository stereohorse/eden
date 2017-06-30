package storage

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func createBoltStorage(t *testing.T) *BoltStorage {
	workDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("unable to create temp work dir", err)
	}

	bolt, err := NewBoltStorage(workDir)
	if err != nil {
		t.Fatalf("unable to open bolt storage file")
	}

	return bolt
}

func TestBoltRecall(t *testing.T) {
	bolt := createBoltStorage(t)

	if err := bolt.Remember("hi there"); err != nil {
		t.Fatalf("unable to remember", err)
	}

	if err := bolt.Remember("hi you"); err != nil {
		t.Fatalf("unable to remember", err)
	}

	hits, err := bolt.Recall("there")
	if err != nil {
		t.Fatalf("unable to recall", err)
	}

	if len(hits) != 1 {
		t.Fatalf("wrong hits count")
	}
}

func TestTermsSplit(t *testing.T) {
	bolt := createBoltStorage(t)

	terms := bolt.TextToTerms("hi     theRe i       foo 1 345.456 true")
	expectedTerms := []*Term{
		&Term{
			text:    "hi",
			index:   0,
			theType: StringType,
		},
		&Term{
			text:    "there",
			index:   1,
			theType: StringType,
		},
		&Term{
			text:    "foo",
			index:   3,
			theType: StringType,
		},
		&Term{
			text:    "1",
			index:   4,
			theType: IntType,
		},
		&Term{
			text:    "345.456",
			index:   5,
			theType: FloatType,
		},
		&Term{
			text:    "true",
			index:   6,
			theType: BoolType,
		},
	}

	if !reflect.DeepEqual(expectedTerms, terms) {
		t.Fatalf("bad split")
	}
}
