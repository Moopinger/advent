package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Schematic struct {
	Lines []string
	Parts []Part
}

type Part struct {
	//x:y
	Coords []int64
	Size   int64
	Value  int64
	//has special char in perim
	validPart bool
}

// will get the "length" of the int a sting starts with eg. 100...343....3.. will return 3
func getEngineLength(input_row string) int64 {

	var part_name string = ""

	for _, input_char := range input_row {
		if _, err := strconv.Atoi(string(input_char)); err == nil {
			part_name += string(input_char)
		} else {
			break
		}
	}

	return int64(len(part_name))

}

func isPointDigit(input_char string) bool {

	if _, err := strconv.Atoi(string(input_char)); err == nil {
		return true
	} else {
		return false
	}

}

func getEngineParts(input_row string, row_size int64, row_number int64, upper_row string, lower_row string) []Part {

	var parts []Part
	var x_coord int64 = 0

	//add padd for right char search

	for x_coord < row_size {

		if isPointDigit(string(input_row[x_coord])) == true {
			coords := []int64{0, 0}

			engine_length := getEngineLength(input_row[x_coord:])

			engineend_index := x_coord + engine_length
			value := input_row[x_coord:engineend_index]
			part_value, _ := strconv.Atoi(value)

			var left_char string = ""
			var left_most_position int64 = 0
			var right_most_position int64 = 0
			var right_char string = ""

			//x
			coords[0] = x_coord
			//y
			coords[1] = row_number

			//Check for scpecialchar in proximity
			//first char is a numm /
			if x_coord == 0 {
				left_most_position = 0
				left_char = "."
				right_most_position = 1 + engine_length
				right_char = string(input_row[right_most_position-1])
				//last char is num?
			} else if (x_coord + engine_length) >= (row_size - 1) {
				left_most_position = x_coord - 1
				left_char = string(input_row[left_most_position])
				right_most_position = row_size - 1
				right_char = "."
				//neith first nor last is  num
			} else {
				left_most_position = x_coord - 1
				left_char = string(input_row[left_most_position])
				right_most_position = x_coord + engine_length + 1
				right_char = string(input_row[right_most_position-1])
			}

			top_pattern := upper_row[left_most_position:right_most_position]
			lower_pattern := lower_row[left_most_position:right_most_position]

			//pattern := "Left Char: " + left_char + "\nTopPat: " + top_pattern + "\nLowePat: " + lower_pattern + "\nRight-Char: " + right_char
			//pattern := left_char + " || " + top_pattern + " || " + right_char + " || " + lower_pattern
			pattern := left_char + top_pattern + right_char + lower_pattern

			valid := false

			for _, char := range pattern {
				if string(char) != "." {
					valid = true
					break
				}

			}

			fmt.Printf("[+] Got the following id: %s for %d\n", pattern, part_value)

			p := Part{
				Coords:    coords,
				Size:      engine_length,
				Value:     int64(part_value),
				validPart: valid,
			}

			parts = append(parts, p)

			x_coord += engine_length

		} else {
			x_coord += 1
		}
	}

	return parts

}

func NewSchematic(schematic_lines []string) Schematic {

	var all_parts []Part

	total_lines := len(schematic_lines)

	for row_number, schematic_line := range schematic_lines {

		schematic_line_size := len(schematic_line)
		if schematic_line_size == 0 {
			break
		} else {
			//create fake Upper row for 0 and a fake lower for end
			upper_row := ""
			lower_row := ""
			switch row_number {
			case 0:
				upper_row = strings.Repeat(".", schematic_line_size)
				lower_row = schematic_lines[row_number+1]
				//deduct two here beacuse of new lines included in file
			case total_lines - 2:
				upper_row = schematic_lines[row_number-1]
				lower_row = strings.Repeat(".", schematic_line_size)

			default:
				upper_row = schematic_lines[row_number-1]
				lower_row = schematic_lines[row_number+1]
			}

			parts := getEngineParts(schematic_line, int64(schematic_line_size), int64(row_number), upper_row, lower_row)

			all_parts = append(all_parts, parts...)

		}
	}

	s := Schematic{

		Lines: schematic_lines,
		Parts: all_parts,
	}

	return s
}
