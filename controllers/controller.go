package controllers

import "errors"

type Function interface {
	Run() (interface{}, error)
}

type Controller struct{}

var errInvalidResource = errors.New("invalid resource requested")

func (c Controller) GetFunction(resource string) (Function, error) {
	switch resource {
	default:
		return nil, errInvalidResource
	}
}
