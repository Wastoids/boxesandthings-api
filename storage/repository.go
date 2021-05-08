package storage

import (
	"github.com/Wastoids/boxesandthings-api/model"
	"github.com/Wastoids/boxesandthings-api/service"
)

type Repository struct{}

func (r Repository) GetTopLevelBoxesForUser(userName string) ([]model.Box, error) {
	d, err := newDao()
	if err != nil {
		return nil, err
	}
	return d.getTopLevelBoxesForUser(userName)
}

func (r Repository) SaveBox(b model.Box) error {
	d, err := newDao()
	if err != nil {
		return err
	}
	return d.saveBox(b)
}

func (r Repository) SaveThing(t model.Thing, boxID string) error {
	d, err := newDao()
	if err != nil {
		return err
	}
	return d.saveThing(t, boxID)
}

func (r Repository) GetBoxContent(boxID string) (service.BoxContentResult, error) {
	d, err := newDao()
	if err != nil {
		return service.BoxContentResult{}, err
	}
	return d.getBoxContent(boxID)

}

func NewRepository() Repository {
	return Repository{}
}
