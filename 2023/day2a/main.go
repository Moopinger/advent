package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var games []Game
	var total int64 = 0

	dat, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	game_inputs := strings.Split(string(dat), "\n")

	for _, game_input := range game_inputs {

		if game_input == "" {
			break
		}
		game := NewGame(game_input)
		fmt.Println(game)

		games = append(games, game)

		if game.MinValue.Red <= 12 && game.MinValue.Green <= 13 && game.MinValue.Blue <= 14 {
			//fmt.Printf("ID %d Checks out.. Adding id to total... \n", game.Id)
			total += game.Id
		}
	}

	fmt.Printf("[+] Total is: %d\n", total)

}
