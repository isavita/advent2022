package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	cubes := map[Pt3]struct{}{}
	neighbors := []Pt3{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}
	min := Pt3{math.MaxInt, math.MaxInt, math.MaxInt}
	max := Pt3{math.MinInt, math.MinInt, math.MinInt}

	s := scanAll()
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		var cube Pt3
		fmt.Sscanf(line, "%d,%d,%d", &cube.X, &cube.Y, &cube.Z)
		cubes[cube] = struct{}{}
		min = Pt3{Min(min.X, cube.X), Min(min.Y, cube.Y), Min(min.Z, cube.Z)}
		max = Pt3{Max(max.X, cube.X), Max(max.Y, cube.Y), Max(max.Z, cube.Z)}
	}
	min = min.Add(Pt3{-1, -1, -1})
	max = max.Add(Pt3{1, 1, 1})

	faces := 0
	q := []Pt3{min}
	seen := map[Pt3]struct{}{min: {}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, delta := range neighbors {
			next := curr.Add(delta)
			if next.X < min.X ||
				next.Y < min.Y ||
				next.Z < min.Z ||
				next.X > max.X ||
				next.Y > max.Y ||
				next.Z > max.Z {
				continue
			}
			if _, ok := cubes[next]; ok {
				faces++
			} else if _, ok := seen[next]; !ok {
				seen[next] = struct{}{}
				q = append(q, next)
			}
		}
	}
	fmt.Println(faces)
}

type Pt3 struct {
	X, Y, Z int
}

func (p1 Pt3) Add(p2 Pt3) Pt3 {
	return Pt3{p1.X + p2.X, p1.Y + p2.Y, p1.Z + p2.Z}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanAll() *bufio.Scanner {
	file, err := os.Open("day18/input.txt")
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}
