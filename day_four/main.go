package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	lower int
	upper int
}

func main() {
	input := parseInput()
	partOne(input)
	partTwo(input)
}

func parseInput() [][]Elf {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	var elves [][]Elf
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Text()
		intervals := strings.Split(data, ",")
		elfOne := strings.Split(intervals[0], "-")
		elfTwo := strings.Split(intervals[1], "-")
		elfOneLower, _ := strconv.Atoi(elfOne[0])
		elfOneUpper, _ := strconv.Atoi(elfOne[1])
		elfTwoLower, _ := strconv.Atoi(elfTwo[0])
		elfTwoUpper, _ := strconv.Atoi(elfTwo[1])
		elfPair := []Elf{{lower: elfOneLower, upper: elfOneUpper}, {lower: elfTwoLower, upper: elfTwoUpper}}
		elves = append(elves, elfPair)
	}
	return elves
}

func partOne(elves [][]Elf) {
	overlaps := 0
	for _, elfPair := range elves {
		if isTotalOverlap(elfPair[0], elfPair[1]) {
			overlaps += 1
		}
	}
	fmt.Println(overlaps)
}

func partTwo(elves [][]Elf) {
	overlaps := 0
	for _, elfPair := range elves {
		if isPartialOverlap(elfPair[0], elfPair[1]) {
			overlaps += 1
		}
	}
	fmt.Println(overlaps)
}

func isTotalOverlap(elfOne Elf, elfTwo Elf) bool {
	if elfOne.lower <= elfTwo.lower && elfOne.upper >= elfTwo.upper {
		return true
	} else if elfTwo.lower <= elfOne.lower && elfTwo.upper >= elfOne.upper {
		return true
	}
	return false
}

func isPartialOverlap(elfOne Elf, elfTwo Elf) bool {
	if elfOne.lower <= elfTwo.lower && elfTwo.lower <= elfOne.upper {
		return true
	} else if elfTwo.lower <= elfOne.lower && elfOne.lower <= elfTwo.upper {
		return true
	}
	return false
}
