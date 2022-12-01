package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("day1\n answer1: %d\n answer2: %d", day1Task1(), day1Task2())
}

func day1Task2() int64 {
	input := readInput("assets/input1.txt")
	calories := []int64{}

	for _, lines := range strings.Split(input, "\n\n") {
		var sum int64
		for _, line := range strings.Split(lines, "\n") {
			if n, err := strconv.ParseInt(line, 10, 64); err == nil {
				sum += n
			}
		}
		calories = append(calories, sum)
	}

	nth := 3
	indeces := map[int]bool{}
	for k := 0; k < nth; k++ {
		index := 0
		var max int64
		for i := 0; i < len(calories); i++ {
			if _, ok := indeces[i]; ok {
				continue
			}
			if calories[i] > max {
				max = calories[i]
				index = i
			}
		}
		indeces[index] = true
	}

	var sum int64
	for i := range indeces {
		sum += calories[i]
	}

	return sum
}

func day1Task1() int64 {
	input := readInput("assets/input1.txt")
	calories := make([]int64, 0)

	for _, lines := range strings.Split(input, "\n\n") {
		var sum int64
		for _, line := range strings.Split(lines, "\n") {
			if n, err := strconv.ParseInt(line, 10, 64); err == nil {
				sum += n
			}
		}
		calories = append(calories, sum)
	}

	var max int64
	for i := 0; i < len(calories); i++ {
		if calories[i] > max {
			max = calories[i]
		}
	}

	return max
}

func readInput(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(bytes.NewBuffer(file).String())
}
