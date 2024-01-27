package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day14/input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(file))
	fmt.Println(solve(input))
}

func solve(input string) int {
	matrix := parseInput(input)
	originCol := 0
	for i, c := range matrix[0] {
		if c == "+" {
			originCol = i
		}
		matrix[len(matrix)-1][i] = "#"
	}

	ans := 0
	for !dropSand(matrix, originCol) {
		ans++
		// COULD incorporate this into the for loop conditional but then the
		// ordering is important... must check origin cell BEFORE running
		// dropSand... it's easier to read here...
		if matrix[0][originCol] == "o" {
			break
		}
	}

	return ans
}

func parseInput(input string) (matrix [][]string) {
	coordSets := [][][2]int{}
	lowestCol := math.MaxInt64
	highestRow := 0
	for _, line := range strings.Split(input, "\n") {
		rawCoords := strings.Split(line, " -> ")
		coords := [][2]int{}
		for _, rawCoord := range rawCoords {
			rawNums := strings.Split(rawCoord, ",")
			col, row := toInt(rawNums[0]), toInt(rawNums[1])
			coord := [2]int{
				col, row,
			}
			coords = append(coords, coord)

			lowestCol = minInt(lowestCol, col)
			highestRow = maxInt(highestRow, row)
		}
		coordSets = append(coordSets, coords)
	}

	// lowering this number to 1 makes it easier to print the matrix, which I
	// used for part 1... but then needed to up it for part 2... or just have a
	// massive screen and make the terminal text tiny...
	ExtraLeftSpace := 200

	highestCol := 0
	for s, set := range coordSets {
		for i := range set {
			coordSets[s][i][0] -= lowestCol - ExtraLeftSpace
			highestCol = maxInt(highestCol, coordSets[s][i][0])
		}
	}

	matrix = make([][]string, highestRow+3)
	for r := range matrix {
		matrix[r] = make([]string, highestCol+ExtraLeftSpace*2)
	}

	for _, set := range coordSets {
		for i := 1; i < len(set); i++ {
			cols := []int{set[i-1][0], set[i][0]}
			rows := []int{set[i-1][1], set[i][1]}

			sort.Ints(cols)
			sort.Ints(rows)

			if cols[0] == cols[1] {
				for r := rows[0]; r <= rows[1]; r++ {
					matrix[r][cols[0]] = "#"
				}
			} else if rows[0] == rows[1] {
				for c := cols[0]; c <= cols[1]; c++ {
					matrix[rows[0]][c] = "#"
				}
			}
		}
	}

	originCol := 500 - lowestCol + ExtraLeftSpace
	// make it a plus so it's searchable in the next step... or could just
	// return this value too...
	matrix[0][originCol] = "+"

	for i, r := range matrix {
		for j := range r {
			if matrix[i][j] == "" {
				matrix[i][j] = "."
			}
		}
	}

	// printMatrix(matrix)
	return matrix
}

func printMatrix(matrix [][]string) {
	for _, r := range matrix {
		fmt.Println(r)
	}
}

func dropSand(matrix [][]string, originCol int) (fallsIntoAbyss bool) {
	r, c := 0, originCol

	for r < len(matrix)-1 {
		below := matrix[r+1][c]
		diagonallyLeft := matrix[r+1][c-1]
		diagonallyRight := matrix[r+1][c+1]
		if below == "." {
			r++
		} else if diagonallyLeft == "." {
			r++
			c--
		} else if diagonallyRight == "." {
			r++
			c++
		} else {
			matrix[r][c] = "o"
			return false
		}
	}

	return true
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
