package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var caloriesList []int
	currentCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			caloriesList = append(caloriesList, currentCalories)
			currentCalories = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Error converting line to int: %v", err)
		}

		currentCalories += calories
	}

	caloriesList = append(caloriesList, currentCalories)
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesList)))

	topThreeSum := 0
	for i := 0; i < 3 && i < len(caloriesList); i++ {
		topThreeSum += caloriesList[i]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(topThreeSum)
}
