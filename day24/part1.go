package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type state struct {
	pos  image.Point
	step int
}

func main() {
	grid := map[image.Point]byte{}
	input := readAll("day24/input.txt")
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] != '.' {
				grid[image.Pt(x, y)] = line[x]
			}
		}
	}

	bounds := Bounds(getKeys(grid))
	entrance, exit := image.Pt(1, 0), image.Pt(bounds.Max.X-2, bounds.Max.Y-1)
	crossing := steps(grid, bounds, entrance, exit, 0)

	fmt.Println(crossing)
}

var Neighbors4 = []image.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func steps(grid map[image.Point]byte, bounds image.Rectangle, start image.Point, end image.Point, initialStep int) int {
	q := []state{{start, initialStep}}
	seen := map[state]struct{}{}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.pos == end {
			return curr.step
		}
	loop:
		for _, n := range append(Neighbors4, image.Pt(0, 0)) {
			nextstate := state{curr.pos.Add(n), curr.step + 1}
			if _, ok := seen[nextstate]; ok {
				continue
			}
			if !nextstate.pos.In(bounds) {
				continue
			}
			if grid[nextstate.pos] == '#' {
				continue
			}
			if nextstate.pos.Y > 0 && nextstate.pos.Y < bounds.Max.Y-1 {
				for _, bliz := range []byte{'^', '>', 'v', '<'} {
					prev := nextstate.pos.Sub(dirFromByte(bliz).PointR().Mul(nextstate.step)).Mod(bounds.Inset(1))
					if grid[prev] == bliz {
						continue loop
					}
				}
			}
			q = append(q, nextstate)
			seen[nextstate] = struct{}{}
		}
	}
	return -1
}

func readAll(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(file))
}

func getKeys(grid map[image.Point]byte) []image.Point {
	var keys []image.Point
	for k := range grid {
		keys = append(keys, k)
	}
	return keys
}

func Bounds(p []image.Point) image.Rectangle {
	var r image.Rectangle
	for _, pp := range p {
		r = r.Union(image.Rectangle{pp, pp.Add(image.Pt(1, 1))})
	}
	return r.Bounds()
}

type Dir int

const (
	N Dir = iota
	E
	S
	W
)

func (d Dir) Add(n int) Dir {
	return Dir((int(d) + n) % 4)
}

func (d Dir) Next() Dir {
	return (d + 1) % 4
}

func (d Dir) Prev() Dir {
	return (d + 3) % 4
}

func (d Dir) Reverse() Dir {
	return (d + 2) % 4
}

var point = map[Dir]image.Point{N: {0, 1}, E: {1, 0}, S: {0, -1}, W: {-1, 0}}
var pointReversed = map[Dir]image.Point{N: {0, -1}, E: {1, 0}, S: {0, 1}, W: {-1, 0}}

// Y-axis goes up.
func (d Dir) Point() image.Point {
	return point[d]
}

// Y-axis goes down.
func (d Dir) PointR() image.Point {
	return pointReversed[d]
}

var fromByte = map[byte]Dir{
	'N': N,
	'E': E,
	'S': S,
	'W': W,
	'U': N,
	'R': E,
	'D': S,
	'L': W,
	'^': N,
	'>': E,
	'v': S,
	'<': W,
}

func dirFromByte(b byte) Dir {
	return fromByte[b]
}
