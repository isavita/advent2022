package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	grid := map[image.Point]struct{}{}
	s := scanAll()
	for s.Scan() {
		sp := strings.Split(s.Text(), " -> ")
		pts := make([]image.Point, len(sp))
		for i := range sp {
			fmt.Sscanf(sp[i], "%d,%d", &pts[i].X, &pts[i].Y)
		}
		for i := 0; i < len(pts)-1; i++ {
			if pts[i].X == pts[i+1].X {
				for y := min(pts[i].Y, pts[i+1].Y); y <= max(pts[i].Y, pts[i+1].Y); y++ {
					grid[image.Pt(pts[i].X, y)] = struct{}{}
				}
			} else {
				for x := min(pts[i].X, pts[i+1].X); x <= max(pts[i].X, pts[i+1].X); x++ {
					grid[image.Pt(x, pts[i].Y)] = struct{}{}
				}
			}
		}
	}

	fmt.Println(fill(grid))
}

func scanAll() *bufio.Scanner {
	file, err := os.Open("day14/input.txt")
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func fill(grid map[image.Point]struct{}) int {
	floor := bounds(grid).Max.Y + 1
	sands, firstFloorTouch := 0, 0
	for full := false; !full; _, full = grid[image.Pt(500, 0)] {
		for sand, settled := image.Pt(500, 0), false; !settled; sand, settled = next(grid, sand) {
			if sand.Y == floor {
				if firstFloorTouch == 0 {
					firstFloorTouch = sands
				}
				grid[sand] = struct{}{}
				break
			}
		}
		sands++
	}
	return firstFloorTouch
}

var D = dirFromByte('D').PointR()
var L = dirFromByte('L').PointR()
var R = dirFromByte('R').PointR()

func next(grid map[image.Point]struct{}, sand image.Point) (image.Point, bool) {
	for _, n := range []image.Point{sand.Add(D), sand.Add(D).Add(L), sand.Add(D).Add(R)} {
		if _, ok := grid[n]; !ok {
			return n, false
		}
	}
	grid[sand] = struct{}{}
	return sand, true
}

func bounds(grid map[image.Point]struct{}) image.Rectangle {
	pts := make([]image.Point, 0, len(grid))
	for p := range grid {
		pts = append(pts, p)
	}
	return boundss(pts)
}

func boundss(p []image.Point) image.Rectangle {
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
