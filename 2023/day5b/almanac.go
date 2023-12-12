package main

import (
	"strconv"
	"strings"
)

type Almanac struct {
	Seeds    []SeedChunk
	Mappings []Mapping
	//Mappings   map[string]map[int]int
}

type SeedChunk struct {
	startInd  int
	chunkSize int
}

type Mapping struct {
	Name  string
	Rules []Rule
}

type Rule struct {
	destAddr int
	srcAddr  int
	Length   int
}

func NewSeedChunk(start int, size int) SeedChunk {
	s := SeedChunk{
		startInd:  start,
		chunkSize: size,
	}

	return s

}

func NewRule(dst int, src int, length int) Rule {
	r := Rule{
		destAddr: dst,
		srcAddr:  src,
		Length:   length,
	}

	return r
}

func (m *Mapping) Convertit(seed_id int) int {
	var result int

	for _, rule := range m.Rules {

		src_range_end := rule.srcAddr + rule.Length
		if (seed_id >= rule.srcAddr) && (seed_id < src_range_end) {

			difference := seed_id - rule.srcAddr

			result = rule.destAddr + difference

			break

		} else {
			result = seed_id
		}
	}

	return result

}

func NewMapping(name string, rules []Rule) Mapping {

	m := Mapping{
		Name:  name,
		Rules: rules,
	}

	return m
}

func NewAlmanac(almanac_data []string) Almanac {

	var mappings []Mapping
	var seedchunks []SeedChunk

	seeds_line := almanac_data[0]
	//remove "seeds: "
	seeds_line = seeds_line[7:]
	seeds_string := strings.Split(seeds_line, " ")
	seeds_input_count := len(seeds_string)

	for i := 0; i < seeds_input_count; i += 2 {
		seed_start_int, _ := strconv.Atoi(seeds_string[i])
		seed_size_int, _ := strconv.Atoi(seeds_string[i+1])

		mySeedChunk := NewSeedChunk(seed_start_int, seed_size_int)

		seedchunks = append(seedchunks, mySeedChunk)

	}

	almanac_data = almanac_data[2:]
	almanac_single_string := strings.Join(almanac_data, "\n") + "\n"
	almanac_conversions := strings.Split(almanac_single_string, "\n\n")
	//strip empty list item Thanks additional
	almanac_conversions = almanac_conversions[:len(almanac_conversions)-1]

	for _, almanac_conversion := range almanac_conversions {

		var rules []Rule

		almanac_conversion := strings.Split(almanac_conversion, "\n")

		mapping_name := almanac_conversion[0]

		mapping_length := len(almanac_conversion)

		for i := 1; i < mapping_length; i++ {
			//add this because of additional line read from text file
			conversion_values_string := strings.Split(almanac_conversion[i], " ")

			dest_range, _ := strconv.Atoi(conversion_values_string[0])
			src_range, _ := strconv.Atoi(conversion_values_string[1])
			range_length, _ := strconv.Atoi(conversion_values_string[2])

			myRule := NewRule(dest_range, src_range, range_length)

			rules = append(rules, myRule)

			//full := fmt.Sprintf("dest_range = %d src_range = %d chunk size is %d", dest_range, src_range, range_length)

			//modifications = append(modifications, full)
		}

		myMapping := NewMapping(mapping_name, rules)
		mappings = append(mappings, myMapping)

	}

	a := Almanac{
		Seeds:    seedchunks,
		Mappings: mappings,
	}
	return a

}
