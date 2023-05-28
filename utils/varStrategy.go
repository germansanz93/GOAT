package utils

import (
	"fmt"
)

type VarStrategy struct{}

func (v *VarStrategy) Add(key string, d myMap, ft *FullTest) {
	m := map[string]interface{}{}
	vars := getVars(key, d)
	for _, v := range vars {
		m[fmt.Sprint(v)] = ""
	}
	ft.Vars = &m
}

func getVars(k string, m myMap) []string {
	va := m[k].([]interface{})
	vars := make([]string, len(va))
	for i, v := range va {
		vars[i] = fmt.Sprint(v)
	}
	return vars
}
