package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	maxSize        = 100000
	totalDiskSpace = 70000000
	spaceNeeded    = 30000000
)

func main() {
	fileSizes := parseCommands()
	directorySizes := getDirectorySizes(fileSizes)
	totalSize := partOne(directorySizes)
	fmt.Println("The output of part one is ", totalSize)
	directoryToDelete := partTwo(directorySizes)
	fmt.Println("THe output of part two is ", directoryToDelete)
}

func parseCommands() map[string]int {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	dir := ""
	fileSize := make(map[string]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")
		if args[0] == "$" {
			if args[1] == "cd" {
				dir = path.Join(dir, args[2])
			}
		} else if args[0] != "dir" {
			size, _ := strconv.Atoi(args[0])
			fileSize[path.Join(dir, args[1])] = size
		}
	}
	return fileSize
}
func getDirectorySizes(fileSizes map[string]int) map[string]int {
	directorySizes := map[string]int{}

	for file, size := range fileSizes {
		for dir := path.Dir(file); dir != "/"; dir = path.Dir(dir) {
			directorySizes[dir] += size
		}
		directorySizes["/"] += size
	}
	return directorySizes
}

func partOne(directorySizes map[string]int) int {
	totalSum := 0
	for _, size := range directorySizes {
		if size <= maxSize {
			totalSum += size
		}
	}
	return totalSum
}

func partTwo(directorySizes map[string]int) int {
	toDelete := spaceNeeded - (totalDiskSpace - directorySizes["/"])

	dirToDelete := directorySizes["/"]
	for _, size := range directorySizes {
		if size >= toDelete && size < dirToDelete {
			dirToDelete = size
		}
	}
	return dirToDelete
}
