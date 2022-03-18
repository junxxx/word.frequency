package main

import (
	"flag"
	"fmt"
	"os"

	v3 "github.com/junxxx/word.frequency/v3"
)

var f string
var trimCharacter = []string{"\"", ".", ",", "“"}

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
	defer r.Close()
	wordCount, err := v3.NewCounter(r).Trim(trimCharacter).Count()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wordCount)
}
