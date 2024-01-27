package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

func main() {
	file, err := os.Open("day18/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cubes := make(map[Point]bool)

	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		cubes[Point{x, y, z}] = true
	}

	surfaceArea := 0
	for cube := range cubes {
		surfaceArea += calculateExposedSides(cube, cubes)
	}

	fmt.Println(surfaceArea)
}

func calculateExposedSides(p Point, cubes map[Point]bool) int {
	directions := []Point{
		{1, 0, 0}, {-1, 0, 0}, // x directions
		{0, 1, 0}, {0, -1, 0}, // y directions
		{0, 0, 1}, {0, 0, -1}, // z directions
	}

	exposedSides := 6
	for _, dir := range directions {
		adjacent := Point{p.x + dir.x, p.y + dir.y, p.z + dir.z}
		if cubes[adjacent] {
			exposedSides--
		}
	}
	return exposedSides
}
