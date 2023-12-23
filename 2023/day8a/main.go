package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Read file into a byte slice
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//remove trailing line
	data = data[:len(data)-1]
	//boom
	myMap := NewMap(data)

	//fmt.Println(myMap)
	fmt.Printf("[+] Loaded %d nodes and a map with pattern: %s\n", len(myMap.Nodes), myMap.Pattern)

	pointer := "AAA"
	move_counter := 0

	for pointer != "ZZZ" {

		for _, direction := range myMap.Pattern {
			for _, node := range myMap.Nodes {
				if node.Name == pointer {
					if string(direction) == "L" {
						move_counter += 1
						pointer = node.Left
						break
					} else {
						move_counter += 1
						pointer = node.Right
						break
					}
				}
			}

		}

	}

	fmt.Printf("[+]Z Solved. Took %d moves\n", move_counter)

}
