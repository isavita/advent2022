package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseRange(r string) (int, int) {
	parts := strings.Split(r, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return start, end
}

func main() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		if len(ranges) != 2 {
			continue
		}
		start1, end1 := parseRange(ranges[0])
		start2, end2 := parseRange(ranges[1])

		if (start1 <= start2 && end1 >= end2) || (start2 <= start1 && end2 >= end1) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(count)
}
