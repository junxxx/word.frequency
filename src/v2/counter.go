package v2

import (
	"io"
	"sort"
	"strconv"
	"strings"
)

type Counter interface {
	io.Reader
	io.Writer
	Calculate()
}

var trimCharacter = []string{"\"", ".", ",", "“"}

type Word struct {
	word  string
	count int
}

// counter do count the word, other thing should do by caller
type sortedWord []Word

func (sw sortedWord) Len() int           { return len(sw) }
func (sw sortedWord) Less(i, j int) bool { return sw[i].count > sw[j].count }
func (sw sortedWord) Swap(i, j int)      { sw[i], sw[j] = sw[j], sw[i] }

type WordCounter struct {
	source []byte
	result map[string]int
	sortedWord
}

func (w *WordCounter) Read(p []byte) {
	w.source = make([]byte, len(p))
	copy(w.source, p)
}

func (w *WordCounter) Write() {

}

func trim(s *string, c []string) {
	for _, i := range c {
		*s = strings.Trim(*s, i)
	}
}

func (w *WordCounter) Calculate() *WordCounter {
	w.result = make(map[string]int)
	w.sortedWord = make([]Word, 0)
	for _, word := range strings.Fields(string(w.source)) {
		trim(&word, trimCharacter)
		w.result[word]++
	}
	w.Sort()

	return w
}

func (w *WordCounter) Sort() {
	w.sortedWord = make([]Word, 0)
	for word, cnt := range w.result {
		w.sortedWord = append(w.sortedWord, Word{word, cnt})
	}
	sort.Sort(w.sortedWord)
}

func (w WordCounter) String() string {
	var out string
	for _, word := range w.sortedWord {
		out += word.word + ":" + strconv.Itoa(word.count) + "\n"
	}
	return out
}
