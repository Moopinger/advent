package main

import (
	"math"
	"strconv"
	"strings"
)

type ScratchCard struct {
	Id             int
	WinningNumbers []int
	Guesses        []int
	Points         int
}

func FilterEmptyConvert(input []string) []int {

	var filtered []int
	for _, value := range input {
		if value != "" {

			value_int, _ := strconv.Atoi(value)
			filtered = append(filtered, value_int)
		}
	}
	return filtered
}

func getInt(input_string string) int {

	var id_string string = ""

	for _, input_char := range input_string {
		if _, err := strconv.Atoi(string(input_char)); err == nil {
			id_string += string(input_char)
		} else {
			break
		}
	}

	id_int, _ := strconv.Atoi(id_string)

	return id_int

}

// Stole this score function from reddit
func getScore(score int) int {
	matches := score
	return int(math.Pow(2, float64(matches-1)))
}

func bulkAddScratchCards(scratchcards []string) []ScratchCard {

	var completed_cards []ScratchCard

	for _, scratch_card := range scratchcards {

		if scratch_card == "" {
			break
		}

		myScratchCard := NewScratchCard(scratch_card)

		completed_cards = append(completed_cards, myScratchCard)

	}

	return completed_cards

}

func NewScratchCard(card_info string) ScratchCard {

	counter := 0
	card_info = card_info[5:]
	card_id := getInt(card_info)
	strip_length := len(strconv.Itoa(card_id)) + 1
	card_info = card_info[strip_length:]

	cards_results := strings.Split(card_info, " | ")

	winning_numbers := strings.Split(cards_results[0], " ")
	guesses := strings.Split(cards_results[1], " ")

	winning_numbers_int := FilterEmptyConvert(winning_numbers)
	guesses_int := FilterEmptyConvert(guesses)

	for _, guess := range guesses_int {

		for _, winning_number := range winning_numbers_int {
			if winning_number == guess {
				counter += 1
			}
		}

	}

	converted_score := getScore(counter)

	s := ScratchCard{
		Id:             card_id,
		WinningNumbers: winning_numbers_int,
		Guesses:        guesses_int,
		Points:         converted_score,
	}

	return s

}