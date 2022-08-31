package dev04

import (
	"sort"
	"strings"
)

func Run(input *[]string) map[string][]string {
	result := make(map[string][]string)
	nominal := make(map[string]string)
	for _, str := range *input {
		if len(str) < 2 {
			continue
		}
		sorting(str, &nominal, &result)
	}
	for k, v := range result {
		if len(v) == 0 {
			delete(result, k)
		}
	}
	return result
}
func sorting(str string, n *map[string]string, r *map[string][]string) {
	lower := strings.ToLower(str)
	byteLower := []rune(lower)
	sortByte := []rune(lower)
	sort.Slice(sortByte, func(i, j int) bool {
		return sortByte[i] <= sortByte[j]
	})
	if v, ok := (*n)[string(sortByte)]; ok {
		(*r)[v] = append((*r)[v], string(byteLower))
	} else {
		(*n)[string(sortByte)] = string(byteLower)
		(*r)[string(byteLower)] = []string{}
	}
}
