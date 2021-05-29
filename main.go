package main

import (
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
	return response, nil
}

func getError(statusCode int, err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body:            err.Error(),
		IsBase64Encoded: false,
	}
}
