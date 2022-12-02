package main

import (
	"bufio"
	"fmt"
	"math"
	"sort"

	//"sort"
	"log"
	"os"
	"strconv"
)

func main() {
	input := parseInput()
	partOne(input)
	partTwo(input)
	partTwoOptimized(input)

}

func partOne(data []string) {
	maximumCalorie := math.MinInt
	elvesCalorie := 0

	for _, calorie := range data {
		if calorie == "." {
			maximumCalorie = max(maximumCalorie, elvesCalorie)
			elvesCalorie = 0
		} else {
			currentCalorie, err := strconv.Atoi(calorie)
			if err != nil {
				log.Fatal(err)
			}
			elvesCalorie += currentCalorie
		}
	}

	maximumCalorie = max(maximumCalorie, elvesCalorie)
	fmt.Println("Maximum elves calorie is ", maximumCalorie)
}

func partTwo(data []string) {
	var CalorieList []int
	var currentElf int
	for _, calorie := range data {
		if calorie == "." {
			CalorieList = append(CalorieList, currentElf)
			currentElf = 0
		} else {
			currentCalorie, err := strconv.Atoi(calorie)
			if err != nil {
				log.Fatal(err)
			}
			currentElf += currentCalorie
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(CalorieList)))

	sumOfThree := 0
	for i := 0; i < 3; i++ {
		sumOfThree += CalorieList[i]
	}

	fmt.Println("Total of the top 3 elves is : ", sumOfThree)
}

func partTwoOptimized(data []string) {
	topThree := [3]int{0, 0, 0}

	var currentElf int
	for _, calorie := range data {
		if calorie == "." {
			maxThree(&topThree, currentElf)
			currentElf = 0
		} else {
			currentCalorie, err := strconv.Atoi(calorie)
			if err != nil {
				log.Fatal(err)
			}
			currentElf += currentCalorie
		}
	}

	sumOfThree := 0
	for i := 0; i < 3; i++ {
		sumOfThree += topThree[i]
	}

	fmt.Println("Total of the top 3 elves is : ", sumOfThree)
}

func parseInput() []string {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			data = append(data, ".")
		} else {
			data = append(data, line)
		}
	}

	return data
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxThree(topThree *[3]int, curr int) {
	for i, val := range topThree {
		if curr >= val {
			topThree[i], curr = curr, val
		}
	}
}
