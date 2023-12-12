package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	//var answers []int

	//keep track of lowest value since we no longer store the results in an array
	lowest := 1000000000

	dat, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	almanac_lines := strings.Split(string(dat), "\n")

	myAlmanac := NewAlmanac(almanac_lines)

	//fmt.Println(myAlmanac)

	for _, seed_range := range myAlmanac.Seeds {

		//fmt.Printf("[+] Seed start Value is: %d and size is: %d\n", seed_range.startInd, seed_range.chunkSize)
		end_length := seed_range.startInd + seed_range.chunkSize

		for i := seed_range.startInd; i < end_length; i++ {

			seed_value := i
			//fmt.Printf("[+] Seed start value is %d", seed_value)
			for _, mapping := range myAlmanac.Mappings {
				seed_value = mapping.Convertit(seed_value)
			}
			if seed_value < lowest {
				fmt.Printf("[+]New Lowest found: %d\n", seed_value)
				lowest = seed_value
			}
			//fmt.Printf("[+] Seed final value is %d\n\n", seed_value)
		}

		// sort.Ints(answers)

	}
	fmt.Printf("Smallest value is: %d\n", lowest)
}
