package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("Elf Logger 5000\n")

	var list []string
	var elves []Elf
	var backpack []int64
	var elfid int64

	dat, err := os.ReadFile("./list.txt")

	if err != nil {
		log.Fatal(err)
	}

	list = strings.Split(string(dat), "\n")

	elfid = 1

	for _, item_calories := range list {
		if item_calories != "" {

			i, err := strconv.ParseInt(item_calories, 10, 64)
			if err != nil {
				panic(err)
			}
			backpack = append(backpack, i)
		} else {
			myElf := NewElf(elfid, backpack)
			elves = append(elves, myElf)
			elfid += 1
			backpack = nil
		}
	}

	check(err)

	//Sort them
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].Calories > elves[j].Calories
	})

	fmt.Printf("[+]The top three are carrying: %v\n", elves[0].Calories+elves[1].Calories+elves[2].Calories)

}
