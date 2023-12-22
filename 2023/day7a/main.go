package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Read file into a byte slice
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//remove trailing line
	data = data[:len(data)-1]
	//boom
	myCamelComp := NewCamelCompetition(data)
	total := 0

	for index, hand := range myCamelComp.Hands {
		bid_points := (index + 1) * hand.Bid
		total += bid_points
		fmt.Printf("Cards in hand: %s with a type of %d\n", hand.printHand(), hand.Type)

		// for _, card := range hand.Cards {
		// 	fmt.Printf("The card %s has a value of %d\n", card.Value, card.Points)
		// }
	}
	fmt.Printf("Total points are %d\n", total)
}
