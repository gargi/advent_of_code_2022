package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items     []int
	operation func(int) int
	test      int
	ifTrue    int
	ifFalse   int
}

var worryFactor = 1

func main() {
	input := parseInput()
	input2 := append([]monkey(nil), input...)
	partOne(input)
	partTwo(input2)
}

func parseInput() []monkey {
	var monkeys []monkey
	bytes, _ := os.ReadFile("input")
	info := strings.Split(string(bytes), "\n\n")
	for _, attributes := range info {
		var items []int
		line := strings.Split(attributes, "\n")
		startingItems := strings.TrimPrefix(strings.TrimSpace(line[1]), "Starting items: ")
		for _, item := range strings.Split(startingItems, ", ") {
			intVal, _ := strconv.Atoi(item)
			items = append(items, intVal)
		}
		operations := strings.Split(strings.TrimSpace(line[2]), " ")
		op := operations[4]
		operand := operations[5]
		operation := func(old int) int {
			switch op {
			case "+":
				if operand == "old" {
					return old + old
				}
				val, _ := strconv.Atoi(operand)
				return old + val
			case "*":
				if operand == "old" {
					return old * old
				}
				val, _ := strconv.Atoi(operand)
				return old * val
			default:
				panic(op)
			}
		}
		test := strings.Split(strings.TrimSpace(line[3]), " ")
		divisibleBy, _ := strconv.Atoi(test[len(test)-1])
		worryFactor *= divisibleBy
		ifTrue := strings.Split(strings.TrimSpace(line[4]), " ")
		truemonkey, _ := strconv.Atoi(ifTrue[len(ifTrue)-1])
		ifFalse := strings.Split(strings.TrimSpace(line[5]), " ")
		falseMonkey, _ := strconv.Atoi(ifFalse[len(ifFalse)-1])
		m := monkey{
			items:     items,
			operation: operation,
			test:      divisibleBy,
			ifTrue:    truemonkey,
			ifFalse:   falseMonkey,
		}
		monkeys = append(monkeys, m)
	}
	return monkeys
}

func partOne(monkeys []monkey) {
	counts := make([]int, len(monkeys))
	for round := 0; round < 20; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				item := m.operation(item)
				item = item / 3
				if item%m.test == 0 {
					monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, item)
				} else {
					monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, item)
				}
				counts[i] += 1
			}
			monkeys[i].items = nil
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	fmt.Println("Output of part 1 is ", counts[0]*counts[1])
}

func partTwo(monkeys []monkey) {
	counts := make([]int, len(monkeys))
	for round := 0; round < 10000; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				item := m.operation(item)
				item = item % worryFactor
				if item%m.test == 0 {
					monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, item)
				} else {
					monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, item)
				}
				counts[i] += 1
			}
			monkeys[i].items = nil
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	fmt.Println("Output of part 2 is ", counts[0]*counts[1])
}
