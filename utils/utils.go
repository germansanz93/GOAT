package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const INVALID_STATUS_CODE_ERROR = "Invalid statusCode, expected %v, got %v\n"

type ApiTest struct {
	Api      string
	Method   string
	Headers  map[string]interface{}
	Body     map[string]interface{}
	Expected map[string]interface{}
}

func Greet(p string) {
	log.Println("GOAT")
	log.Println("---GOlang Api Tester---")
	log.Printf("Scanning for YAML documents in filesPath: %s\n", p)
}

func AddHeaders(r *http.Request, h map[string]interface{}) {
	for k, v := range h {
		r.Header.Add(k, v.(string))
	}
}

func MakeTest(t map[string]interface{}) (ApiTest, error) {
	e, ok := t["expected"].(map[string]interface{})
	if !ok {
		log.Fatal("Erorr reading test!")
	}
	if len(e) == 0 {
		log.Println("No expected values, setting default assertion codes 200, 204")
		e["code"] = []string{"200", "204"}
	}
	h, ok := t["headers"].(map[string]interface{})
	if !ok {
		log.Fatal("Error reding headers")
	}

	at := ApiTest{
		Api:      t["api"].(string),
		Expected: e,
		Headers:  h,
		Method:   t["method"].(string),
	}
	return at, nil
}

func RunTest(at ApiTest, c *http.Client) (bool, error) {
	req, err := http.NewRequest(at.Method, at.Api, nil)
	AddHeaders(req, at.Headers)
	r, err := c.Do(req)
	if err != nil {
		log.Println("Error calling api: ", at.Api, err)
		return false, err
	}
	passed, err := executeAsserts(r, at)
	if err != nil {
		return false, err
	}
	if passed {
		return true, nil
	}
	return false, nil
}

func executeAsserts(r *http.Response, at ApiTest) (bool, error) {
	if at.Expected["statusCode"] != r.StatusCode {
		return false, errors.New(fmt.Sprintf(INVALID_STATUS_CODE_ERROR, r.StatusCode, at.Expected["statusCode"]))
	}
	return true, nil
}
