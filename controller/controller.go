package controller

import "errors"

var errInvalidResource = errors.New("invalid resource requested")

type Function interface {
	Run() (interface{}, error)
}

type Controller struct{}

func (c Controller) GetFunction(resource string) (Function, error) {
	switch resource {
	default:
		return nil, errInvalidResource
	}
}

func NewController() Controller {
	return Controller{}
}
