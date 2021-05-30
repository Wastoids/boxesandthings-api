package service

import (
	"encoding/json"
	"errors"

	"github.com/Wastoids/boxesandthings-api/model"
	"github.com/google/uuid"
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
	s.b.ID = uuid.NewString()
	return nil, s.db.SaveBox(s.b)
}

func NewSaveBox(db Storage, b model.Box) SaveBox {
	return SaveBox{b: b, db: db}
}

func GetBoxFromRequest(body string) model.Box {
	var req struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return model.Box{}
	}
	return model.Box{Name: req.Name}
}
