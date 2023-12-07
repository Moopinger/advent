package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Calibration struct {
	AllItems     string
	IntegerItems []int
	Sumof        int
}

func getInts(items string) ([]int, error) {

	var integers []int

	if items == "" {
		return integers, errors.New("getInts: Empty Items list....")
	}

	for _, item := range items {
		if digit, err := strconv.Atoi(string(item)); err == nil {
			integers = append(integers, digit)
		}
	}
	return integers, nil
}

func getSumof(items []int) (int, error) {

	//We use a string to temp store the combined values
	var sum_string = ""

	if len(items) < 1 {
		return -1, errors.New("Sumof: empty list, nothing to add")
	} else {
		//convert to
		sum_string = strconv.Itoa(items[0]) + strconv.Itoa(items[len(items)-1])
		// sum += items[0]
		// sum += items[len(items)-1]

		//convert back to int
		sum_int, err := strconv.Atoi(sum_string)
		if err != nil {
			return -1, errors.New("Could not convert value to int")
		}
		return sum_int, nil
	}

}

func replaceWord(items string) (string, error) {

	wordnumbers := map[string]string{"one": "o1e", "two": "t2o", "three": "t3e", "four": "f4r", "five": "f5e", "six": "s6x", "seven": "s7n", "eight": "e8t", "nine": "n9e"}

	for numword, num := range wordnumbers {
		items = strings.ReplaceAll(items, numword, num)
	}

	return items, nil
}

func NewCallibration(items string) Calibration {

	items, err := replaceWord(items)

	if err != nil {
		fmt.Println("f0 failed:", err)
	}

	integers, err := getInts(items)

	if err != nil {
		fmt.Println("f1 failed:", err)
	}

	sumof, err := getSumof(integers)

	if err != nil {
		fmt.Println("f2 failed:", err)
	}

	c := Calibration{
		AllItems:     items,
		IntegerItems: integers,
		Sumof:        sumof,
	}

	return c
}
