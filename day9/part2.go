package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input := readAll()
	fmt.Println(visited(input, 10))
}

func visited(input string, ropelen int) int {
	rope := make([]image.Point, ropelen)
	visited := map[image.Point]struct{}{}
	for _, line := range strings.Split(input, "\n") {
		var b byte
		var n int
		fmt.Sscanf(line, "%c %d", &b, &n)
		d := dirFromByte(b)
		for i := 0; i < n; i++ {
			rope[0] = rope[0].Add(d.Point())
			for j := 1; j < ropelen; j++ {
				rope[j] = next(rope[j-1], rope[j])
			}
			visited[rope[ropelen-1]] = struct{}{}
		}
	}
	return len(visited)
}

func next(head, tail image.Point) image.Point {
	if abs(head.X-tail.X) <= 1 && abs(head.Y-tail.Y) <= 1 {
		return tail
	}
	return tail.Add(image.Pt(sign(head.X-tail.X), sign(head.Y-tail.Y)))
}

func readAll() string {
	file, err := os.ReadFile("day9/input.txt")
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

func sign(n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	return 1
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

var fromPoint = map[image.Point]Dir{{0, 1}: N, {1, 0}: E, {0, -1}: S, {-1, 0}: W}
var fromPointReversed = map[image.Point]Dir{{0, -1}: N, {1, 0}: E, {0, 1}: S, {-1, 0}: W}

// Y-axis goes up.
func DirFromPoint(p image.Point) Dir {
	return fromPoint[p]
}

// Y-axis goes down.
func DirFromPointR(p image.Point) Dir {
	return fromPointReversed[p]
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
