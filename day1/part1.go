package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxCalories := 0
	currentCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Error converting line to int: %v", err)
		}

		currentCalories += calories
	}

	if currentCalories > maxCalories {
		maxCalories = currentCalories
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(maxCalories)
}
