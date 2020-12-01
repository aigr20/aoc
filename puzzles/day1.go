package puzzles

import (
	"aoc/util"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func part1(numbers []int) string {
	for _, i := range numbers {
		toFind := 2020 - i
		found := sort.SearchInts(numbers, toFind)
		if numbers[found] == toFind {
			return fmt.Sprintf("%v * %v = %v\n", i, toFind, i*toFind)
		}
	}

	return "Kunde inte lösas :("
}

func part2(numbers []int) string {
	for _, i := range numbers {
		for _, k := range numbers {
			toFind := 2020 - i - k
			found := sort.SearchInts(numbers, toFind)
			if numbers[found] == toFind {
				return fmt.Sprintf("%v * %v * %v = %v\n", i, k, toFind, i*k*toFind)
			}
		}
	}

	return "Kunde inte lösas :("
}

// Day1 https://adventofcode.com/2020/day/1
func Day1() {
	input, err := util.ReadInput("inputs/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(input, "\n")
	numbers := make([]int, 0)
	for _, i := range inputs {
		num, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, num)
	}
	sort.Ints(numbers)

	p1Ans := part1(numbers)
	p2Ans := part2(numbers)

	fmt.Printf("%s%s", p1Ans, p2Ans)
}
