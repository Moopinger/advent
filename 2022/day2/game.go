package main

import (
	"errors"
	"log"
)

type Game struct {
	Id     int64
	OpHand string
	MyHand string
	Score  int64
	Result string
}

func ProcEntry(entry string) (string, string, error) {

	var myhand string
	var ophand string

	if entry == "" {
		return "", "", errors.New("Cannot read")
	} else {
		ophand = entry[0:1]
		myhand = entry[len(entry)-1:]
	}

	return ophand, myhand, nil

}

func GetScore(ophand string, myhand string) (int64, string, error) {
	if len(ophand) != 1 || len(myhand) != 1 {
		return 0, "", errors.New("Strings too big to be valid moves")
	} else {

		var score int64 = 0
		var result string

		//convert our hand to abc style
		if myhand == "X" {
			myhand = "A"
			score = 1
		} else if myhand == "Y" {
			myhand = "B"
			score = 2
		} else if myhand == "Z" {
			myhand = "C"
			score = 3
		} else {
			return 0, "", errors.New("Not a valid player hand provided....")
		}

		//draw
		if myhand == ophand {
			score += 3
			result = "DRAW"
		} else if myhand == "A" {
			if ophand == "B" {
				result = "LOSE"
			} else if ophand == "C" {
				score += 6
				result = "WIN"
			}
		} else if myhand == "B" {
			if ophand == "A" {
				score += 6
				result = "WIN"
			} else if ophand == "C" {
				result = "LOSE"
			}
		} else if myhand == "C" {
			if ophand == "A" {
				result = "LOSE"
			} else if ophand == "B" {
				score += 6
				result = "WIN"
			}
		}

		return score, result, nil

	}

}

func NewGame(id int64, entry string) Game {

	var score int64
	var result string

	ophand, myhand, err := ProcEntry(entry)
	if err != nil {
		log.Printf("[-]cannot proc this entry")
	}

	score, result, err = GetScore(ophand, myhand)

	if err != nil {
		log.Fatal("[-]cannot get score for this entry....")
	}

	g := Game{
		Id:     id,
		OpHand: ophand,
		MyHand: myhand,
		Score:  score,
		Result: result,
	}

	return g
}
