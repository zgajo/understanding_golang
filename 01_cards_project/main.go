package main

func main() {
	// Two ways to define variables,
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // only to define new variable

	// array vs slice
	// Array - fixed length list of things
	// Slice - array that cang grow and shrink
	// slice
	cards := deck{"Ace", newCard()}

	cards.print()
}

func newCard() string {
	return "Test"
}
