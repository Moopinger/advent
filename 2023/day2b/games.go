package main

import (
	"strconv"
	"strings"
)

type Game struct {
	Id       int64
	Rounds   []Round
	MinValue Round
}

type Round struct {
	Red   int64
	Green int64
	Blue  int64
}

func getEmptyRound() Round {
	var r int64 = 0
	var g int64 = 0
	var b int64 = 0

	ro := Round{
		Red:   r,
		Green: g,
		Blue:  b,
	}

	return ro
}

func newRound(round_info string) Round {
	var r int64 = 0
	var g int64 = 0
	var b int64 = 0

	rounds_str := strings.Split(round_info, ", ")

	for _, round_str := range rounds_str {
		num_string := getInt(round_str)
		num_int, _ := strconv.Atoi(num_string)
		num_length := len(num_string) + 1
		color_code := string(round_str[num_length])

		if color_code == "r" {
			r += int64(num_int)
		} else if color_code == "g" {
			g += int64(num_int)
		} else if color_code == "b" {
			b += int64(num_int)
		}

	}

	ro := Round{
		Red:   r,
		Green: g,
		Blue:  b,
	}

	return ro
}
//I really need to come up with a better function name here
func getInt(input_string string) string {

	var id_string string = ""

	for _, input_char := range input_string {
		if _, err := strconv.Atoi(string(input_char)); err == nil {
			id_string += string(input_char)
		} else {
			break
		}
	}

	return id_string

}

func NewGame(game_results string) Game {
	//		 0 1 2 3 4 5
	//strip "G a m e   1"
	var rounds []Round

	game_results = game_results[5:]
	//fmt.Printf("DEBUG: %s \n", game_results)

	//store the number of charachters in the lemgth
	//var id_length int64 = 0

	//Store the game id before we convert AtoI
	var id_string string = ""

	id_string = getInt(game_results)
	//append 2 to get rid of ": "
	id_length := len(id_string) + 2
	id_int, _ := strconv.Atoi(id_string)

	minValue := getEmptyRound()

	//chop off the id: and
	game_results = game_results[id_length:]
	rounds_result := strings.Split(game_results, "; ")

	for _, round_string := range rounds_result {
		//Calculate min values
		myRound := newRound(round_string)

		if myRound.Red > minValue.Red {
			minValue.Red = myRound.Red
		}

		if myRound.Green > minValue.Green {
			minValue.Green = myRound.Green
		}

		if myRound.Blue > minValue.Blue {
			minValue.Blue = myRound.Blue
		}

		rounds = append(rounds, myRound)
	}

	g := Game{
		Id:       int64(id_int),
		Rounds:   rounds,
		MinValue: minValue,
	}

	return g

}
