package services

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"os"
)

const (
	ErrorFailedToMap = "failed to map bookmark"
	ErrorFailedToGetBookMark     = "failed to get bookmark"
	ErrorInvalidBookMark     = "invalid bookmark"
	ErrorCouldNotMapBookMark     = "could not map bookmark"
	ErrorCouldNotDeleteBookMark      = "could not delete"
	ErrorCouldNotDynamoUpdateBookMark   = "could not dynamo update bookmark"
	ErrorBookMarkAlreadyExists   = "BookMark already exists"
	ErrorBookMarkDoesNotExists   = "BookMark does not exist"
	TableName = "FavoulinksDB"
)

var dynaClient dynamodbiface.DynamoDBAPI

type BookMark struct {
	Title       string `json:"title" validate:"required"`
	URL 		string `json:"url" validate:"required"`
	Category	string `json:"category"`
}

func init()  {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return
	}
	dynaClient = dynamodb.New(awsSession)
}

func GetBookMark(url string) (*BookMark, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"url": {
				S: aws.String(url),
			},
		},
		TableName: aws.String(TableName),
	}

	data, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToGetBookMark)

	}

	bookmark := new(BookMark)

	if err = dynamodbattribute.UnmarshalMap(data.Item, bookmark);  err != nil {
		return nil, errors.New(ErrorFailedToMap)
	}

	return bookmark, nil
}

func GetAllBookMarks() (*[]BookMark, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}
	data, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToGetBookMark)
	}
	bookMark := new([]BookMark)
	err = dynamodbattribute.UnmarshalListOfMaps(data.Items, bookMark)
	return bookMark, nil
}

func CreateBookMark(req events.APIGatewayProxyRequest) (*BookMark, error) {
	var bookMark BookMark
	if err := json.Unmarshal([]byte(req.Body), &bookMark); err != nil {
		return nil, errors.New(ErrorInvalidBookMark)
	}

	initialBookMark, _ := GetBookMark(bookMark.URL)
	if initialBookMark != nil && len(initialBookMark.URL) != 0 {
		return nil, errors.New(ErrorBookMarkAlreadyExists)
	}

	if err := checkBookMark(bookMark); err != nil {
		return nil, err
	}

	newBookMark, err := dynamodbattribute.MarshalMap(bookMark)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMapBookMark)
	}

	putItem := &dynamodb.PutItemInput{
		Item:      newBookMark,
		TableName: aws.String(TableName),
	}

	_, err = dynaClient.PutItem(putItem)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoUpdateBookMark)
	}
	return &bookMark, nil
}

func UpdateBookMark(req events.APIGatewayProxyRequest) (*BookMark,error) {
	var bookMark BookMark
	if err := json.Unmarshal([]byte(req.Body), &bookMark); err != nil {
		return nil, errors.New(ErrorBookMarkDoesNotExists)
	}

	initialBookMark, _ := GetBookMark(bookMark.URL)
	if initialBookMark != nil && len(initialBookMark.URL) == 0 {
		return nil, errors.New(ErrorBookMarkDoesNotExists)
	}

	newBookMark, err := dynamodbattribute.MarshalMap(bookMark)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMapBookMark)
	}

	putItem := &dynamodb.PutItemInput{
		Item:      newBookMark,
		TableName: aws.String(TableName),
	}

	_, err = dynaClient.PutItem(putItem)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoUpdateBookMark)
	}
	return &bookMark, nil
}

func DeleteBookMark(req events.APIGatewayProxyRequest) error {
	url := req.QueryStringParameters["url"]
	bookMark := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"url": {
				S: aws.String(url),
			},
		},
		TableName: aws.String(TableName),
	}
	_, err := dynaClient.DeleteItem(bookMark)
	if err != nil {
		return errors.New(ErrorCouldNotDeleteBookMark)
	}

	return nil
}

func checkBookMark(b BookMark) error {
	err := ""
	if b.URL == "" {
		err += "URL empty. "
	}
	if b.Title == "" {
		err += "Title empty."
	}
	if err == "" {
		return nil
	} else {
		return errors.New(err)
	}
}