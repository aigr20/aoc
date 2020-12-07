package puzzles

import (
	"aoc/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func validatePassword1(passDetails []string) bool {
	strRange := strings.Split(passDetails[0], "-")
	min, err := strconv.Atoi(strRange[0])
	if err != nil {
		log.Fatal(err)
	}

	max, err := strconv.Atoi(strRange[1])
	if err != nil {
		log.Fatal(err)
	}

	search := string(passDetails[1][0])
	testPassword := passDetails[2]

	occurences := strings.Count(testPassword, search)

	return (occurences <= max) && (occurences >= min)
}

func validatePassword2(passDetails []string) bool {
	strPositions := strings.Split(passDetails[0], "-")
	first, err := strconv.Atoi(strPositions[0])
	if err != nil {
		log.Fatal(err)
	}

	last, err := strconv.Atoi(strPositions[1])
	if err != nil {
		log.Fatal(err)
	}

	search := string(passDetails[1][0])
	testPassword := passDetails[2]

	passOk := false

	firstOk := string(testPassword[first-1]) == search
	lastOk := string(testPassword[last-1]) == search

	if firstOk && lastOk {
		passOk = false
	} else if firstOk || lastOk {
		passOk = true
	}

	return passOk
}

// Day2 https://adventofcode.com/2020/day/2
func Day2() {
	input, err := util.ReadInput("inputs/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	valids1 := 0
	valids2 := 0

	passwords := strings.Split(input, "\n")
	for _, password := range passwords {
		passInfoSlice := strings.Split(password, " ")
		pass1Ok := validatePassword1(passInfoSlice)
		if pass1Ok {
			valids1++
		}
		pass2Ok := validatePassword2(passInfoSlice)
		if pass2Ok {
			valids2++
		}
	}

	fmt.Println(valids1)
	fmt.Println(valids2)
}
