package service

import (
	"errors"

	"github.com/Wastoids/boxesandthings-api/model"
)

var (
	errCannotSaveEmptyBox = errors.New("cannot save an empty box")
)

type SaveBox struct {
	b  model.Box
	db Storage
}

func (s SaveBox) Run() (interface{}, error) {
	if s.b.Equals(model.Box{}) {
		return nil, errCannotSaveEmptyBox
	}

	return nil, s.db.SaveBox(s.b)
}
