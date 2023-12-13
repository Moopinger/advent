package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Print("BoatGambler5000\n")

	dat, err := os.ReadFile("./input.txt")
	//trainling line

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(dat), "\n")
	//trailing line
	lines = lines[:len(lines)-1]

	myComp, err := NewCompetition(lines)

	if err != nil {
		log.Fatalf("Bad values for competition provided: %s", err)
	}

	fmt.Printf("[+] Loaded a total of %d races. Here are the stats: \n", len(myComp.Races))
	for _, race := range myComp.Races {
		var winning_answers []int
		fmt.Printf("Race length time a total of: %d milliseconds\nWorld Record is %d millimetres. Calculating odds...\n", race.Time, race.RecordDist)
		for i := 1; i <= race.Time; i += 1 {
			if race.RaceRecord(i) == true {
				winning_answers = append(winning_answers, i)
			}
		}
		fmt.Printf("[+] Completed calculating. There are a total of %d winning combinattions, \n\n", len(winning_answers))

	}
	//fmt.Printf("There is an overall edge of: %d \n\n", multiplySlice(winning_amounts))
}
