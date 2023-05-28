package main

import (
	"flag"
	u "germansanz93/goat/utils"
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

	// ch := make(chan string)

	//Set client for http calls
	// var client *http.Client = http.DefaultClient

	// tests := u.ReadTests(settings.filesPath)

	// log.Println(tests)
	glog.Info("Starting Suite Creation...")
	suite := u.ReadTests(settings.filesPath)
	glog.Info(suite.End)
	// for _, at := range tests {
	// 	go u.RunTest(at, client, ch)
	// }

	// for i := 0; i < len(tests); i++ {
	// 	fmt.Println(<-ch)
	// }

}
