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

func (b Box) GetTopLevelBoxesForUser(userName string) ([]model.Box, error) {
	if len(userName) == 0 {
		return nil, errInvalidUser
	}
	return b.db.GetTopLevelBoxesForUser(userName)
}

func NewBoxService(db Storage) Box {
	return Box{db: db}
}
