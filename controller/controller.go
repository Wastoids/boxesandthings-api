package controller

import (
	"errors"
	"net/http"

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
	repo := storage.NewRepository()
	email := e.RequestContext.Authorizer["claims"].(map[string]interface{})["email"].(string)
	switch router(e) {
	case "getTopBoxes":
		return service.NewGetTopBoxesService(repo, email), nil
	case "saveBox":
		return service.NewSaveBox(repo, service.GetBoxFromRequest(e.Body)), nil
	case "saveThing":
		return service.NewSaveThing(
				repo,
				service.GetThingFromRequest(e.Body),
				e.PathParameters["boxID"]),
			nil
	case "boxContent":
		return service.NewBoxContent(repo, e.PathParameters["boxID"]), nil
	case "createTopBox":
		return service.NewSaveTopBox(repo, email, service.GetBoxFromRequest(e.Body)), nil
	default:
		return nil, errInvalidResource
	}
}

func router(e events.APIGatewayProxyRequest) string {
	var result string
	resource := e.PathParameters["resource"]

	if resource == "top" {
		if e.HTTPMethod == http.MethodGet {
			return "getTopBoxes"
		}
		if e.HTTPMethod == http.MethodPost {
			return "createTopBox"
		}
	}

	if resource == "box" {
		if e.HTTPMethod == http.MethodPost {
			return "saveBox"
		}
	}

	if resource == "thing" {
		if e.HTTPMethod == http.MethodPost {
			return "saveThing"
		}
	}
	return result
}

func NewController() Controller {
	return Controller{}
}
