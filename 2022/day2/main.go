package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("Game checker 5000\n")

	var games []Game
	var entries []string
	var gameid int64 = 1
	var total int64 = 0

	dat, err := os.ReadFile("./list.txt")

	if err != nil {
		log.Fatal(err)
	}

	entries = strings.Split(string(dat), "\n")

	for _, entry := range entries {

		if entry == "" {
			log.Printf("Skipping blank entry")
		} else {
			thisgame := NewGame(gameid, entry)

			games = append(games, thisgame)
			gameid += 1
		}

	}

	for _, game := range games {

		total += game.Score
	}

	fmt.Printf("[+]Total is: %v\n", total)

	check(err)

}
