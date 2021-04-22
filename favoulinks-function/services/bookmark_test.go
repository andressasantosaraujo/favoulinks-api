package services

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	bbc = BookMark{"BBC","bbc.co.uk","News" }
	globo = BookMark{"Globo", "globo.com","News" }
	twitter = BookMark{ "Twitter", "twitter.com", "Social" }
)

func mockRequest(body BookMark) events.APIGatewayProxyRequest {
	strBody, _ := json.Marshal(body)
	return events.APIGatewayProxyRequest {
			Body: string(strBody),
			QueryStringParameters: map[string]string{"url": body.URL},
		}
}

func createAll(t *testing.T) {
	listBookMarks := []BookMark{bbc, globo, twitter}
	for _, v := range listBookMarks {
		bookmark, err := CreateBookMark(mockRequest(v))
		assert.Nil(t, err)
		assert.Equal(t, v.Title, bookmark.Title)
		assert.Equal(t, v.URL, bookmark.URL)
		assert.Equal(t, v.Category, bookmark.Category)
	}
}

func resetBookMarkTest(t *testing.T) {
	listBookMarks := []BookMark{bbc, globo, twitter}
	for _, v := range listBookMarks {
		err := DeleteBookMark(mockRequest(v))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestBookMark_ReturnBookMark_WhenCreateBookMarkAndGetBookMark(t *testing.T) {
	resetBookMarkTest(t)
	createAll(t)
	listBookMarks := []BookMark{bbc, globo, twitter}
	for _, v := range listBookMarks {
		bookmark, err := GetBookMark(v.URL)
		assert.Nil(t, err)
		assert.Equal(t, v.Title, bookmark.Title)
		assert.Equal(t, v.URL, bookmark.URL)
		assert.Equal(t, v.Category, bookmark.Category)
	}
}

func TestBookMark_ReturnUpdatedBookMark_WhenUpdateAfterCreation(t *testing.T) {
	resetBookMarkTest(t)
	createAll(t)
	bbc.Category = "Teste"
	globo.Category = "Teste"
	twitter.Category = "Teste"
	listBookMarks := []BookMark{bbc, globo, twitter}
	for _, v := range listBookMarks {
		bookmark, err := UpdateBookMark(mockRequest(v))
		assert.Nil(t, err)
		assert.Equal(t, v.Title, bookmark.Title)
		assert.Equal(t, v.URL, bookmark.URL)
		assert.Equal(t, v.Category, bookmark.Category)
	}
}

func TestBookMark_ReturnNotFound_WhenUpdatedIdDoesntExist(t *testing.T) {
	resetBookMarkTest(t)
	listBookMarks := []BookMark{bbc, globo, twitter}
	for _, v := range listBookMarks {
		bookmark, err := UpdateBookMark(mockRequest(v))
		assert.Nil(t, bookmark)
		assert.Equal(t, errors.New(ErrorBookMarkDoesNotExists), err)
	}
}

func TestBookMark_ReturnNil_WhenDelete(t *testing.T) {
	resetBookMarkTest(t)
	createAll(t)
	listBookMarks := []BookMark{bbc, globo, twitter}
	for _, v := range listBookMarks {
		err := DeleteBookMark(mockRequest(v))
		assert.Nil(t, err)
	}
}

func TestBookMark_ReturnAllBookMarks(t *testing.T) {
	resetBookMarkTest(t)
	createAll(t)
	bookMarks, err := GetAllBookMarks()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(*bookMarks), 3 )
}
