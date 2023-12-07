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

		games = append(games, game)

		power := game.MinValue.Red * game.MinValue.Green * game.MinValue.Blue

		fmt.Printf("[+]Calculated the power for %v it is %d\n", game.Id, power)

		total += power
	}

	fmt.Printf("[+] Total is: %d\n", total)

}
