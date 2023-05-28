package utils

import (
	"fmt"
	"io/ioutil"
	"time"

	glog "github.com/magicsong/color-glog"
	"gopkg.in/yaml.v3"
)

const INVALID_STATUS_CODE_ERROR = "Invalid statusCode, expected %v, got %v\n"
const INVALID_BODY_ERROR = "Invalid body, expected %v, got %v\n"
const VARS = "vars"

// Custom map
type myMap map[string]interface{}

type Expected struct {
	StatusCodes []uint8
	Body        map[string]interface{}
}

type Suite struct {
	Start time.Time
	End   time.Time
	Files []*File
}

type File struct {
	FileName string
	Tests    []*FullTest
}

type FullTest struct {
	Name string
	Apis []*ApiTest
	Vars *map[string]interface{}
}

type ApiTest struct {
	Name     string
	Url      string
	Method   string
	Headers  map[string]string
	Body     map[string]interface{}
	Expected Expected
}

func Greet(p string) {
	glog.Info("GOAT")
	glog.Info("---GOlang Api Tester---")
	glog.Info("Scanning for YAML documents in filesPath: ", p)
}

func ReadTests(fp string) *Suite {
	//Initialize Suite
	s := Suite{
		Start: time.Now(),
	}
	glog.Info("Starting time: ", s.Start)
	//Getting files
	files, err := ioutil.ReadDir("./files/")
	if err != nil {
		glog.Fatal("Unexpected error: ", err)
	}
	//Iterate over each file
	for i, f := range files {
		glog.Info("reading file: ", i)
		//Read file
		yf, err := ioutil.ReadFile(fp + f.Name())
		if err != nil {
			glog.Warning("Skipping file: %s because error: %s\n", f.Name(), err)
		}
		file := &File{FileName: f.Name()}
		s.Files = append(s.Files, file)
		data := myMap{}
		//Decode file
		err = yaml.Unmarshal(yf, data)
		//For each test in yaml file get content
		for e := range data {
			glog.Info("Creating fulltest for file ", i)
			ft := &FullTest{}
			ft.Name = e
			keys := getKeys(data.get(e))
			for _, k := range keys {
				getStrategy(k).Add(k, data.get(e), ft)
			}
			file.Tests = append(file.Tests, ft)
			glog.Info(ft)
		}
		glog.Info(file)
	}
	glog.Info(s)
	return &s
}

func getStrategy(k string) KeyStrategy {
	switch k {
	case VARS:
		glog.Info("Var strategy selected: ", k)
		return &VarStrategy{}
	default:
		glog.Info("Api strategy selected: ", k)
		return &ApiStrategy{}
	}
}

// Get inner map
func (m *myMap) get(k string) myMap {
	return (*m)[k].(myMap)
}

// Get keys in a map
func getKeys(m myMap) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i += 1
	}
	return keys
}

// Get String values in a map
func (m *myMap) getStrValue(key string, val string) string {
	// log.Printf("getStrValue: %s %s", val, (m.get(key))[val])
	return (m.get(key))[val].(string)
}

// Get Map values in a map
func (m *myMap) getMapStrValues(key string, val string) map[string]string {
	input := (m.get(key))[val].(myMap)
	result := make(map[string]string, len(input))
	for k, v := range input {
		result[k] = fmt.Sprint(v)
	}
	// log.Printf("getMapValues %s %s", val, result)
	return result
}

// func ReadTests(fp string) []ApiTest {
// 	var tests []ApiTest
// 	//Getting files
// 	files, err := ioutil.ReadDir("./files/")
// 	if err != nil {
// 		log.Fatal("Unexpected error: ", err)
// 	}
// 	for _, f := range files {
// 		//read file
// 		yf, err := ioutil.ReadFile(fp + f.Name())
// 		if err != nil {
// 			log.Printf("Skipping file: %s because error: %s\n", f.Name(), err)
// 		}
// 		//parse file data to map
// 		data := make(map[string]interface{})
// 		err = yaml.Unmarshal(yf, data)
// 		for e := range data {
// 			t := data[e].(map[string]interface{})
// 			at, err := MakeTest(t)
// 			if err != nil {
// 				log.Println("Error creating test")
// 			}
// 			tests = append(tests, at)
// 		}
// 	}
// 	return tests
// }

// func AddHeaders(r *http.Request, h map[string]interface{}) {
// 	for k, v := range h {
// 		r.Header.Add(k, v.(string))
// 	}
// }

// func MakeTest(t map[string]interface{}) (ApiTest, error) {
// 	log.Println(t)
// 	e, ok := t["expected"].(map[string]interface{})
// 	if !ok {
// 		log.Fatal("error getting expected")
// 		e = map[string]interface{}{}
// 	}
// 	b, ok := t["body"].(map[string]interface{})
// 	if !ok {
// 		log.Println("Request body not found.")
// 		b = map[string]interface{}{}
// 	}
// 	ex := getExpected(e)

// 	at := ApiTest{
// 		Expected: ex,
// 		Api:      t["api"].(string),
// 		Method:   t["method"].(string),
// 		Headers:  t["headers"].(map[string]interface{}),
// 		Body:     b,
// 	}
// 	return at, nil
// }

// func RunTest(at ApiTest, c *http.Client, ch chan string) {
// 	req, err := http.NewRequest(at.Method, at.Api, nil)
// 	AddHeaders(req, at.Headers)
// 	r, err := c.Do(req)
// 	if err != nil {
// 		log.Println("Error calling api: ", at.Api, err)
// 		ch <- fmt.Sprintf("test %s failed! %s", at.Api, at.Expected.Body)
// 	}
// 	passed, err := executeAsserts(r, at)
// 	if err != nil {
// 		ch <- fmt.Sprintf("test %s failed!", at.Api)
// 	}
// 	if passed {
// 		msg := fmt.Sprintf("test %s passed\n", at.Api)
// 		ch <- msg
// 	}
// }

// func getExpected(e map[string]interface{}) expected {
// 	sc := []uint8{}
// 	ec := e["statusCodes"].([]interface{})
// 	for _, c := range ec {
// 		sc = append(sc, uint8(c.(int)))
// 	}
// 	var b map[string]interface{}
// 	if e["body"] != nil {
// 		b = e["body"].(map[string]interface{})
// 	} else {
// 		log.Println("Expected body is not present, setting default empty body.")
// 		b = map[string]interface{}{}
// 	}
// 	ex := expected{
// 		Body:        b,
// 		StatusCodes: sc,
// 	}
// 	return ex
// }

// func executeAsserts(r *http.Response, at ApiTest) (bool, error) {
// 	api := at.Api
// 	result := "failed"
// 	if slices.Contains(at.Expected.StatusCodes, uint8(r.StatusCode)) {
// 		result = "passed"
// 	}
// 	log.Printf("statusCode assertion for %s %s, expected: %v, got: %v", api, result, at.Expected.StatusCodes, r.StatusCode)
// 	return true, nil
// }
