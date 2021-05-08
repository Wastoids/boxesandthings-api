package service

import (
	"encoding/json"
	"errors"

	"github.com/Wastoids/boxesandthings-api/model"
)

var (
	nameRequiredForThing = errors.New("name is required for a thing")
)

type SaveThing struct {
	db    Storage
	t     model.Thing
	boxID string
}

func NewSaveThing(db Storage, t model.Thing, boxID string) SaveThing {
	return SaveThing{db: db, t: t}
}

func (s SaveThing) Run() (interface{}, error) {
	if len(s.t.Name) == 0 {
		return nil, nameRequiredForThing
	}
	return nil, s.db.SaveThing(s.t, s.boxID)
}

func GetThingFromRequest(body string) model.Thing {
	var req struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return model.Thing{}
	}
	return model.Thing{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
	}
}
