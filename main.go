package main

import (
	"flag"
	"fmt"
	u "germansanz93/goat/utils"
	"net/http"
	"runtime/pprof"

	glog "github.com/magicsong/color-glog"
)

type Settings struct {
	filesPath string `yaml:"filesPath"`
}

var settings = Settings{ //TODO hacer esto configurable con un archivo settings.yaml
	filesPath: "./files/",
}
var threadProfie = pprof.Lookup("threadcreate")

func main() {
	flag.Parse()
	//Greet
	u.Greet(settings.filesPath)

	// fmt.Println(runtime.NumCPU())
	// fmt.Println(threadProfie.Count())

	ch := make(chan string)

	//Set client for http calls
	var client *http.Client = http.DefaultClient

	tests := 0
	glog.Info("Starting Suite Creation...")
	suite := u.ReadTests(settings.filesPath)
	glog.Info(suite.End)
	for _, file := range suite.Files {
		glog.Info("Starting tests for file: ", file.FileName)
		glog.Warning("o(n)")
		for _, test := range file.Tests {
			glog.Warning("o(n²)")
			glog.Info("Running test ", test.Name)
			for _, api := range test.Apis {
				glog.Warning("o(n³)")
				glog.Info("Calling api ", api)
				go u.RunTest(*api, client, ch) //TODO No se puede dejar este go aca.. Hay que separar esto a una func aparte para ir a goroutine con el test en particular y no con solo un api
				tests++
			}
		}
	}
	// for _, at := range tests {
	// 	go u.RunTest(at, client, ch)
	// }

	for i := 0; i < tests; i++ {
		fmt.Println(<-ch)
	}

}
