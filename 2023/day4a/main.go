package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	total := 0

	var scratchcards []ScratchCard

	dat, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scratchcards_data := strings.Split(string(dat), "\n")

	scratchcards = bulkAddScratchCards(scratchcards_data)

	for _, scrarch_card := range scratchcards {

		fmt.Println(scrarch_card)
		if scrarch_card.Points > 0 {
			total += scrarch_card.Points
		}
	}

	fmt.Printf("Total id %d \n", total)

}
