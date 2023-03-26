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
	Headers  map[string]interface{}
	Body     map[string]interface{}
	Expected map[string]interface{}
}

func Greet(p string) {
	log.Println("GOAT")
	log.Println("---GOlang Api Tester---")
	log.Printf("Scanning for YAML documents in filesPath: %s\n", p)
}

func SelectStrategy(m string) TestCall {
	switch m {
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
