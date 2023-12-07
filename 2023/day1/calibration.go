package main

import (
	"errors"
	"fmt"
	"strconv"
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

func NewCallibration(items string) Calibration {

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
