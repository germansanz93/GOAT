package utils

const URL string = "url"
const METHOD string = "method"
const HEADERS string = "headers"
const EXPECTED_BODY string = "body"

type ApiStrategy struct{}

func (a *ApiStrategy) Add(key string, d myMap, ft *FullTest) {
	at := ApiTest{
		Name:     key,
		Url:      d.getStrValue(key, URL),
		Method:   d.getStrValue(key, METHOD),
		Headers:  d.getMapStrValues(key, HEADERS),
		Expected: getExpected(d.get(key)),
		//TODO Falta trabajar el caso de los puts o los posts, donde necesitamos enviar un body
	}

	ft.Apis = append(ft.Apis, &at)

}

func getExpected(m myMap) Expected {
	return Expected{} // TODO completar esta funcion
}
