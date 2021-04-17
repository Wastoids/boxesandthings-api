package service

import (
	"errors"
)

var (
	errInvalidUser = errors.New("invalid user")
)

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
