package v3

import (
	"io"
	"strconv"
)

type Counter struct {
	r     io.Reader
	count map[string]int
}

func NewCounter(r io.Reader) *Counter {
	return &Counter{r: r, count: make(map[string]int)}
}

func isCharacter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func (c *Counter) Count() error {
	buf, err := io.ReadAll(c.r)
	if err != nil {
		return err
	}
	word := make([]byte, 1)
	for _, char := range buf {
		if isCharacter(char) {
			word = append(word, char)
		} else {
			if len(word) > 0 {
				c.count[string(word)]++
				word = nil
			}
		}
	}
	return nil
}

func (c Counter) String() string {
	var out string
	for word, num := range c.count {
		out += word + "    " + strconv.Itoa(num) + "\n"
	}
	return out
}
