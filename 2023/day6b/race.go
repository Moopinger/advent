package main

import (
	"errors"
	"regexp"
	"strconv"
)

type Competition struct {
	Races []Race
}

type Race struct {
	Time       int
	RecordDist int
}

func (r *Race) RaceRecord(ignit_time int) bool {

	if ignit_time >= r.Time || ignit_time == 0 {
		return false
	}

	available_time := r.Time - ignit_time
	//ignit time is our accelration
	dist_traveled := available_time * ignit_time
	//Get distance traveled

	if dist_traveled <= r.RecordDist {
		return false
	}

	return true
}

func newRace(time int, record_dist int) Race {
	r := Race{
		Time:       time,
		RecordDist: record_dist,
	}

	return r
}

// modify this to achive challenge b values
func convertToInt(lines []string) int {

	var temp_string_value string

	for _, line := range lines {
		temp_string_value = temp_string_value + line
	}
	line_converted, _ := strconv.Atoi(temp_string_value)

	return line_converted
}

func NewCompetition(lines []string) (Competition, error) {

	var races []Race
	//turns out regex is very useful here..my old way looks insane now
	re := regexp.MustCompile(`\d+`)
	times_str := re.FindAllString(lines[0], -1)
	record_distance_str := re.FindAllString(lines[1], -1)
	var myComp Competition

	if len(times_str) != len(record_distance_str) {
		err := errors.New("Different number of times and distances")
		return myComp, err
	}

	times_int := convertToInt(times_str)
	record_distance_int := convertToInt(record_distance_str)

	//fmt.Printf("Time is: %d and the world record is %d\n", time, record_distance_int[ind])
	myRace := newRace(times_int, record_distance_int)
	races = append(races, myRace)

	c := Competition{
		Races: races,
	}

	return c, nil

}
