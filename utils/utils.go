package utils

import (
	"log"
)

const GET = "GET"
const POST = "POST"
const PUT = "PUT"
const DELETE = "DELETE"

type ApiTest struct {
	Api      string
	Method   string
	Expected map[string]interface{}
}

func SelectStrategy(at ApiTest) TestCall {
	switch at.Method {
	case GET:
		return GetStrategy{}
	case POST:
		log.Println("POST")
		return nil
	case PUT:
		log.Println("PUT")
		return nil
	case DELETE:
		log.Println("DELETE")
		return nil
	default:
		log.Println("MISSING METHOD")
		return nil
	}
}

func MakeTest(test map[string]interface{}) (ApiTest, error) {
	expected, ok := test["expected"].(map[string]interface{})
	if !ok {
		log.Fatal("Erorr reading test!")
	}
	if len(expected) == 0 {
		log.Println("No expected values, setting default assertion codes 200, 204")
		expected["code"] = []string{"200", "204"}
	}
	log.Println(expected)
	at := ApiTest{
		Api:      test["api"].(string),
		Expected: expected,
		Method:   test["method"].(string),
	}
	return at, nil
}
