package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	x := []int{1}
	for _, line := range strings.Split(readAll("day10/input.txt"), "\n") {
		switch line {
		case "noop":
			x = append(x, x[len(x)-1])
		default:
			var n int
			fmt.Sscanf(line, "addx %d", &n)
			x = append(x, x[len(x)-1])
			x = append(x, x[len(x)-1]+n)
		}
	}

	sum := 0
	for i := range x {
		if (i-19)%40 == 0 {
			sum += (i + 1) * x[i]
		}
	}
	fmt.Println(sum)
}

func readAll(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
