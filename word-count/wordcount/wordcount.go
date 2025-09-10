package wordcount

import "strings"

// Count counts words in the given string s and returns a map of the counts of each word in the string
func Count(s string) map[string]int {
	split := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '\t'
	})
	m := make(map[string]int)
	for _, word := range split {
		m[word]++
	}
	return m
}
