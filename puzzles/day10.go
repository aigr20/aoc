package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Day10A https://adventofcode.com/2020/day/10
func Day10A() {
	input, err := os.Open("inputs/day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)
	var jolts []int
	jolts = append(jolts, 0) // charging outlet is 0
	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			num = 0
		}
		jolts = append(jolts, num)
	}
	sort.Ints(jolts)

	var (
		diff1 int
		diff3 int
	)
	for i, jolt := range jolts {
		if i == len(jolts)-1 {
			diff3++ // device port difference is always 3
			break
		}

		switch jolts[i+1] - jolt {
		case 1:
			diff1++
			break
		case 3:
			diff3++
			break
		}
	}

	fmt.Println(diff1 * diff3)
}

func paths(pos int, jolts []int, memo map[int]int) int {
	for k := range memo {
		if pos == k {
			return memo[pos]
		}
	}
	if pos == len(jolts)-1 {
		return 1
	}

	var ways int

	for i := pos + 1; i < len(jolts); i++ {
		if jolts[i]-jolts[pos] > 3 {
			break
		}

		ans := paths(i, jolts, memo)
		ways += ans
		memo[i] = ans
	}

	return ways
}

// Day10B https://adventofcode.com/2020/day/10#part2
func Day10B() {
	input, err := os.Open("inputs/day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)
	var jolts []int

	jolts = append(jolts, 0)
	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			num = 0
		}
		jolts = append(jolts, num)
	}
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	fmt.Println(paths(0, jolts, make(map[int]int)))
}
