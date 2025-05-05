package files

import (
	"github.com/AlexLuminare/read_advisor_bot/lib/e"
	"github.com/AlexLuminare/read_advisor_bot/storage"
	"os"
	"path/filepath"
)

const (
	defaultPerm = 0774
)

type Storage struct {
	basePath string
}

func NewStorage(basePath string) *Storage {
	return &Storage{basePath}
}

func (s *Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can't save page", err) }()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err = os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
