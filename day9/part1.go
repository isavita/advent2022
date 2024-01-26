package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct{ x, y int }

func main() {
	file, _ := os.Open("day9/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	head, tail := point{0, 0}, point{0, 0}
	visited := map[point]bool{tail: true}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir, steps := line[0], line[1]
		numSteps, _ := strconv.Atoi(steps)

		for i := 0; i < numSteps; i++ {
			switch dir {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}

			if abs(head.x-tail.x) > 1 || abs(head.y-tail.y) > 1 {
				if head.x != tail.x && head.y != tail.y {
					if head.x > tail.x {
						tail.x++
					} else {
						tail.x--
					}
					if head.y > tail.y {
						tail.y++
					} else {
						tail.y--
					}
				} else {
					if head.x > tail.x {
						tail.x++
					} else if head.x < tail.x {
						tail.x--
					}
					if head.y > tail.y {
						tail.y++
					} else if head.y < tail.y {
						tail.y--
					}
				}
			}

			visited[tail] = true
		}
	}

	fmt.Println(len(visited))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
