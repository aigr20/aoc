package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func isValid(preamble []int, considered int) bool {
	for i, preNum := range preamble {
		diff := considered - preNum
		for _, compNum := range preamble[i:] {
			if compNum == diff {
				return true
			}
		}
	}
	return false
}

func sum(arr []int) int {
	var sum int
	for _, num := range arr {
		sum += num
	}
	return sum
}

// Day9A https://adventofcode.com/2020/day/9
func Day9A() {
	input, err := os.Open("inputs/day9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)
	var (
		numbers          []int
		preambleMax      = 25
		firstPreamblePos = 0
	)
	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			num = 0
		}
		numbers = append(numbers, num)
	}

	for {
		preamble := numbers[firstPreamblePos:preambleMax]
		consideredNumbers := numbers[preambleMax:]

		if !isValid(preamble, consideredNumbers[0]) {
			fmt.Println(consideredNumbers[0])
			break
		}
		firstPreamblePos++
		preambleMax++
	}
}

// Day9B https://adventofcode.com/2020/day/9
func Day9B() {
	input, err := os.Open("inputs/day9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)

	var (
		numbers []int
		try     []int
		target  = 258585477
	)

	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			num = 0
		}
		numbers = append(numbers, num)
	}

	for i := range numbers {
		for _, tryAdd := range numbers[i:] {
			try = append(try, tryAdd)
			if sum(try) > target {
				try = make([]int, 0)
				break
			} else if sum(try) == target {
				sort.Ints(try)
				fmt.Println(try[0] + try[len(try)-1])
				return
			}
		}
	}
}
