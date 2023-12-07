package main

import (
	"errors"
	"strings"
)

type TravelGroup struct {
	Id        int64
	BackPacks []BackPack
	Badge     byte
}

func NewTravelGroup(Id int64, Backpacks []BackPack) (TravelGroup, error) {

	var itemBadge byte
	if len(Backpacks) != 3 {
		//fmt.Printf(">: %d \n", len(Backpacks))
		return TravelGroup{}, errors.New("Only 3 backpacks at a time..., you provided")
	}

	for _, item := range Backpacks[0].AllGear {

		if strings.Contains(Backpacks[1].AllGear, string(item)) && strings.Contains(Backpacks[2].AllGear, string(item)) {

			//convert rune to byte
			itemBadge = byte(item)
			//fmt.Printf("[+] Calculated group is: %s \r\n", string(item))
			break
		}

	}

	tg := TravelGroup{
		Id:        Id,
		BackPacks: Backpacks,
		Badge:     itemBadge,
	}

	return tg, nil

}
