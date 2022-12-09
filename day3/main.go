package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	partOne()
	data := parseInput()
	partTwo(data)
}

func partOne() {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	var totalScore int32
	for scanner.Scan() {
		line := scanner.Text()
		common := make(map[rune]bool)
		left := line[:len(line)/2]

		for _, val := range left {
			common[val] = true
		}

		right := line[len(line)/2:]
		for _, val := range right {
			if common[val] {
				totalScore += getPriority(val)
				common[val] = false // as we do not want to count multiple instances of a letter
			}
		}
	}

	fmt.Println("The total priority is", totalScore)
}

func partTwo(data []string) {
	var totalScore int32

	for i := 0; i < len(data); i += 3 {
		common := make(map[rune]int)
		ruck0 := data[i]
		ruck1 := data[i+1]
		ruck2 := data[i+2]

		for _, val := range ruck0 {
			common[val] = 1
		}
		for _, val := range ruck1 {
			if _, found := common[val]; found {
				common[val] = 2 // First I was going to increment it but there could be duplicates in the strings and that could give false values
			}
		}

		for _, val := range ruck2 {
			if count, found := common[val]; found && count == 2 {
				totalScore += getPriority(val)
				// as every item only appears once
				break
			}
		}
	}

	fmt.Println("Total score of second part is ", totalScore)
}

func parseInput() []string {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

func getPriority(val int32) int32 {
	if val >= 65 && val <= 90 {
		return val - 'A' + 27
	} else {
		return val - 'a' + 1
	}
}
