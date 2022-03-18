package v3

import (
	"io"
	"strings"
)

type Counter struct {
	r        io.Reader
	trimChar []string
}

func NewCounter(r io.Reader) *Counter {
	return &Counter{r: r}
}

func (c *Counter) Trim(char []string) *Counter {
	c.trimChar = char
	return c
}

func (c *Counter) Count() (count map[string]int, err error) {
	count = make(map[string]int)
	buf, err := io.ReadAll(c.r)
	if err != nil {
		return
	}
	for _, word := range strings.Fields(string(buf)) {
		trim(&word, c.trimChar)
		if isWord(word) {
			count[word]++
		}
	}
	return
}

// could optimization
func isWord(s string) bool {
	for _, c := range s {
		if (c >= 'A' && c <= 'Z') ||
			(c >= 'a' && c <= 'z') || c == '\'' {

		} else {
			return false
		}
	}
	return true
}

func trim(s *string, c []string) {
	for _, i := range c {
		*s = strings.Trim(*s, i)
	}
}
