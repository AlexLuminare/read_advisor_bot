package storage

import (
	"crypto/sha1"
	"fmt"
	"github.com/AlexLuminare/read_advisor_bot/lib/e"
	"io"
)

// методы взаимодействия с хранилищим
type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
	//Created  time.Time
}

func (p Page) Hash() (string, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
