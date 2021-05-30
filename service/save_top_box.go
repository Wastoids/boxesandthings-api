package service

import (
	"errors"

	"github.com/Wastoids/boxesandthings-api/model"
	"github.com/google/uuid"
)

var (
	errEmptyBox = errors.New("invalid box")
)

type SaveTopBox struct {
	b  model.Box
	db Storage
	u  string
}

func NewSaveTopBox(db Storage, userName string, b model.Box) SaveTopBox {
	return SaveTopBox{
		b:  b,
		db: db,
		u:  userName,
	}
}

func (s SaveTopBox) Run() (interface{}, error) {
	if s.b.Equals(model.Box{}) {
		return nil, errEmptyBox
	}
	s.b.ID = uuid.NewString()
	return s.b, s.db.SaveTopBox(s.u, s.b)
}
