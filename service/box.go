package service

import (
	"errors"

	"github.com/Wastoids/boxesandthings-api/model"
)

var (
	errInvalidUser = errors.New("invalid user")
)

type Storage interface {
	GetTopLevelBoxesForUser(userID string) ([]model.Box, error)
}

type Box struct {
	db Storage
}

func (b Box) GetTopLevelBoxesForUser(userID string) ([]model.Box, error) {
	if len(userID) == 0 {
		return nil, errInvalidUser
	}
	return b.db.GetTopLevelBoxesForUser(userID)
}

func NewBoxService(db Storage) Box {
	return Box{db: db}
}
