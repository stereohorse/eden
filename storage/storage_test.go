package storage

import (
	"os"
	"testing"
)

func TestGetWorkDirPath(t *testing.T) {
	workDir := getWorkDirPath()
	if len(workDir) == 0 {
		t.Error("should return work dir path")
	}
}

func TestCreateWorkDir(t *testing.T) {
	GetStorage()

	if !pathExists(getWorkDirPath(), t) {
		t.Error("should create work dir")
	}
}

func TestGetIndexPath(t *testing.T) {
	indexPath := getIndexPath()
	if len(indexPath) == 0 {
		t.Error("should construct index path")
	}
}

func pathExists(path string, t *testing.T) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	t.Error("unable to stat path " + path)
	return false
}
