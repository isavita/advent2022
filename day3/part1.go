package main

import (
	"bufio"
	"fmt"
	"os"
)

func itemPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	}
	return int(item-'A') + 27
}

func main() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		half := len(line) / 2
		firstCompartment := line[:half]
		secondCompartment := line[half:]

		compartmentMap := make(map[rune]int)
		for _, item := range firstCompartment {
			compartmentMap[item]++
		}
		for _, item := range secondCompartment {
			if _, exists := compartmentMap[item]; exists {
				sum += itemPriority(item)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(sum)
}
