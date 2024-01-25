package main

import (
	"bufio"
	"fmt"
	"os"
)

func itemPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item-'a') + 1
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

	var sum, groupLineCounter int
	groupItems := make([]map[rune]int, 3)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		itemsMap := make(map[rune]int)
		for _, item := range line {
			itemsMap[item]++
		}
		groupItems[groupLineCounter] = itemsMap
		groupLineCounter++

		if groupLineCounter == 3 {
			commonItems := make(map[rune]int)
			for item := range groupItems[0] {
				if groupItems[1][item] > 0 && groupItems[2][item] > 0 {
					commonItems[item]++
				}
			}
			for item := range commonItems {
				sum += itemPriority(item)
				break // Since we need only one common item per group
			}
			groupLineCounter = 0
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(sum)
}
