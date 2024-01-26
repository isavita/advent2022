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
	visible := map[image.Point]struct{}{}
	s := scanAll()
	y := 0
	for s.Scan() {
		for x, b := range s.Bytes() {
			grid[image.Pt(x, y)] = int(b - '0')
		}
		y++
	}

	for p := range grid {
		for _, n := range Neighbors4 {
			next := p
			for {
				next = next.Add(n)
				if _, ok := grid[next]; ok {
					if grid[next] >= grid[p] {
						break
					}
				} else {
					visible[p] = struct{}{}
					break
				}
			}
		}
	}
	fmt.Println(len(visible))
}

func scanAll() *bufio.Scanner {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}
