package storage

import "github.com/Wastoids/boxesandthings-api/model"

type Repository struct{}

func (r Repository) GetTopLevelBoxesForUser(userName string) ([]model.Box, error) {
	d, err := newDao()
	if err != nil {
		return nil, err
	}
	return d.getTopLevelBoxesForUser(userName)
}

func NewRepository() Repository {
	return Repository{}
}
