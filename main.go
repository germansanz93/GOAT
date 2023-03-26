package main

import (
	u "germansanz93/goat/utils"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	filesPath string `yaml:"filesPath"`
}

var settings = Settings{ //TODO hacer esto configurable con un archivo settings.yaml
	filesPath: "./files/",
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
			at, err := u.MakeTest(test)
			if err != nil {
				log.Println("Error creating test")
			}
			st := u.SelectStrategy(at)
			response, err := st.Call(at)
			if err != nil {
				log.Println("Error calling api: ", at.Api, err)
			} else {
				io.Copy(os.Stdout, response.Body)
			}
			log.Println(at)
		}
	}
}
