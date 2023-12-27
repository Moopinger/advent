package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getNextRow(input_row []int) []int {
	var nextRow []int

	for i := 1; i < len(input_row); i += 1 {
		difference := input_row[i] - input_row[i-1]
		nextRow = append(nextRow, difference)
	}

	return nextRow

}

func allZero(row []int) bool {
	for _, val := range row {
		if val != 0 {
			return false
		}
	}
	return true
}

func predictNext(rows [][]int) int {

	height := len(rows) - 1

	//appends last zero
	//rows[height] = append(rows[height], 0)
	rows[height] = append([]int{0}, rows[height]...)
	for i := (height - 1); i >= 0; i = i - 1 {
		predictedval := rows[i][0] - rows[i+1][0]
		rows[i] = append([]int{predictedval}, rows[i]...)
	}

	return rows[0][0]

}

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//remove trailing line
	data = data[:len(data)-1]
	challenge_inputs := strings.Split(string(data), "\n")

	next_answers := []int{}

	for _, challenge_input := range challenge_inputs {

		steps := []int{}
		rows := [][]int{}

		inputs := strings.Split(challenge_input, " ")

		for _, input := range inputs {
			int_val, _ := strconv.Atoi(input)
			steps = append(steps, int_val)
		}

		rows = append(rows, steps)

		for true {

			row := rows[len(rows)-1]
			next_row := getNextRow(row)
			rows = append(rows, next_row)

			if allZero(next_row) == true {
				break
			}
		}

		answer := predictNext(rows)
		fmt.Printf("The answer is... %d\n", answer)
		next_answers = append(next_answers, answer)
	}
	total := 0
	for _, val := range next_answers {
		total += val
	}
	fmt.Printf("[+]Total values are %d\n", total)
}
