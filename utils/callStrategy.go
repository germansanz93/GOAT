package utils

import "net/http"

type TestCall interface {
	Call(at ApiTest) (*http.Response, error)
}
