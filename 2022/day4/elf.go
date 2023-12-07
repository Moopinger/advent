package main

import (
	"log"
)

type Elf struct {
	Id       int64
	Items    []int64
	Calories int64
}

func AddCalories(items []int64) (int64, error) {

	var total int64
	total = 0

	for _, calories := range items {

		total += calories
	}
	return total, nil
}

func NewElf(Id int64, Items []int64) Elf {

	calories_total, err := AddCalories(Items)

	if err != nil {
		log.Fatal("wrong")
	}

	e := Elf{
		Id:       Id,
		Items:    Items,
		Calories: calories_total,
	}
	return e
}
