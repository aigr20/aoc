package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	action   string
	val      int
	executed bool
}

// Day8A https://adventofcode.com/2020/day/8
func Day8A() {
	input, err := os.Open("inputs/day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)
	var (
		program     []instruction
		accumulator int
		pos         int // current instruction location
	)

	for sc.Scan() {
		instr := strings.Fields(sc.Text())
		number, err := strconv.Atoi(instr[1])
		if err != nil {
			log.Fatal(err)
		}
		program = append(program, instruction{
			action:   instr[0],
			val:      number,
			executed: false,
		})
	}

	for {
		if program[pos].executed {
			fmt.Println(accumulator)
			return
		}

		program[pos].executed = true
		switch program[pos].action {
		case "acc":
			accumulator += program[pos].val
			break
		case "jmp":
			pos += program[pos].val - 1
			break
		}
		pos++
	}
}

// Day8B https://adventofcode.com/2020/day/8
func Day8B() {
	input, err := os.Open("inputs/day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)
	var program []instruction
	for sc.Scan() {
		instr := strings.Fields(sc.Text())
		number, err := strconv.Atoi(instr[1])
		if err != nil {
			log.Fatal(err)
		}
		program = append(program, instruction{
			action:   instr[0],
			val:      number,
			executed: false,
		})
	}

	for pos := range program {
		if program[pos].action == "acc" {
			continue
		}

		programCopy := make([]instruction, len(program))
		copy(programCopy, program)
		programCopy[pos].action = swap(programCopy[pos].action)

		booted, accumulator := tryBoot(programCopy)
		if booted {
			fmt.Println(accumulator)
			return
		}
	}
}

func tryBoot(program []instruction) (bool, int) {
	var (
		pos         int
		accumulator int
	)

	for {
		if pos == len(program) {
			return true, accumulator
		}
		if program[pos].executed {
			return false, accumulator
		}

		program[pos].executed = true
		switch program[pos].action {
		case "acc":
			accumulator += program[pos].val
			break
		case "jmp":
			pos += program[pos].val - 1
			break
		}
		pos++
	}
}

func swap(action string) string {
	if action == "jmp" {
		return "nop"
	}
	return "jmp"
}
