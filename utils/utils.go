package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

const INVALID_STATUS_CODE_ERROR = "Invalid statusCode, expected %v, got %v\n"
const INVALID_BODY_ERROR = "Invalid body, expected %v, got %v\n"

type expected struct {
	StatusCodes []uint8
	Body        map[string]interface{}
}

type ApiTest struct {
	Api      string
	Method   string
	Headers  map[string]interface{}
	Body     map[string]interface{}
	Expected expected
}

func Greet(p string) {
	log.Println("GOAT")
	log.Println("---GOlang Api Tester---")
	log.Printf("Scanning for YAML documents in filesPath: %s\n", p)
}

func ReadTests(fp string) []ApiTest {
	var tests []ApiTest
	//Getting files
	files, err := ioutil.ReadDir("./files/")
	if err != nil {
		log.Fatal("Unexpected error: ", err)
	}
	for _, f := range files {
		//read file
		yf, err := ioutil.ReadFile(fp + f.Name())
		if err != nil {
			log.Printf("Skipping file: %s because error: %s\n", f.Name(), err)
		}
		//parse file data to map
		data := make(map[string]interface{})
		err = yaml.Unmarshal(yf, data)
		for e := range data {
			t := data[e].(map[string]interface{})
			at, err := MakeTest(t)
			if err != nil {
				log.Println("Error creating test")
			}
			tests = append(tests, at)
		}
	}
	return tests
}

func AddHeaders(r *http.Request, h map[string]interface{}) {
	for k, v := range h {
		r.Header.Add(k, v.(string))
	}
}

func MakeTest(t map[string]interface{}) (ApiTest, error) {
	log.Println(t)
	e, ok := t["expected"].(map[string]interface{})
	if !ok {
		log.Fatal("error getting expected")
		e = map[string]interface{}{}
	}
	b, ok := t["body"].(map[string]interface{})
	if !ok {
		log.Println("Request body not found.")
		b = map[string]interface{}{}
	}
	ex := getExpected(e)

	at := ApiTest{
		Expected: ex,
		Api:      t["api"].(string),
		Method:   t["method"].(string),
		Headers:  t["headers"].(map[string]interface{}),
		Body:     b,
	}
	return at, nil
}

func RunTest(at ApiTest, c *http.Client, ch chan string) {
	req, err := http.NewRequest(at.Method, at.Api, nil)
	AddHeaders(req, at.Headers)
	r, err := c.Do(req)
	if err != nil {
		log.Println("Error calling api: ", at.Api, err)
		ch <- fmt.Sprintf("test %s failed! %s", at.Api, at.Expected.Body)
	}
	passed, err := executeAsserts(r, at)
	if err != nil {
		ch <- fmt.Sprintf("test %s failed!", at.Api)
	}
	if passed {
		msg := fmt.Sprintf("test %s passed\n", at.Api)
		ch <- msg
	}
}

func getExpected(e map[string]interface{}) expected {
	sc := []uint8{}
	ec := e["statusCodes"].([]interface{})
	for _, c := range ec {
		sc = append(sc, uint8(c.(int)))
	}
	var b map[string]interface{}
	if e["body"] != nil {
		b = e["body"].(map[string]interface{})
	} else {
		log.Println("Expected body is not present, setting default empty body.")
		b = map[string]interface{}{}
	}
	ex := expected{
		Body:        b,
		StatusCodes: sc,
	}
	return ex
}

func executeAsserts(r *http.Response, at ApiTest) (bool, error) {
	api := at.Api
	result := "failed"
	if slices.Contains(at.Expected.StatusCodes, uint8(r.StatusCode)) {
		result = "passed"
	}
	log.Printf("statusCode assertion for %s %s, expected: %v, got: %v", api, result, at.Expected.StatusCodes, r.StatusCode)
	return true, nil
}
