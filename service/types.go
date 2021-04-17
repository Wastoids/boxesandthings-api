package service

import "github.com/Wastoids/boxesandthings-api/model"

type Storage interface {
	GetTopLevelBoxesForUser(userID string) ([]model.Box, error)
}
