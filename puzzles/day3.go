package puzzles

import (
	"aoc/util"
	"fmt"
	"log"
	"strings"
)

func treesOnPath(rows []string, xStep int, yStep int) int {
	repeatAfter := len(rows[0])
	trees := 0
	checkPos := 0

	for i := 0; i < len(rows); i += yStep {
		if string(rows[i][checkPos]) == "#" {
			trees++
		}
		if (checkPos + xStep) > repeatAfter-1 {
			checkPos = (checkPos + xStep) % repeatAfter
		} else {
			checkPos += xStep
		}
	}

	return trees
}

// Day3 https://adventofcode.com/2020/day/3
func Day3() {
	input, err := util.ReadInput("inputs/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(input, "\n")
	p1Trees := treesOnPath(rows, 3, 1)
	trees := treesOnPath(rows, 1, 1)
	trees *= p1Trees
	trees *= treesOnPath(rows, 5, 1)
	trees *= treesOnPath(rows, 7, 1)
	trees *= treesOnPath(rows, 1, 2)

	fmt.Printf("Träffade på %v träd i första delen\n", p1Trees)
	fmt.Printf("Träden på alla vägar multiplicerade blir %v\n", trees)
}
