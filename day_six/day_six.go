package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input")
	stream := strings.Split(string(bytes), "")

	ans := findUnique(stream, 4)
	fmt.Println("The answer of part One is", ans)

	ans = findUnique(stream, 14)
	fmt.Println("The answer of part Two is", ans)
}

func findUnique(stream []string, marker int) int {
	for i := 0; i < len(stream)-marker; i++ {
		buffer := make(map[string]bool)
		for j := i; j < i+marker; j++ {
			if _, ok := buffer[stream[j]]; ok {
				break
			}
			buffer[stream[j]] = true
			if len(buffer) == marker {
				return i + marker
			}
		}
	}
	return -1
}
