package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

const (
	noOP = "noop"
	addx = "addx"
)

var OpCycle = map[string]int{
	noOP: 1,
	addx: 2,
}

func partOne() {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	cycles := 0
	register := 1
	sumOfSignals := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		iterations := OpCycle[line[0]]
		for i := 0; i < iterations; i++ {
			cycles++
			sumOfSignals += isImportantCycle(cycles, register)
		}
		if line[0] == addx {
			val, _ := strconv.Atoi(line[1])
			register += val
		}
	}
	fmt.Println("Sum of signals is ", sumOfSignals)
}

func partTwo() {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	spritePosition := 1
	scanner := bufio.NewScanner(f)
	cycles := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		iterations := OpCycle[line[0]]
		for i := 0; i < iterations; i++ {
			cycles++
			printPixel(cycles, spritePosition)
		}
		if line[0] == addx {
			val, _ := strconv.Atoi(line[1])
			spritePosition += val
		}
	}
}

func isImportantCycle(cycles, register int) int {
	if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
		fmt.Println(cycles, register)
		return cycles * register
	}
	return 0
}

func printPixel(cycles, spritePosition int) {
	if (cycles-1)%40 == 0 {
		fmt.Println()
	}
	if (cycles-1)%40 >= spritePosition-1 && (cycles-1)%40 <= spritePosition+1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
}
