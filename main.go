package main

import (
	u "germansanz93/goat/utils"
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

	//Greet
	u.Greet(settings.filesPath)

	//Getting files
	files, err := ioutil.ReadDir("./files/")
	if err != nil {
		log.Fatal("Unexpected error: ", err)
	}

	//Reading each file
	for _, f := range files {
		//read file
		yf, err := ioutil.ReadFile(settings.filesPath + f.Name())
		if err != nil {
			log.Printf("Skipping file: %s because error: %s\n", f.Name(), err)
		}

		//parse file data to map
		data := make(map[string]interface{})
		err = yaml.Unmarshal(yf, data)
		for e := range data {
			t := data[e].(map[string]interface{})
			at, err := u.MakeTest(t)
			if err != nil {
				log.Println("Error creating test")
			}
			st := u.SelectStrategy(at.Method)
			r, err := st.Call(at)
			if err != nil {
				log.Println("Error calling api: ", at.Api, err)
			} else {
				// io.Copy(os.Stdout, r.Body)
				log.Printf("api: %s, status: %d", at.Api, r.StatusCode)
			}
			log.Println(at)
		}
	}
}
