package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

type move struct {
	num  int
	from int
	to   int
}

func main() {
	stacks, moves := parseInput()
	partOne(stacks, moves)
	stacks, moves = parseInput()
	partTwo(stacks, moves)
}

func parseInput() ([][]string, []move) {
	bytes, _ := os.ReadFile("input")
	data := strings.Split(string(bytes), "\n\n")

	// first half
	crates := strings.Split(data[0], "\n")

	// Get last line to get the number of stacks
	totalRows := strings.Fields(crates[len(crates)-1])
	stacks := make([][]string, len(totalRows))

	for i := len(crates) - 2; i >= 0; i-- {
		for position := 1; position <= len(crates[i]); position += 4 {
			if crates[i][position] != ' ' {
				index := position / 4
				stacks[index] = append(stacks[index], string(crates[i][position]))
			}
		}
	}

	var moves []move
	lines := strings.Split(data[1], "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		moves = append(moves, move{
			num:  convertToInt(words[1]),
			from: convertToInt(words[3]) - 1,
			to:   convertToInt(words[5]) - 1,
		})
	}

	return stacks, moves
}

func partOne(stacks [][]string, moves []move) {
	for _, move := range moves {
		for i := 0; i < move.num; i++ {
			size := len(stacks[move.from])
			movingFrom := stacks[move.from]
			element := movingFrom[len(movingFrom)-1]
			stacks[move.from] = stacks[move.from][:size-1]
			stacks[move.to] = append(stacks[move.to], element)
		}
	}

	var output bytes.Buffer
	for _, stack := range stacks {
		length := len(stack)
		if length > 0 {
			output.WriteString(stack[length-1])
		}
	}
	fmt.Println("The answer to part 1 is", output.String())
}

func partTwo(stacks [][]string, moves []move) {
	for _, move := range moves {
		size := len(stacks[move.from])
		elementsMoved := stacks[move.from][size-move.num:]
		stacks[move.from] = stacks[move.from][:size-move.num]
		stacks[move.to] = append(stacks[move.to], elementsMoved...)
	}

	var output bytes.Buffer
	for _, stack := range stacks {
		length := len(stack)
		if length > 0 {
			output.WriteString(stack[length-1])
		}
	}
	fmt.Println("The answer to part 2 is", output.String())
}

func convertToInt(val string) int {
	intVal, _ := strconv.Atoi(val)
	return intVal
}
