package storage

import "github.com/Wastoids/boxesandthings-api/model"

type Repository struct{}

func (r Repository) GetTopLevelBoxesForUser(userID string) ([]model.Box, error) {
	return nil, nil
}
