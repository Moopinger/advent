package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var total int64 = 0

	dat, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	schematic_lines := strings.Split(string(dat), "\n")

	mySchematic := NewSchematic(schematic_lines)

	for _, part := range mySchematic.Parts {

		fmt.Println(part)

		if part.isGear == true {
			total += part.Value
		}
	}

	//fmt.Printf("[+] Total is %d", total)

}
