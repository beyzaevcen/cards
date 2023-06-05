package main

import "fmt"

func main() {
	cards := newDeck() //slices
	fmt.Println(cards.convertDeck2String())

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

}
