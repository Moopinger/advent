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
			move_counter += 1
			node := myMap.Nodes[pointer]
			if string(direction) == "L" {
				pointer = node.Left
			} else {
				pointer = node.Right
			}
		}
	}
	fmt.Printf("[+]Z Solved. Took %d moves\n", move_counter)
}
