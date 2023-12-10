package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	var answers []int

	dat, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	almanac_lines := strings.Split(string(dat), "\n")

	myAlmanac := NewAlmanac(almanac_lines)

	//fmt.Println(myAlmanac)

	for _, seed_id := range myAlmanac.Seeds {

		fmt.Printf("[+] Seed Value is: %d\n", seed_id)
		for _, mapping := range myAlmanac.Mappings {

			seed_id = mapping.Convertit(seed_id)

			//fmt.Println(mapping)
			//fmt.Printf("converted value for %s is: %d\n", mapping.Name, seed_id)

		}

		//fmt.Printf("converted value is: %d\n", seed_id)
		answers = append(answers, seed_id)

	}
	sort.Ints(answers)
	fmt.Printf("Smallest value is: %d\n", answers[0])

}
