package service

import (
	"errors"
)

var (
	errInvalidBoxID = errors.New("invalid box")
)

type BoxContent struct {
	db    Storage
	boxID string
}

func NewBoxContent(db Storage, boxID string) BoxContent {
	return BoxContent{db: db, boxID: boxID}
}

func (b BoxContent) Run() (interface{}, error) {
	if len(b.boxID) == 0 {
		return nil, errInvalidBoxID
	}
	return b.db.GetBoxContent(b.boxID)
}
