package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type sensor struct {
	pos    image.Point
	beacon image.Point
	dist   int
}

func main() {
	var sensors []sensor
	input := readAll("day15/input.txt")
	for _, line := range strings.Split(input, "\n") {
		var s sensor
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.pos.X, &s.pos.Y, &s.beacon.X, &s.beacon.Y)
		s.dist = manhattan(s.pos, s.beacon)
		sensors = append(sensors, s)
	}
	fmt.Println(distress(sensors, 4000000))
}

func distress(sensors []sensor, maxcoord int) int {
	for x := 0; x <= maxcoord; x++ {
		for y := 0; y <= maxcoord; y++ {
			p := image.Pt(x, y)
			detected := false
			skip := 0
			for _, s := range sensors {
				if manhattan(s.pos, p) <= s.dist {
					detected = true
					dist := s.dist - abs(s.pos.X-x)
					skip = max(skip, dist+s.pos.Y-y)
				}
			}
			if !detected {
				return x*4000000 + y
			}
			y += skip
		}
	}
	return -1
}

func readAll(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func manhattan(p, q image.Point) int {
	return abs(p.X-q.X) + abs(p.Y-q.Y)
}
