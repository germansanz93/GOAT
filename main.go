package main

import (
	u "germansanz93/goat/utils"
	"log"
	"net/http"
)

type Settings struct {
	filesPath string `yaml:"filesPath"`
}

var settings = Settings{ //TODO hacer esto configurable con un archivo settings.yaml
	filesPath: "./files/",
}

func main() {

	//Greet
	u.Greet(settings.filesPath)

	//Set client for http calls
	var client *http.Client = http.DefaultClient

	tests := u.ReadTests(settings.filesPath)

	for _, at := range tests {
		passed, err := u.RunTest(at, client)
		if err != nil {
			log.Println(err)
		}
		if passed {
			log.Println("passed!")
		}
	}
}
