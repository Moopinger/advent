package main

import (
	"fmt"
	"log"
	"os"
)

// Math magic - basicaly find the fist number that is divisible by all the move counts to reach a "Z" last char  see var []int divisnumbers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find the Least Common Multiple of two numbers
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to find the LCM of an array of numbers
func lcmArray(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
}

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

	var starterAddresses []string
	var divis_numbers []int

	//fmt.Println(myMap)
	fmt.Printf("[+] Loaded %d nodes and a map with pattern: %s\n", len(myMap.Nodes), myMap.Pattern)

	for _, node := range myMap.Nodes {
		if node.Name[2:] == "A" {
			starterAddresses = append(starterAddresses, node.Name)
		}
	}

	// for _, starter := range starterAddresses {
	// 	fmt.Printf("SS: %s\n", starter)
	// }
	//pointer := "AAA"

	for _, pointer := range starterAddresses {
		move_counter := 0
		for pointer[2:] != "Z" {

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
		divis_numbers = append(divis_numbers, move_counter)
		fmt.Printf("[+]%s found. Took %d moves\n", pointer, move_counter)
	}

	fmt.Print("Calculating number...\n")

	answer := lcmArray(divis_numbers)

	fmt.Printf("[+] Answer is: %d\n", answer)

}
