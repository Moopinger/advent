package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var total int64 = 0

	var calibrations []Calibration
	dat, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	calibration_inputs := strings.Split(string(dat), "\n")

	for _, calib_input := range calibration_inputs {

		if calib_input == "" {
			break
		}
		calib := NewCallibration(calib_input)
		calibrations = append(calibrations, calib)
		total += int64(calib.Sumof)
		fmt.Println(calib)
	}

	fmt.Printf("Total is: %d\n", total)

}
