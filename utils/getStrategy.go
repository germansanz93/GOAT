package utils

import (
	"errors"
	"log"
	"net/http"
)

const API_CALL_ERROR = "Error calling api: "

type GetStrategy struct{}

func (g GetStrategy) Call(at ApiTest) (*http.Response, error) {
	response, err := http.Get(at.Api)
	if err != nil {
		log.Println(API_CALL_ERROR, at.Api)
		return nil, errors.New(API_CALL_ERROR + at.Api)
	}
	return response, nil
}
