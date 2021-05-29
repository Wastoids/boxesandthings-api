package main

import (
	"encoding/json"
	"log"

	"github.com/Wastoids/boxesandthings-api/controller"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(event events.APIGatewayProxyRequest) (interface{}, error) {
	f, err := controller.NewController().GetFunction(event)
	if err != nil {
		return getError(400, err), nil
	}
	response, err := f.Run()
	if err != nil {
		return getError(400, err), nil
	}
	json, err := json.Marshal(response)
	if err != nil {
		return getError(500, err), nil
	}
	res := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:            string(json),
		IsBase64Encoded: false,
	}
	return res, nil
}

func getError(statusCode int, err error) events.APIGatewayProxyResponse {
	log.Printf("there was an error: %v", err)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body:            err.Error(),
		IsBase64Encoded: false,
	}
}
