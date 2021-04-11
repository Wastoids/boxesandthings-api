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
	f, err := controller.NewController().GetFunction(event.PathParameters["resource"])
	if err != nil {
		panic("invalid resource path")
	}
	return f.Run()
}
