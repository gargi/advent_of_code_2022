package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	tree := parseInput()
	count := partOne(tree)
	fmt.Println("Output of part one is", count)
	maxScore := partTwo(tree)
	fmt.Println("Output of part two is", maxScore)
}

func parseInput() [][]int {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	var data [][]int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		var row []int
		for _, num := range line {
			intVal, _ := strconv.Atoi(num)
			row = append(row, intVal)
		}
		data = append(data, row)

	}
	return data
}

func partOne(tree [][]int) int {
	count := 0

	maxRow := len(tree)
	maxCol := len(tree[0])

	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxCol; j++ {
			// edge rows
			if i == 0 || j == 0 || i == maxRow-1 || j == maxCol-1 {
				count += 1
				continue
			}
			current := tree[i][j]
			if isVisibleHorizontal(current, tree[i][:j]) || isVisibleHorizontal(current, tree[i][j+1:]) || isVisibleVertical(tree, current, i, j) {
				count += 1
			}
		}
	}
	return count
}

func isVisibleHorizontal(current int, row []int) bool {
	for _, tree := range row {
		if current <= tree {
			return false
		}
	}
	return true
}

func isVisibleVertical(tree [][]int, current, row, col int) bool {
	var above []int
	for i := 0; i < row; i++ {
		above = append(above, tree[i][col])
	}
	maxNeighbor := getMaxNeighbor(above)
	if current > maxNeighbor {
		return true
	}

	var below []int
	for i := row + 1; i < len(tree); i++ {
		below = append(below, tree[i][col])
	}
	maxNeighbor = getMaxNeighbor(below)
	if current > maxNeighbor {
		return true
	}

	return false
}

func getMaxNeighbor(neighbors []int) int {
	maxNeighbor := math.MinInt
	for _, neighbor := range neighbors {
		if neighbor > maxNeighbor {
			maxNeighbor = neighbor
		}
	}
	return maxNeighbor
}

func partTwo(tree [][]int) int {
	maxScore := math.MinInt
	maxRow := len(tree)
	maxCol := len(tree[0])

	for i := 1; i < maxRow-1; i++ {
		for j := 1; j < maxCol-1; j++ {
			current := tree[i][j]
			// reverse as the view from the tree would be reverse
			leftScore := highestHorizontal(current, reverse(tree[i][:j]))
			rightScore := highestHorizontal(current, tree[i][j+1:])
			verticalScore := highestVertical(tree, current, i, j)
			score := leftScore * rightScore * verticalScore
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func highestHorizontal(current int, row []int) int {
	count := 0
	for _, tree := range row {
		count += 1
		if tree >= current {
			break
		}
	}
	return count
}

func highestVertical(tree [][]int, current, row, col int) int {
	var above []int
	for i := 0; i < row; i++ {
		above = append(above, tree[i][col])
	}
	aboveView := reverse(above)
	aboveScore := 0
	for i := 0; i < len(aboveView); i++ {
		aboveScore += 1
		if current <= aboveView[i] {
			break
		}
	}

	var below []int
	for i := row + 1; i < len(tree); i++ {
		below = append(below, tree[i][col])
	}
	belowScore := 0
	for i := 0; i < len(below); i++ {
		belowScore += 1
		if current <= below[i] {
			break
		}
	}
	return aboveScore * belowScore
}

func reverse(numbers []int) []int {
	var output []int
	for i := len(numbers) - 1; i >= 0; i-- {
		output = append(output, numbers[i])
	}
	return output
}
