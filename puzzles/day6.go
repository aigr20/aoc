package puzzles

import (
	"aoc/util"
	"fmt"
	"log"
	"strings"
)

func countUnanimousYes(answers [][]string) int {
	yesCounts := make(map[string]int)

	for _, answeredQuestions := range answers {
		for _, answer := range answeredQuestions {
			yesCounts[answer]++
		}
	}

	unanimousYes := 0
	for _, questionYesCount := range yesCounts {
		if questionYesCount == len(answers) {
			unanimousYes++
		}
	}
	return unanimousYes
}

func countGroupYes(answers [][]string) int {
	uniqueAnswers := make([]string, 0)
	for _, answeredQuestions := range answers {
		for _, answer := range answeredQuestions {
			alreadyAnswered := false
			for _, answered := range uniqueAnswers {
				if answer == answered {
					alreadyAnswered = true
					break
				}
			}
			if !alreadyAnswered {
				uniqueAnswers = append(uniqueAnswers, answer)
			}
		}
	}

	return len(uniqueAnswers)
}

// Day6 https://adventofcode.com/2020/day/6
func Day6() {
	input, err := util.ReadInput("inputs/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	groups := strings.Split(input, "\n\n")
	totalYes := 0
	unanimousYes := 0
	for _, group := range groups {
		answers := strings.Split(group, "\n")
		groupAnswers := make([][]string, 0)

		for _, personAnswers := range answers {
			groupAnswers = append(groupAnswers, strings.Split(personAnswers, ""))
		}

		totalYes += countGroupYes(groupAnswers)
		unanimousYes += countUnanimousYes(groupAnswers)
	}

	fmt.Printf("Totalt var det %v ja-svar\n", totalYes)
	fmt.Printf("Total var det %v eniga ja-svar\n", unanimousYes)
}
