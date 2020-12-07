package puzzles

import (
	"aoc/util"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
)

type seat struct {
	row    int
	column int
}

func lowerHalf(min int, max int) int {
	return int(math.Floor((float64(max) + float64(min)) / 2))
}

func upperHalf(min int, max int) int {
	return int(math.Ceil(float64(max)/2 - float64(min)/2))
}

func calculateSeat(boardingPass string) seat {
	minRow := 0
	maxRow := 127
	rowSpec := boardingPass[:7]
	colSpec := boardingPass[7:]

	var calculatedSeat seat
	for _, spec := range rowSpec {
		switch spec {
		case 'F':
			maxRow = lowerHalf(minRow, maxRow)
			break
		case 'B':
			minRow += upperHalf(minRow, maxRow)
			break
		}
	}

	if minRow == maxRow {
		calculatedSeat.row = maxRow
	} else {
		log.Fatalf("Raduträkningen misslyckades.\nminRow: %v\nmaxRow: %v\n", minRow, maxRow)
	}

	minCol := 0
	maxCol := 7
	for _, spec := range colSpec {
		switch spec {
		case 'L':
			maxCol = lowerHalf(minCol, maxCol)
			break
		case 'R':
			minCol += upperHalf(minCol, maxCol)
			break
		}
	}

	if minCol == maxCol {
		calculatedSeat.column = maxCol
	} else {
		log.Fatalf("Kolumnuträkningen misslyckades.\nminCol: %v\nmaxCol: %v\n", minCol, maxRow)
	}

	return calculatedSeat
}

// Day5 https://adventofcode.com/2020/day/5
func Day5() {
	input, err := util.ReadInput("inputs/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	boardingPasses := strings.Split(input, "\n")

	highestID := 0
	allIDs := make([]int, 0)
	myID := 0
	for _, pass := range boardingPasses {
		passSeat := calculateSeat(pass)
		id := passSeat.row*8 + passSeat.column
		if id > highestID {
			highestID = id
		}

		allIDs = append(allIDs, id)
	}

	sort.Ints(allIDs)

	for i, id := range allIDs {
		if !(allIDs[i+1] == id+1) {
			myID = id + 1
			break
		}
	}

	fmt.Println(highestID)
	fmt.Println(myID)
}
