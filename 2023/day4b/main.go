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

	for index, scrarch_card := range scratchcards {

		for i2 := 1; i2 <= scrarch_card.Copies; i2++ {

			if scrarch_card.Points > 0 {

				for i := 1; i <= scrarch_card.Points; i++ {

					scratchcards[index+i].CreateCopy()

				}

			}

		}
		fmt.Println(scrarch_card)
	}

	for _, scratchcard := range scratchcards {

		total += scratchcard.Copies

	}

	fmt.Printf("Total id %d \n", total)

}
