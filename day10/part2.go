package main

import (
	"fmt"
	"image"
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

	grid := map[image.Point]struct{}{}
	for i := range x {
		crtx, crty := i%40, i/40
		if abs(crtx-x[i]) <= 1 {
			grid[image.Pt(crtx, crty)] = struct{}{}
		} else {
			delete(grid, image.Pt(crtx, crty))
		}
	}

	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			if _, ok := grid[image.Pt(x, y)]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
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
