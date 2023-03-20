package main

import "fmt"

// Create a new type of 'deck'
// Just a slice of string

type deck []string

func (d deck) newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Bich", "Xi", "Ro", "Co"}
	cardValues := []string{"Hai", "Ba", "Bon", "Nam", "Sau", "Bay", "Tam", "Chin", "Muoi", "Boi", "Dam", "Gia", "Xi"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" "+suite)

		}

	}
	// for _, card := range cards {
	// 	fmt.Println(card)
	// }
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
