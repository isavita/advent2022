package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		opponent, yourMove := line[0], line[2]

		score := 0
		if yourMove == 'X' {
			score = 1
		} else if yourMove == 'Y' {
			score = 2
		} else if yourMove == 'Z' {
			score = 3
		}

		if (opponent == 'A' && yourMove == 'Y') || (opponent == 'B' && yourMove == 'Z') || (opponent == 'C' && yourMove == 'X') {
			score += 6
		} else if opponent == 'A' && yourMove == 'X' || opponent == 'B' && yourMove == 'Y' || opponent == 'C' && yourMove == 'Z' {
			score += 3
		}

		totalScore += score
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	fmt.Println(totalScore)
}
