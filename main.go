package main

func main() {
	cards := newDeck()

	// cards.print()

	hand, remainigCards := deal(cards, 5)

	hand.print()
	remainigCards.print()
}
