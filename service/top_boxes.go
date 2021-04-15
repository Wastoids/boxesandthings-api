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

type GetTopBoxes struct {
	db       Storage
	userName string
}

func (b GetTopBoxes) Run() (interface{}, error) {
	if len(b.userName) == 0 {
		return nil, errInvalidUser
	}
	return b.db.GetTopLevelBoxesForUser(b.userName)
}

func NewGetTopBoxesService(db Storage, userName string) GetTopBoxes {
	return GetTopBoxes{db: db, userName: userName}
}
