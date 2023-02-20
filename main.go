package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	filesPath string `yaml:"filesPath"`
}

var settings = Settings{ //TODO hacer esto configurable con un archivo settings.yaml
	filesPath: "./files/",
}

func main() {

	log.Println("GOAT")
	log.Println("---GOlang Api Tester---")
	log.Printf("Scanning for YAML documents in filesPath: %s\n", settings.filesPath)

	files, err := ioutil.ReadDir("./files/")
	if err != nil {
		log.Fatal("Unexpected error: ", err)
	}

	for _, file := range files {
		yamlFile, err := ioutil.ReadFile(settings.filesPath + file.Name())
		if err != nil {
			log.Printf("Skipping file: %s because error: %s\n", file.Name(), err)
		}
		data := make(map[interface{}]interface{})
		err = yaml.Unmarshal(yamlFile, data)
		for endpoint := range data {
			log.Println(endpoint)
		}
	}
}
