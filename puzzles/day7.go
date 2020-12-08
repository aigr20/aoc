package puzzles

import (
	"aoc/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type bags map[string][]string

var (
	holders          = make([]string, 0)
	bagContentAmount = make(map[string]int)
	day7part2        = 0
)

func goldTotal(types bags, start string) {
	for _, content := range types[start] {
		for i := 1; i < bagContentAmount[content]; i++ {
			day7part2 += bagContentAmount[content] - 1
			goldTotal(types, content)
		}
	}
}

func findHolders(types bags, searchFor string) {
	for bagType, bagContent := range types {
		for _, content := range bagContent {
			if content == searchFor {
				holders = append(holders, bagType)
				findHolders(types, bagType)
			}
		}
	}
}

// Day7 https://adventofcode.com/2020/day/7
func Day7() {
	input, err := util.ReadInput("inputs/day7_example.txt")
	if err != nil {
		log.Fatal(err)
	}

	rules := strings.Split(input, "\n")
	bagTypes := make(bags)
	for _, rule := range rules {
		rule = strings.TrimSuffix(rule, ".")
		ruleParts := strings.Split(rule, " contain ")
		bagColor := strings.TrimSuffix(ruleParts[0], " bags")
		bagTypes[bagColor] = make([]string, 0)
		contents := strings.Split(ruleParts[1], ", ")

		for _, bag := range contents {
			legalContent := strings.Join(strings.Split(bag, " ")[1:3], " ")
			contentAmount, err := strconv.Atoi(strings.Split(bag, " ")[0])
			if err != nil {
				if bag != "no other bags" {
					log.Fatal(err)
				} else {
					contentAmount = 0
				}
			}
			bagContentAmount[bagColor] += contentAmount
			bagTypes[bagColor] = append(bagTypes[bagColor], legalContent)
		}
	}

	findHolders(bagTypes, "shiny gold")
	alreadyCounted := make([]string, 0)
	found := 0

outer:
	for _, color := range holders {
		for _, counted := range alreadyCounted {
			if color == counted {
				continue outer
			}
		}
		alreadyCounted = append(alreadyCounted, color)
		found++
	}
	fmt.Println(found)
	goldTotal(bagTypes, "shiny gold")
	fmt.Println(day7part2)
}
