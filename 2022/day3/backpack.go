package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type BackPack struct {
	Id          int64
	ItemsCount  int64
	TopPouch    string
	BottomPouch string
	AllGear     string
	CommonItem  byte
}

// stolen from https://github.com/Nikscorp/advent_of_code_2022/blob/master/days/day3.go
func Score(bpitem byte) int64 {
	if bpitem >= 'a' && bpitem <= 'z' {
		return int64(bpitem - 'a' + 1)
	}
	return int64(bpitem - 'A' + 27)
}

func packBackPack(items string) (string, string, int64, byte, error) {

	var itemCount int64
	itemCount = int64(len(items))
	var common_item byte

	a := fmt.Sprintf("Length: %d", itemCount)

	if itemCount%2 != 0 || itemCount == 0 {
		return "", "", 0, 0, errors.New(a)
	}

	pack_amount := itemCount / 2
	toppack := items[:pack_amount]
	botpack := items[pack_amount:]

	for _, topitem := range toppack {

		if strings.Contains(botpack, string(topitem)) {

			common_item = byte(topitem)
			break

		}

	}

	return toppack, botpack, itemCount, common_item, nil
}

func NewBackPack(Id int64, items string) (BackPack, error) {

	TopPouch, BottomPouch, ItemsCount, CommonItem, err := packBackPack(items)

	if err != nil {
		log.Printf("[-] Cannot create backpack %s", err)
		return BackPack{}, err
	}

	e := BackPack{
		Id:          Id,
		ItemsCount:  ItemsCount,
		TopPouch:    TopPouch,
		BottomPouch: BottomPouch,
		AllGear:     items,
		CommonItem:  CommonItem,
	}
	return e, nil
}
