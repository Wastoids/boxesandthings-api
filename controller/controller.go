package controller

import (
	"errors"

	"github.com/Wastoids/boxesandthings-api/service"
	"github.com/Wastoids/boxesandthings-api/storage"
	"github.com/aws/aws-lambda-go/events"
)

var errInvalidResource = errors.New("invalid resource requested")

type Function interface {
	Run() (interface{}, error)
}

type Controller struct{}

func (c Controller) GetFunction(e events.APIGatewayProxyRequest) (Function, error) {

	switch e.PathParameters["resource"] {
	case "topBoxes":
		return service.NewGetTopBoxesService(storage.NewRepository(), e.QueryStringParameters["username"]), nil
	case "saveBox":
		return service.NewSaveBox(storage.NewRepository(), service.GetBoxFromRequest(e.Body)), nil
	default:
		return nil, errInvalidResource
	}
}

func NewController() Controller {
	return Controller{}
}
