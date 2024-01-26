package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

var Neighbors4 = []image.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func main() {
	grid := map[image.Point]int{}
	s := scanAll()
	y := 0
	for s.Scan() {
		for x, b := range s.Bytes() {
			grid[image.Pt(x, y)] = int(b - '0')
		}
		y++
	}

	maxScore := 0
	for p := range grid {
		score := 1
		for _, n := range Neighbors4 {
			next, view := p, 0
			for {
				next = next.Add(n)
				if _, ok := grid[next]; ok {
					view++
					if grid[next] >= grid[p] {
						score *= view
						break
					}
				} else {
					score *= view
					break
				}
			}
		}

		if score > maxScore {
			maxScore = score
		}
	}
	fmt.Println(maxScore)
}

func scanAll() *bufio.Scanner {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}
