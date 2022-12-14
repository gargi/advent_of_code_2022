package main

import (
	"fmt"
	"math"
	"os"
)

type Coord struct {
	row, col int
}

func main() {
	minDistance := math.MaxInt
	input, startCoord, endCoord := parseInput('S', 'E')
	starts := findStarts('a')
	distance := bfs(input, startCoord, endCoord)
	fmt.Println("Part 1 :", distance)

	minDistance = min(minDistance, distance)
	for _, start := range starts {
		distance = bfs(input, start, endCoord)
		if distance != -1 { 
		minDistance = min(minDistance, bfs(input, start, endCoord))
	}
	}
	fmt.Println("Part 2 :", minDistance)
}

func parseInput(startPoint, endpoint rune) (grid []string, startCoords Coord, endCoords Coord) {
	bytes, _ := os.ReadFile("input")
	input := string(bytes)
	row := 0
	col := 0
	line := ""
	start := Coord{}
	end := Coord{}
	var lines []string
	for _, c := range input {
		if c == '\n' {
			lines = append(lines, line)
			line = ""
			row += 1
			col = 0
			continue
		}
		if c == startPoint {
			start.row = row
			start.col = col
			c = 'a'
		}
		if c == endpoint {
			end.row = row
			end.col = col
			c = 'z'
		}
		col += 1
		line += string(c)
	}
	lines = append(lines, line)
	return lines, start, end
}


func findStarts(startPoint rune) []Coord {
	var starts []Coord
	bytes, _ := os.ReadFile("input")
	input := string(bytes)
	row := 0
	col := 0
	line := ""
	var lines []string
	for _, c := range input {
		if c == '\n' {
			lines = append(lines, line)
			line = ""
			row += 1
			col = 0
			continue
		}
		if c == startPoint {
			starts = append(starts, Coord{row, col})
		}
		col += 1
		line += string(c)
	}
	lines = append(lines, line)
	return starts
}


func isValid(lines []string, coord Coord, from byte) bool {
	if coord.row < 0 || coord.row >= len(lines) || coord.col < 0 || coord.col >= len(lines[0]) {
		return false
	}
	to := lines[coord.row][coord.col]

	return int(to)-int(from) <= 1
}

func getNeighbors(lines []string, c Coord, currVal byte) []Coord {
	var neighbors []Coord
	left := Coord{c.row - 1, c.col}
	right := Coord{c.row + 1, c.col}
	up := Coord{c.row, c.col - 1}
	down := Coord{c.row, c.col + 1}

	if isValid(lines, left, currVal){
		neighbors = append(neighbors, left)
	}

	if isValid(lines, right, currVal){
		neighbors = append(neighbors, right)
	}

	if isValid(lines, up, currVal){
		neighbors = append(neighbors, up)
	}

	if isValid(lines, down, currVal){
		neighbors = append(neighbors, down)
	}

	return neighbors
}

func bfs(lines []string, start Coord, end Coord) (count int) {
	visited := make(map[Coord]int)
	var q []Coord
	var curr Coord

	q = append(q, start)
	visited[start] = 0

	for len(q) > 0 {
		curr = q[0]

		q = q[1:]
		v := visited[curr]

		neighbors := getNeighbors(lines, curr, lines[curr.row][curr.col])
		for _, n := range neighbors {
			if n.row == end.row && n.col == end.col {
				return v + 1
			}
			_, ok := visited[n]
			if !ok {
				visited[n] = v + 1
				q = append(q, n)
			}
		}
	}
	return -1
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}