package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day21/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	jobs := make(map[string]string)
	results := make(map[string]int)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		jobs[parts[0]] = parts[1]
	}

	fmt.Println(calculate("root", jobs, results))
}

func calculate(monkey string, jobs map[string]string, results map[string]int) int {
	if val, ok := results[monkey]; ok {
		return val
	}

	job, ok := jobs[monkey]
	if !ok {
		panic("Monkey not found: " + monkey)
	}

	if num, err := strconv.Atoi(job); err == nil {
		results[monkey] = num
		return num
	}

	parts := strings.Fields(job)
	a := calculate(parts[0], jobs, results)
	b := calculate(parts[2], jobs, results)

	var result int
	switch parts[1] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("Unknown operation: " + parts[1])
	}

	results[monkey] = result
	return result
}
