package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

const rockstr = `####

 # 
###
 # 

  #
  #
###

#
#
#
#

##
##`

func main() {
	jetPattern := []byte(readAll("day17/input.txt"))

	rocks := getRocks()
	grid := map[image.Point]struct{}{}
	for x := 0; x < 7; x++ {
		grid[image.Pt(x, 0)] = struct{}{}
	}
	floor, j := 0, 0
	repeat := map[[2]int][2]int{}

	for i, curr := 0, 0; ; i, curr = i+1, (curr+1)%len(rocks) {
		key := [2]int{curr, j}
		if r, ok := repeat[key]; ok {
			previ, prevFloor := r[0], r[1]
			if (1000000000000-i)%(i-previ) == 0 {
				fmt.Println(floor + (1000000000000-i)/(i-previ)*(floor-prevFloor))
				break
			}
		}
		repeat[key] = [2]int{i, floor}
		currRock := rocks[curr]
		pos := image.Pt(2, floor+4)
		for {
			jet := jetPattern[j]
			j = (j + 1) % len(jetPattern)
			pos = pos.Add(dirFromByte(jet).Point())
			if collision(grid, currRock, pos) {
				pos = pos.Sub(dirFromByte(jet).Point())
			}
			pos = pos.Add(S.Point())
			if collision(grid, currRock, pos) {
				pos = pos.Sub(S.Point())
				for p := range currRock {
					grid[p.Add(pos)] = struct{}{}
					if p.Add(pos).Y > floor {
						floor = p.Add(pos).Y
					}
				}
				break
			}
		}
	}
}

func collision(grid, rock map[image.Point]struct{}, pos image.Point) bool {
	for p := range rock {
		_, ok := grid[p.Add(pos)]
		if ok || p.Add(pos).X < 0 || p.Add(pos).X > 6 {
			return true
		}
	}
	return false
}

func getRocks() []map[image.Point]struct{} {
	rocks := []map[image.Point]struct{}{}
	for i, rock := range strings.Split(rockstr, "\n\n") {
		rocks = append(rocks, map[image.Point]struct{}{})
		lines := strings.Split(rock, "\n")
		for y, line := range lines {
			for x := 0; x < len(line); x++ {
				if line[x] == '#' {
					rocks[i][image.Pt(x, len(lines)-1-y)] = struct{}{}
				}
			}
		}
	}
	return rocks
}

func readAll(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(file))
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
