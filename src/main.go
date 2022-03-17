package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	v2 "github.com/junxxx/word.frequency/v2"
)

var f string

func init() {
	flag.StringVar(&f, "f", "", "cal the world of the file")
}

func main() {
	flag.Parse()
	r, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	text, err := io.ReadAll(r)
	var wordCounter v2.WordCounter
	wordCounter.Read([]byte(text))
	wordCounter.Calculate()
	fmt.Println(wordCounter)
}
