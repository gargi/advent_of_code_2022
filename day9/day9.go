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

const n = 10

func main() {
	partOne()
	partTwo()
}

type Position struct {
	x int
	y int
}

func partOne() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	seen := make(map[Position]bool)
	seen[Position{0, 0}] = true
	head := Position{0, 0}
	tail := Position{0, 0}
	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		steps, _ := strconv.Atoi(instruction[1])
		for i := 0; i < steps; i++ {
			moveHead(instruction[0], &head)
			moveTail(&head, &tail)
			seen[tail] = true
		}
	}
	fmt.Println("Positions visited in first part", len(seen))
}

func partTwo() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	knots := [n]Position{}
	scanner := bufio.NewScanner(file)
	seen := make(map[Position]bool)
	seen[Position{0, 0}] = true
	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		steps, _ := strconv.Atoi(instruction[1])
		for i := 0; i < steps; i++ {
			moveHead(instruction[0], &knots[0])
			for j := 1; j < n; j++ {
				moveTail(&knots[j-1], &knots[j])
			}
			seen[knots[n-1]] = true
		}
	}
	fmt.Println("Positions visited in second part", len(seen))
}

func moveHead(direction string, head *Position) {
	if direction == "R" {
		head.x += 1
	}
	if direction == "L" {
		head.x -= 1
	}
	if direction == "U" {
		head.y += 1
	}
	if direction == "D" {
		head.y -= 1
	}
}

func moveTail(head *Position, tail *Position) {
	if math.Abs(float64(head.x-tail.x)) >= 2 || math.Abs(float64(head.y-tail.y)) >= 2 {
		if head.x != tail.x {
			if head.x > tail.x {
				tail.x += 1
			} else {
				tail.x -= 1
			}
		}
		if head.y != tail.y {
			if head.y > tail.y {
				tail.y += 1
			} else {
				tail.y -= 1
			}
		}
	}
}
