package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SensorBeacon struct {
	SensorX, SensorY, BeaconX, BeaconY int
}

func parseInput(filename string) ([]SensorBeacon, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sensorsBeacons []SensorBeacon
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		fmt.Println(parts, len(parts))
		sensorX, _ := strconv.Atoi(strings.Split(parts[2], "=")[1])
		sensorY, _ := strconv.Atoi(strings.Split(parts[3], "=")[1])
		beaconX, _ := strconv.Atoi(strings.Split(parts[8], "=")[1])
		beaconY, _ := strconv.Atoi(strings.Split(parts[9], "=")[1][:len(parts[9])-1])

		sensorsBeacons = append(sensorsBeacons, SensorBeacon{sensorX, sensorY, beaconX, beaconY})
	}

	return sensorsBeacons, scanner.Err()
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(sensorsBeacons []SensorBeacon, row int) int {
	var totalCovered int
	for _, sb := range sensorsBeacons {
		distance := manhattanDistance(sb.SensorX, sb.SensorY, sb.BeaconX, sb.BeaconY)
		remainingDistance := distance - abs(sb.SensorY-row)
		if remainingDistance >= 0 {
			totalCovered += 2*remainingDistance + 1
		}
	}

	beaconsOnRow := 0
	for _, sb := range sensorsBeacons {
		if sb.SensorY == row && sb.BeaconX >= 0 {
			beaconsOnRow++
		}
	}

	return totalCovered - beaconsOnRow
}

func main() {
	filename := "day15/input.txt"
	sensorsBeacons, err := parseInput(filename)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	row := 2000000
	result := part1(sensorsBeacons, row)
	fmt.Printf("Part 1: %d\n", result)
}
