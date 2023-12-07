package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var contents []string
	var backpacks []BackPack
	var bagid int64 = 1
	var backpack BackPack
	var travelgroups []TravelGroup
	var total int64 = 0
	var team_total int64 = 0

	dat, err := os.ReadFile("./list.txt")

	if err != nil {
		log.Fatal(err)
	}

	contents = strings.Split(string(dat), "\n")

	for _, items := range contents {

		backpack, err = NewBackPack(bagid, items)
		if err != nil {
			break
		}

		total += Score(backpack.CommonItem)

		backpacks = append(backpacks, backpack)
		//fmt.Println(backpack)

	}

	size := 3
	var j int
	for i := 0; i < len(backpacks); i += size {
		j += size
		if j > len(backpacks) {
			j = len(backpacks)
		}

		MyGroup, err := NewTravelGroup(1, backpacks[i:j])

		if err != nil {
			log.Fatalf("aaaaa: %s", err)

		}

		travelgroups = append(travelgroups, MyGroup)

	}

	for _, travelgroup := range travelgroups {
		team_total += Score(travelgroup.Badge)
	}
	fmt.Printf("[+] Your total is: %d \r\n", team_total)

}
