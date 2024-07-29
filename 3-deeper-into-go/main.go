package main

func main() {
	// Saving data to the hard drive
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// Reading from the hard drive
	cards := newDeckFromFile("my_cards")
	cards.print()
}
