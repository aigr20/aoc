package puzzles

import (
	"aoc/util"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var requiredFields = []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}

func validateField(field string, passport string) bool {
	fields := strings.Fields(passport)
	sort.Strings(fields)
	found := sort.SearchStrings(fields, field)
	data := strings.Split(fields[found], ":")[1]
	ok := false

	switch field {
	case "byr:":
		byr, err := strconv.Atoi(data)
		if err != nil {
			return false
		}
		ok = (byr >= 1920) && (byr <= 2002)
		break
	case "iyr:":
		iyr, err := strconv.Atoi(data)
		if err != nil {
			return false

		}
		ok = (iyr >= 2010) && (iyr <= 2020)
		break
	case "eyr:":
		eyr, err := strconv.Atoi(data)
		if err != nil {
			return false
		}
		ok = (eyr >= 2020) && (eyr <= 2030)
		break
	case "hgt:":
		ok = strings.HasSuffix(data, "in") || strings.HasSuffix(data, "cm")
		if ok == false {
			return false
		}
		unit := data[len(data)-2:]
		hgt, err := strconv.Atoi(data[:len(data)-2])
		if err != nil {
			return false
		}

		if unit == "in" {
			ok = (hgt >= 59) && (hgt <= 76)
		} else if unit == "cm" {
			ok = (hgt >= 150) && (hgt <= 193)
		}
		break
	case "hcl:":
		var err error
		regex := "^#[a-f|A-F|0-9]+$"
		ok, err = regexp.MatchString(regex, data)
		if err != nil {
			log.Fatal(err)
		}
		break
	case "ecl:":
		validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, color := range validColors {
			ok = data == color
			if ok {
				break
			}
		}
		break
	case "pid:":
		if len(data) != 9 {
			return false
		}
		_, err := strconv.Atoi(data)
		if err != nil {
			return false
		}
		ok = true
	}

	return ok
}

func countValidatedPassports(passports []string) int {
	validPassports := 0

passport_loop:
	for _, passport := range passports {
		for _, field := range requiredFields {
			if strings.Index(passport, field) < 0 {
				continue passport_loop
			}
			fieldOk := validateField(field, passport)
			if fieldOk == false {
				continue passport_loop
			}
		}
		validPassports++
	}

	return validPassports
}

func countValidPassports(passports []string) int {
	validPassports := 0

passport_loop:
	for _, passport := range passports {
		for _, field := range requiredFields {
			if strings.Index(passport, field) < 0 {
				continue passport_loop
			}
		}
		validPassports++
	}

	return validPassports
}

// Day4 https://adventofcode.com/2020/day/4
func Day4() {
	input, err := util.ReadInput("inputs/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	passports := strings.Split(input, "\n\n")
	validPassports := countValidPassports(passports)
	validatedPassports := countValidatedPassports(passports)

	fmt.Printf("Det finns %v giltiga pass\n", validPassports)
	fmt.Printf("Det finns %v pass som uppfyller valideringskraven\n", validatedPassports)
}
