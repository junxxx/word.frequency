package main

import (
	"fmt"

	v1 "github.com/junxxx/word.frequency/v1"
)

var trimCharacter = []string{"\"", ".", ",", "“"}

func main() {
	text := `Three, "European prime ministers visited Kyiv on Tuesday. The prime ministers of neighboring Poland, Czechia and Slovenia came to Ukraine by train for the first such visit by foreign leaders since the Russian invasion.

	"It is our duty to be where history is forged. Because it's not about us, but about the future of our children who deserve to live in a world free from tyranny," said Polish Prime Minister Mateusz Morawiecki. Czech Prime Minister Petr Fiala and Janez Jansa of Slovenia joined him on the trip.
	
	Fiala said the visit was to show complete “support of the entire European Union for the sovereignty and independence of Ukraine."
	`
	for word, cnt := range v1.Counter(text, trimCharacter) {
		fmt.Println(word, cnt)
	}
}
