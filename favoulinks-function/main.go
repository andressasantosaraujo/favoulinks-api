package main

import (
	"favoulinks-function/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)



func main() {
	lambda.Start(router)
}

func router(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetBookMark(req)
	case "POST":
		return handlers.CreateBookMark(req)
	case "PUT":
		return handlers.UpdateBookMark(req)
	case "DELETE":
		return handlers.DeleteBookMark(req)
	default:
		return handlers.MethodNotAllowed()
	}
}