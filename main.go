package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	filesPath string `yaml:"filesPath"`
}

type ApiTest struct {
	api      string
	method   string
	expected map[string]interface{}
}

var settings = Settings{ //TODO hacer esto configurable con un archivo settings.yaml
	filesPath: "./files/",
}

func makeTest(test map[string]interface{}) (ApiTest, error) {
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
		api:      test["api"].(string),
		expected: expected,
		method:   test["method"].(string),
	}
	return at, nil
}

func main() {

	//Greet
	log.Println("GOAT")
	log.Println("---GOlang Api Tester---")
	log.Printf("Scanning for YAML documents in filesPath: %s\n", settings.filesPath)

	//Getting files
	files, err := ioutil.ReadDir("./files/")
	if err != nil {
		log.Fatal("Unexpected error: ", err)
	}

	//Reading each file
	for _, file := range files {
		//read file
		yamlFile, err := ioutil.ReadFile(settings.filesPath + file.Name())
		if err != nil {
			log.Printf("Skipping file: %s because error: %s\n", file.Name(), err)
		}

		//parse file data to map
		data := make(map[string]interface{})
		err = yaml.Unmarshal(yamlFile, data)
		for endpoint := range data {
			test := data[endpoint].(map[string]interface{})
			at, err := makeTest(test)
			if err != nil {
				log.Println("Error creating test")
			}
			response, err := http.Get(at.api)
			if err != nil {
				log.Println("Error calling api: ", at.api, err)
			} else {
				io.Copy(os.Stdout, response.Body)
			}
			log.Println(at)
		}
	}
}
