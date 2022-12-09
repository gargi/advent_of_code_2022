package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){
	partOne()
	partTwo()
}

func partOne() {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		current := strings.Split(line, " ")
		if current[1] == "X"{ // rock
			totalScore += 1
			if current[0] == "A" { // rock
			totalScore += 3
			}
			if current[0] == "C" { //scissor
				totalScore += 6
			}
		}
		if current[1] == "Y" { // paper
			totalScore += 2
			if current[0] == "A" {
			totalScore += 6
			}
			if current[0] == "B" { // paper
			 totalScore += 3
			}
		}
		if current[1] == "Z" { // scissor
			totalScore += 3
			if current[0] == "B" { // paper
				totalScore += 6
			}
			if current[0] == "C" {
				totalScore += 3
			}
		}
	}

	fmt.Println("Total score of part 1 :", totalScore)
}

func partTwo() {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		current := strings.Split(line, " ")
		if current[1] == "X"{ // lose
			if current[0] == "A" { // rock -> select scissors
				totalScore += 3
			}
			if current[0] == "B" { // paper -> select rock
				totalScore += 1
			}
			if current[0] == "C" { //scissor -> select paper
				totalScore += 2
			}
		}
		if current[1] == "Y" { // draw
			totalScore += 3
			if current[0] == "A" { // rock -> select rock
				totalScore += 1
			}
			if current[0] == "B" { // paper -> select paper
				totalScore += 2
			}
			if current[0] == "C" { // scissor -> select scissor
				totalScore += 3
			}
		}
		if current[1] == "Z" { // win
			totalScore += 6
			if current[0] == "A" { // rock -> select paper
				totalScore += 2
			}
			if current[0] == "B" { // paper -> select scissor
				totalScore += 3
			}
			if current[0] == "C" { // scissor -> select rock
				totalScore += 1
			}
		}
	}

	fmt.Println("Total score of part 2 :", totalScore)
}
