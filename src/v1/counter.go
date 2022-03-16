package v1

import "strings"

func trim(s *string, c []string) {
	for _, i := range c {
		*s = strings.Trim(*s, i)
	}
}

func Counter(resource string, trimCharacter []string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range strings.Fields(resource) {
		trim(&word, trimCharacter)
		wordCount[word]++
	}
	return wordCount
}
