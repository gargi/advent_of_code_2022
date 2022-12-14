package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

var dividers = [][]any{{[]any{float64(2)}}, {[]any{float64(6)}}}

func main() {
	packetPairs := parseInput()

	packets := partOne(packetPairs)

	partTwo(packets)
}

type bySize [][]any

func (p bySize) Len() int {
	return len(p)
}

func (p bySize) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p bySize) Less(i, j int) bool {
	return Compare(p[i], p[j], 0) <= 0
}

func parseInput() [][2][]any {
	bytes, _ := os.ReadFile("input")
	signals := strings.Split(string(bytes), "\n\n")
	var packetPairs [][2][]any

	for _, signal := range signals {
		packets := strings.Split(signal, "\n")
		leftPacket, rightPacket := parsePacket(packets[0]), parsePacket(packets[1])
		pair := [2][]any{leftPacket, rightPacket}
		packetPairs = append(packetPairs, pair)
	}

	return packetPairs
}

func partOne(packetPairs [][2][]any) [][]any {
	rightOrderSum := 0
	var packets [][]any

	for i, pairs := range packetPairs {
		packets = append(packets, pairs[0], pairs[1])
		leftPacket, rightPacket := pairs[0], pairs[1]
		if Compare(leftPacket, rightPacket, 0) <= 0 {
			rightOrderSum += i + 1
		}
	}

	fmt.Println("Output of part 1 is : ", rightOrderSum)

	return packets
}
func partTwo(packets [][]any) {
	packets = append(packets, dividers...)

	sort.Sort(bySize(packets))

	var key = 1
	for i, p := range packets {
		if reflect.DeepEqual(p, dividers[0]) || reflect.DeepEqual(p, dividers[1]) {
			key *= i + 1
		}
	}

	fmt.Println("Output of part 2 is : ", key)
}

func parsePacket(packet string) []any {
	var list []interface{}
	err := json.Unmarshal([]byte(packet), &list)
	if err != nil {
		panic(err)
	}
	return list
}

func Compare(left, right []any, index int) int {
	if index >= len(left) && index >= len(right) {
		return 0 // end of both lists
	}
	if index >= len(left) {
		return -1 // left ended first
	}

	if index >= len(right) {
		return 1 // right ended first
	}

	// Check if the first element is a float
	li := left[index]
	ri := right[index]
	leftElement, islFloat := li.(float64)
	rightElement, isRFloat := ri.(float64)

	if islFloat && isRFloat {
		if leftElement == rightElement {
			return Compare(left, right, index+1)
		} else {
			return int(leftElement - rightElement)
		}
	}

	// Check if list
	llist, isllist := li.([]any)
	rlist, isrlist := ri.([]any)

	if !isllist {
		llist = []any{li} // mixed, convert to list
	}
	if !isrlist {
		rlist = []any{ri} // mixed, convert to list
	}

	res := Compare(llist, rlist, 0)

	if res == 0 { // tie, continue to next item
		return Compare(left, right, index+1)
	}
	return res
}
