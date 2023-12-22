package main

import (
	"slices"
	"strconv"
	"strings"
)

type CamelCompetition struct {
	Hands []Hand
}

type Hand struct {
	Cards  []Card
	Bid    int
	Type   int
	Points int

	//0 unknown
	//1 High card
	//2 1pair
	//3 2pair
	//4	3 of a kind
	//5 full house
	//6 4 of a kind
	//7 5 of a kind

}

func calculateHandType(cards []Card) int {

	// counting sort https://www.geeksforgeeks.org/counting-sort/ initial steps
	groups := make(map[Card]int)
	for _, card := range cards {
		groups[card]++
	}

	distinctCount := len(groups)
	maxGroup := 0
	for _, count := range groups {
		if count > maxGroup {
			maxGroup = count
		}
	}

	switch {
	case distinctCount == 1 && maxGroup == 5:
		return 7
	case distinctCount == 2 && maxGroup == 4:
		return 6
	case distinctCount == 2 && maxGroup == 3:
		return 5
	case distinctCount == 3 && maxGroup == 3:
		return 4
	case distinctCount == 3 && maxGroup == 2:
		return 3
	case distinctCount == 4 && maxGroup == 2:
		return 2
	case distinctCount == 5 && maxGroup == 1:
		return 1
	default:
		return 0
	}
}

type Card struct {
	Points int
	Value  string //0..9 T J Q K A
}

func NewCamelCompetition(data []byte) CamelCompetition {

	var cards string
	var bid int
	var hands []Hand

	data_lines := strings.Split(string(data), "\n")

	for _, data_line := range data_lines {

		values := strings.Split(data_line, " ")
		cards = values[0]
		bid, _ = strconv.Atoi(values[1])

		myHand := newHand(cards, bid)
		hands = append(hands, myHand)
	}

	sortHands(hands)

	c := CamelCompetition{
		Hands: hands,
	}

	return c
}

// found this beautiful sort online and adapted here
func sortHands(hands []Hand) []Hand {

	slices.SortFunc(hands, func(a, b Hand) int {
		if a.Type == b.Type {
			for i := 0; i < len(a.Cards); i++ {
				if a.Cards[i].Points == b.Cards[i].Points {
					continue
				}
				return a.Cards[i].Points - b.Cards[i].Points
			}
		}
		return a.Type - b.Type
	})

	return hands
}

// debug feature
func (h *Hand) printHand() string {
	hand := ""

	for _, card := range h.Cards {
		hand = hand + string(card.Value)
	}

	return hand

}

// func getHandScore(hand []Card) int {

// }

func newHand(cards_string string, bid int) Hand {

	var cards []Card
	score := 0

	for _, card := range cards_string {

		card_s := string(card)

		myCard := newCard(card_s)
		score += myCard.Points
		cards = append(cards, myCard)
	}

	hand_type := calculateHandType(cards)

	h := Hand{
		Cards:  cards,
		Type:   hand_type,
		Bid:    bid,
		Points: score,
	}

	return h
}

func newCard(card_v string) Card {
	number := 0
	number, err := strconv.Atoi(card_v)

	if err != nil {
		switch card_v {
		case "T":
			number = 10
		case "J":
			number = 11
		case "Q":
			number = 12
		case "K":
			number = 13
		case "A":
			number = 14
		default:
			number = 0
		}
	}

	c := Card{
		Points: number,
		Value:  card_v,
	}

	return c
}
