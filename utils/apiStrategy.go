package utils

import "log"

const URL string = "url"
const METHOD string = "method"

type ApiStrategy struct{}

func (a *ApiStrategy) Add(key string, d myMap, ft *FullTest) {
	at := ApiTest{
		Name:   key,
		Url:    getStrValue(key, URL, d),
		Method: getStrValue(key, METHOD, d),
		//TODO crear metodo getHeaders y los metodos que faltan para completar la creacion completa del apitest
	}

	ft.Apis = append(ft.Apis, &at)

}

func getStrValue(key string, val string, m myMap) string {
	log.Printf("getStrValue: %s %s", val, m.get(key)[val])
	return m.get(key)[val].(string)
}
