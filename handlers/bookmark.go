package handlers

import (
	"encoding/json"
	"favoulinks-function/services"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"net/http"
)

var ErrorMethodNotAllowed = "method Not allowed"

type Error struct {
	Msg *string `json:"error,omitempty"`
}

func GetBookMark(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	url := req.QueryStringParameters["url"]
	if len(url) > 0 {
		data, err := services.GetBookMark(url)
		if err != nil {
			return buildResponse(http.StatusBadRequest, Error{aws.String(err.Error())})
		}

		return buildResponse(http.StatusOK, data)
	}

	result, err := services.GetAllBookMarks()
	if err != nil {
		return buildResponse(http.StatusBadRequest, Error{aws.String(err.Error())})
	}
	return buildResponse(http.StatusOK, result)
}

func CreateBookMark(req events.APIGatewayProxyRequest) ( *events.APIGatewayProxyResponse, error) {
	data, err := services.CreateBookMark(req)
	if err != nil {
		return buildResponse(http.StatusBadRequest, Error{aws.String(err.Error())})
	}
	return buildResponse(http.StatusCreated, data)
}

func UpdateBookMark(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	data, err := services.UpdateBookMark(req)
	if err != nil {
		return buildResponse(http.StatusBadRequest, Error{aws.String(err.Error())})
	}
	return buildResponse(http.StatusOK, data)
}

func DeleteBookMark(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if err := services.DeleteBookMark(req); err != nil {
		return buildResponse(http.StatusBadRequest, Error{aws.String(err.Error())})
	}
	return buildResponse(http.StatusOK, nil)
}

func MethodNotAllowed() (*events.APIGatewayProxyResponse, error) {
	return buildResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}

func buildResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	strBody, _ := json.Marshal(body)
	return &events.APIGatewayProxyResponse{
		Headers: map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: status,
		Body: string(strBody),
	}, nil
}