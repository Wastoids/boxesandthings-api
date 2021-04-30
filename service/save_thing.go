package service

import "github.com/Wastoids/boxesandthings-api/storage"

type SaveThing struct {
	db storage.Repository
}

func NewSaveThing(db storage.Repository) SaveThing {
	return SaveThing{db: db}
}

func (s SaveThing) Run() (interface{}, error) {
	return nil, nil
}
