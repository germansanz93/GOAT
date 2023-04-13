package main

import (
	"fmt"
	u "germansanz93/goat/utils"
	"net/http"
	"runtime/pprof"
)

type Settings struct {
	filesPath string `yaml:"filesPath"`
}

var settings = Settings{ //TODO hacer esto configurable con un archivo settings.yaml
	filesPath: "./files/",
}
var threadProfie = pprof.Lookup("threadcreate")

func main() {

	//Greet
	u.Greet(settings.filesPath)

	// fmt.Println(runtime.NumCPU())
	// fmt.Println(threadProfie.Count())

	ch := make(chan string)

	//Set client for http calls
	var client *http.Client = http.DefaultClient

	tests := u.ReadTests(settings.filesPath)

	for _, at := range tests {
		go u.RunTest(at, client, ch)
	}

	for i := 0; i < len(tests); i++ {
		fmt.Println(<-ch)
	}

}
