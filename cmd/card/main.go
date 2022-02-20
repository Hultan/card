package main

import (
	"fmt"

	"card/internal/decks"
)

func main() {
	d := decks.NewDecks(2)
	d.Main().Shuffle()
	for i := 0; i < decks.NumberOfCardsInDeck; i++ {
		fmt.Println(d.Main().Deal(2))
	}
	t := d.Main().Deal(1)
	if t != nil {
		fmt.Println("MAJOR BUG!")
	}
}
