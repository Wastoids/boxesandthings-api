package service

import "github.com/Wastoids/boxesandthings-api/model"

type Storage interface {
	GetTopLevelBoxesForUser(userID string) ([]model.Box, error)
	SaveBox(b model.Box) error
	SaveThing(t model.Thing) error
}
