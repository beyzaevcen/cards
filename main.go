package main

func main() {
	cards := newDeck() //slices

	hand, remainingCards := deal(cards, 2)
	hand.print()
	remainingCards.print()

}
