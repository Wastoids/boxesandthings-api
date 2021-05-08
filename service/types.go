package service

import (
	"reflect"

	"github.com/Wastoids/boxesandthings-api/model"
)

type Storage interface {
	GetTopLevelBoxesForUser(userID string) ([]model.Box, error)
	SaveBox(b model.Box) error
	SaveThing(t model.Thing, boxID string) error
	GetBoxContent(boxID string) (BoxContentResult, error)
}

type BoxContentResult struct {
	Boxes  []model.Box   `json:"box"`
	Things []model.Thing `json:"thing"`
}

func (b BoxContentResult) Equals(that BoxContentResult) bool {
	return reflect.DeepEqual(b.Boxes, that.Boxes) && reflect.DeepEqual(b.Things, that.Things)
}
