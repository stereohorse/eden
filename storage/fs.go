package storage

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
	u "github.com/stereohorse/eden/utils"
	"io"
	"os"
	"os/user"
	"path"
)

const (
	storageName = "storage"
)

var (
	storageBytesOrder = binary.LittleEndian
)

type FsStorage struct {
	InMemoryStorage

	workDirPath string
}

func (self *FsStorage) storagePath() string {
	return path.Join(self.workDirPath, storageName)
}

func (self *FsStorage) Remember(text string) (err error) {
	err = self.InMemoryStorage.Remember(text)
	if err != nil {
		return u.NewError("unable to remember doc", err)
	}

	var docBytes bytes.Buffer
	enc := gob.NewEncoder(&docBytes)
	err = enc.Encode(Document{
		Body: text,
	})
	if err != nil {
		return u.NewError("unable to encode doc", err)
	}

	f, err := os.OpenFile(self.storagePath(),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return u.NewError("unable to open storage", err)
	}

	defer func() {
		if e := f.Close(); e != nil {
			err = u.NewError("unable to close storage", e)
		}
	}()

	docBytesLen := uint64(docBytes.Len())

	bf := bufio.NewWriter(f)
	err = binary.Write(bf, storageBytesOrder, &docBytesLen)
	if err != nil {
		return u.NewError("unable to write doc header", err)
	}

	_, err = bf.Write(docBytes.Bytes())
	if err != nil {
		return u.NewError("unable to write doc", err)
	}

	if err = bf.Flush(); err != nil {
		return u.NewError("unable to flush docs", err)
	}

	return nil
}

func NewFsStorage(dirPath string) (*FsStorage, error) {
	fss := &FsStorage{
		workDirPath: dirPath,
	}

	if err := os.MkdirAll(dirPath, 0700); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(fss.storagePath(),
		os.O_CREATE|os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	defer func() {
		if e := f.Close(); e != nil {
			fss = nil
			err = e
		}
	}()

	bf := bufio.NewReader(f)
	for {
		var docBytesLen uint64
		err = binary.Read(bf, storageBytesOrder, &docBytesLen)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, u.NewError("unable to read doc header", err)
		}

		docBytes := make([]byte, docBytesLen)
		n, err := bf.Read(docBytes)
		if err != nil {
			return nil, u.NewError("unable to read doc", err)
		}
		if uint64(n) != docBytesLen {
			return nil, errors.New("unexpected EOF")
		}

		doc := &Document{}

		dec := gob.NewDecoder(bytes.NewBuffer(docBytes))
		err = dec.Decode(doc)
		if err != nil {
			return nil, u.NewError("unable to decode docs", err)
		}

		fss.InMemoryStorage.rememberDoc(doc)
	}

	return fss, nil
}

func GetDefaultWorkDirPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return path.Join(usr.HomeDir, ".eden"), nil
}
