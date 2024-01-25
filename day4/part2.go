package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day4/input.txt")
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")

		// Extract ranges
		left := parseRange(pair[0])
		right := parseRange(pair[1])

		// Check if ranges overlap
		if left[0] <= right[1] && left[1] >= right[0] {
			count++
		}
	}

	fmt.Println(count)
}

func parseRange(s string) []int {
	split := strings.Split(s, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])
	return []int{start, end}
}
