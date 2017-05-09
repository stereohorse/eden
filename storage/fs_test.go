package storage

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestWorkDir(t *testing.T) {
	path, err := GetDefaultWorkDirPath()
	if err != nil {
		t.Error(err)
	}

	if path == "" {
		t.Error("default work dir path is null")
	}
}

func TestCreateWorkDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Error(err)
	}

	workDir := path.Join(dir, "some", "dir", "for", "work")
	_, err = NewFsStorage(workDir)
	if err != nil {
		t.Error(err)
	}

	dirExists, err := pathExists(workDir)
	if err != nil {
		t.Error(err)
	}

	if !dirExists {
		t.Error("should create work dir if not exists")
	}

	storageExists, err := pathExists(path.Join(workDir, storageName))
	if err != nil {
		t.Error(err)
	}

	if !storageExists {
		t.Error("should create storage")
	}
}

func TestRecall(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}

	workDir := path.Join(dir, "some", "dir", "for", "work")

	fss, err := NewFsStorage(workDir)
	if err != nil {
		t.Fatal(err)
	}

	records := []string{
		"hi", "hello", "chao",
	}

	for _, r := range records {
		err = fss.Remember(r)
		if err != nil {
			t.Fatal(err)
		}
	}

	if len(fss.docs) != len(records) {
		t.Fatal("should put exact number of records")
	}

	hits, err := fss.Recall("hi")
	if err != nil {
		t.Fatal(err)
	}

	if len(hits) != 1 {
		t.Fatal("should recall exactly one hit")
	}

	if hits[0].Doc.Body != "hi" {
		t.Fatal("should recall exact document")
	}

	fss, err = NewFsStorage(workDir)
	if err != nil {
		t.Fatal(err)
	}

	if len(fss.docs) != len(records) {
		t.Fatalf("should load all docs, actually: %d", len(fss.docs))
	}

	for _, d := range fss.docs {
		t.Log(d)
	}

	hits, err = fss.Recall("hi")
	if err != nil {
		t.Fatal(err)
	}

	if len(hits) != 1 {
		t.Fatalf("should recall exactly 1 hit after reload, actual: %d", len(hits))
	}

	if hits[0].Doc.Body != "hi" {
		t.Error("should recall exact hit after reload")
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, errors.New("unable to check file existence")
}
