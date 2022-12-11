package main

import (
	"bytes"
	"container/list"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("day1\n answer1: %d\n answer2: %d", day1Task1(), day1Task2())
	fmt.Printf("\nday2\n answer1: %d\n answer2: %d", day2Task1(), day2Task2())
	fmt.Printf("\nday3\n answer1: %d\n answer2: %d", day3Task1(), day3Task2())
	fmt.Printf("\nday4\n answer1: %d\n answer2: %d", day4Task1(), day4Task2())
	fmt.Printf("\nday5\n answer1: %s\n answer2: %s", day5Task1(), day5Task2())
	fmt.Printf("\nday6\n answer1: %d\n answer2: %d", day6Task1(), day6Task2())
	fmt.Printf("\nday7\n answer1: %d\n answer2: %d", day7Task1(), day7Task2())
	fmt.Printf("\nday8\n answer1: %d\n answer2: %d", day8Task1(), day8Task2())
	fmt.Printf("\nday9\n answer1: %d\n answer2: %d", day9Task1(), day9Task2())
}

func day9Task2() int {
	input := readInput("assets/input9.txt")

	positions := make(map[Position]int)
	snake := list.New()
	snake.PushBack(Position{x: 5, y: 5, segmentType: 1})
	snake.PushBack(Position{x: 5, y: 5, segmentType: 2})
	snake.PushBack(Position{x: 5, y: 5, segmentType: 3})
	snake.PushBack(Position{x: 5, y: 5, segmentType: 4})
	snake.PushBack(Position{x: 5, y: 5, segmentType: 5})
	snake.PushBack(Position{x: 5, y: 5, segmentType: 6})
	snake.PushBack(Position{x: 5, y: 5, segmentType: 7})
	snake.PushBack(Position{x: 5, y: 5, segmentType: 8})
	snake.PushBack(Position{x: 5, y: 5, segmentType: 9})
	snake.PushFront(Position{x: 5, y: 5, segmentType: Head})
	positions[Position{x: 5, y: 5, segmentType: 9}] = 1

	for _, line := range strings.Split(input, "\n") {
		direction := day9GetDirection(line)
		for ; direction.steps > 0; direction.steps -= 1 {
			headSegment := snake.Front()
			snake.Remove(headSegment)
			headPos := headSegment.Value.(Position)
			newX := headPos.x + direction.x
			newY := headPos.y + direction.y
			newTail := day9UpdateTail(newX, newY, snake, &positions)
			snake = list.New()
			snake.PushFrontList(&newTail)
			snake.PushFront(Position{x: newX, y: newY, segmentType: Head})
		}
	}

	return len(positions)
}

func day9UpdateTail(newX, newY int, snake *list.List, positions *map[Position]int) list.List {
	newTail := list.New()
	for item := snake.Front(); item != nil; item = item.Next() {
		tailPos := Position{item.Value.(Position).x, item.Value.(Position).y, item.Value.(Position).segmentType}
		if day9IsValidPosition(newX, newY, &tailPos) {
			newX, newY = tailPos.x, tailPos.y
			newTail.PushBack(tailPos)
			if tailPos.segmentType == 9 {
				(*positions)[Position{x: tailPos.x, y: tailPos.y, segmentType: tailPos.segmentType}] += 1
			}
		} else {
			deltX, deltY := day9CalcTailUpdate(newX, newY, &tailPos)
			newXTail := tailPos.x + deltX
			newYTail := tailPos.y + deltY
			newX, newY = newXTail, newYTail
			newTail.PushBack(Position{x: newXTail, y: newYTail, segmentType: tailPos.segmentType})
			if tailPos.segmentType == 9 {
				(*positions)[Position{x: newXTail, y: newYTail, segmentType: tailPos.segmentType}] += 1
			}
		}
	}
	return *newTail
}

func day9Task1() int {
	input := readInput("assets/input9.txt")

	positions := make(map[Position]int)
	snake := list.New()
	snake.PushFront(Position{x: 5, y: 5, segmentType: Tail})
	snake.PushFront(Position{x: 5, y: 5, segmentType: Head})
	positions[Position{x: 5, y: 5, segmentType: Tail}] = 1

	for _, line := range strings.Split(input, "\n") {
		direction := day9GetDirection(line)
		for ; direction.steps > 0; direction.steps -= 1 {
			headSegment := snake.Front()
			snake.Remove(headSegment)
			headPos := headSegment.Value.(Position)
			newX := headPos.x + direction.x
			newY := headPos.y + direction.y
			snake.PushFront(Position{x: newX, y: newY, segmentType: Head})
			tailSegment := snake.Back()
			snake.Remove(tailSegment)
			tailPos := tailSegment.Value.(Position)
			if day9IsValidPosition(newX, newY, &tailPos) {
				snake.PushBack(tailPos)
				positions[Position{x: tailPos.x, y: tailPos.y, segmentType: Tail}] += 1
			} else {
				deltX, deltY := day9CalcTailUpdate(newX, newY, &tailPos)
				newXTail := tailPos.x + deltX
				newYTail := tailPos.y + deltY
				snake.PushBack(Position{x: newXTail, y: newYTail, segmentType: Tail})
				positions[Position{x: newXTail, y: newYTail, segmentType: Tail}] += 1
			}
		}
	}

	return len(positions)
}

type Position struct {
	x           int
	y           int
	segmentType int // Head or Tail
}

const (
	Head = iota
	Tail
)

type Direction struct {
	x     int
	y     int
	steps int
}

func day9CalcTailUpdate(xHead, yHead int, tailPos *Position) (int, int) {
	diffX, diffY := xHead-tailPos.x, yHead-tailPos.y

	switch {
	case 2 == diffX && 1 == diffY || 1 == diffX && 2 == diffY || 2 == diffX && 2 == diffY:
		return 1, 1
	case -2 == diffX && -1 == diffY || -1 == diffX && -2 == diffY || -2 == diffX && -2 == diffY:
		return -1, -1
	case 2 == diffX && 0 == diffY:
		return 1, 0
	case 0 == diffX && 2 == diffY:
		return 0, 1
	case -2 == diffX && 0 == diffY:
		return -1, 0
	case 0 == diffX && -2 == diffY:
		return 0, -1
	case 2 == diffX && -1 == diffY || 1 == diffX && -2 == diffY || 2 == diffX && -2 == diffY:
		return 1, -1
	case -2 == diffX && 1 == diffY || -1 == diffX && 2 == diffY || -2 == diffX && 2 == diffY:
		return -1, 1
	default:
		return 0, 0
	}
}

func day9PrintBoard(snake *list.List, positions *map[Position]int) {
	// Create an empty board
	board := [15][15]string{}

	// Set the head and tail positions on the board
	head := snake.Front().Value.(Position)
	tail := snake.Back().Value.(Position)
	board[head.x][head.y] = "H"
	board[tail.x][tail.y] = "T"

	// Print the board to the console
	for i, row := range board {
		for j, col := range row {
			// If the position is empty, print a dot
			if _, ok := (*positions)[Position{j + 1, i + 1, Tail}]; ok {
				fmt.Print("#")
			} else if col == "" {
				fmt.Print(".")
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}

}

func day9IsValidPosition(x, y int, tail *Position) bool {
	maxDistance := 1
	if abs(x-tail.x) > maxDistance || abs(y-tail.y) > maxDistance {
		return false
	}

	return true
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func day9GetDirection(line string) Direction {
	parts := strings.Split(line, " ")
	direction := parts[0]
	distance, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}

	switch direction {
	case "U":
		return Direction{x: 0, y: -1, steps: distance}
	case "D":
		return Direction{x: 0, y: 1, steps: distance}
	case "L":
		return Direction{x: -1, y: 0, steps: distance}
	case "R":
		return Direction{x: 1, y: 0, steps: distance}
	default:
		return Direction{x: 0, y: 0, steps: 0}
	}
}

func day8Task2() int {
	input := readInput("assets/input8.txt")
	trees := day8PrepTrees(input)

	max := math.MinInt
	for i, row := range trees {
		for j := range row {
			score := treeScore(&trees, i, j)
			if max < score {
				max = score
			}
		}
	}

	return max
}

func treeScore(trees *[][]int, row, col int) int {
	rowLen := len(*trees)
	colLen := len((*trees)[0])
	scores := [4]int{}

	for i := row - 1; i >= 0; i-- {
		scores[0]++
		if (*trees)[row][col] <= (*trees)[i][col] {
			break
		}
	}

	for i := row + 1; i < rowLen; i++ {
		scores[1]++
		if (*trees)[row][col] <= (*trees)[i][col] {
			break
		}
	}

	for i := col + 1; i < colLen; i++ {
		scores[2]++
		if (*trees)[row][col] <= (*trees)[row][i] {
			break
		}
	}

	for i := col - 1; i >= 0; i-- {
		scores[3]++
		if (*trees)[row][col] <= (*trees)[row][i] {
			break
		}
	}

	score := 1
	for _, val := range scores {
		score *= val
	}

	return score
}

func day8Task1() int {
	input := readInput("assets/input8.txt")
	trees := day8PrepTrees(input)
	rowLimit := len(trees) - 1
	colLimit := len(trees[0]) - 1

	count := 0
	for i := 1; i < rowLimit; i++ {
		for j := 1; j < colLimit; j++ {
			if day8IsVisible(&trees, i, j) {
				count++
			}
		}
	}

	frameTrees := 2*rowLimit + 2*colLimit
	return frameTrees + count
}

func day8PrepTrees(input string) [][]int {
	rows := strings.Split(input, "\n")
	trees := make([][]int, len(rows))
	for i, row := range rows {
		for _, col := range row {
			if n, err := strconv.Atoi(string(col)); err == nil {
				trees[i] = append(trees[i], n)
			}
		}

	}
	return trees
}
func day8IsVisible(rows *[][]int, row, col int) bool {
	rowLen := len(*rows)
	colLen := len((*rows)[0])

	i := row - 1
	for ; i >= 0; i-- {
		if (*rows)[row][col] <= (*rows)[i][col] {
			break
		}
	}
	if i == -1 {
		return true
	}

	for i = row + 1; i < rowLen; i++ {
		if (*rows)[row][col] <= (*rows)[i][col] {
			break
		}
	}
	if i == rowLen {
		return true
	}

	for i = col + 1; i < colLen; i++ {
		if (*rows)[row][col] <= (*rows)[row][i] {
			break
		}
	}
	if i == colLen {
		return true
	}

	for i = col - 1; i >= 0; i-- {
		if (*rows)[row][col] <= (*rows)[row][i] {
			break
		}
	}

	return i == -1
}

func day7Task2() int {
	input := readInput("assets/input7.txt")
	dirs := day7GetDirsInfo(input)

	var rootIndex int
	for i, dir := range dirs {
		if dir.Path == "/" {
			rootIndex = i
			break
		}

	}

	free := 70000000 - dirs[rootIndex].Size
	required := 30000000 - free
	minSize := math.MaxInt
	for _, dir := range dirs {
		if dir.Size >= required && dir.Size < minSize {
			minSize = dir.Size
		}
	}

	return minSize
}

func day7Task1() int {
	input := readInput("assets/input7.txt")
	dirs := day7GetDirsInfo(input)

	sum := 0
	for _, dir := range dirs {
		if dir.Size <= 100000 {
			sum += dir.Size
		}
	}

	return sum
}

type Dir struct {
	Path string
	Size int
}

func day7GetDirsInfo(input string) []Dir {
	currDir := ""
	dirSizes := map[string]int{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "$ cd") {
			if line == "$ cd /" {
				currDir = "/"
			} else if line == "$ cd .." {
				currDir = currDir[:strings.LastIndex(currDir, "/")]
			} else {
				if currDir != "/" {
					currDir += "/" + strings.TrimPrefix(line, "$ cd ")
				} else {
					currDir += strings.TrimPrefix(line, "$ cd ")
				}
			}
		} else if line == "$ ls" {
			continue
		} else if strings.HasPrefix(line, "dir ") {
			suffix := "/" + strings.TrimPrefix(line, "dir ")
			if currDir != "/" {
				dirSizes[currDir+suffix] += 0
			} else {
				dirSizes[suffix] += 0
			}
		} else {
			items := strings.Fields(line)
			if size, err := strconv.Atoi(items[0]); err == nil {
				dirSizes[currDir] += size
			}
		}
	}

	dirsSorted := make([]Dir, len(dirSizes))

	i := 0
	for k := range dirSizes {
		dirsSorted[i].Path = k
		dirsSorted[i].Size = dirSizes[k]
		i++
	}

	sort.Slice(dirsSorted, func(i, j int) bool {
		iCount := strings.Count(dirsSorted[i].Path, "/")
		jCount := strings.Count(dirsSorted[j].Path, "/")
		if iCount == jCount {
			return len(dirsSorted[i].Path) > len(dirsSorted[j].Path)
		}
		return iCount > jCount
	})

	isDirect := func(s1 string, s2 string) bool {
		if !strings.HasPrefix(s1, s2) {
			return false
		}

		dir := strings.Replace(s1, s2, "", 1)
		return dir == "" || ((strings.HasPrefix(dir, "/") && strings.Count(dir, "/") == 1) || strings.Count(dir, "/") == 0)
	}

	for i := 0; i < len(dirsSorted)-1; i++ {
		for j := i + 1; j < len(dirsSorted); j++ {
			if strings.Count(dirsSorted[i].Path, "/")-strings.Count(dirsSorted[j].Path, "/") > 1 {
				break
			}
			if isDirect(dirsSorted[i].Path, dirsSorted[j].Path) {
				dirsSorted[j].Size += dirsSorted[i].Size
			}
		}
	}

	return dirsSorted
}

func day6Task2() int {
	input := readInput("assets/input6.txt")
	return day6FindPosNUniqChars(input, 14)
}

func day6Task1() int {
	input := readInput("assets/input6.txt")
	return day6FindPosNUniqChars(input, 4)
}

func day6FindPosNUniqChars(input string, n int) int {
	seenChars := make(map[rune]bool, 0)
	count := 0
	for i := 0; i < len(input); i++ {
		if !seenChars[rune(input[i])] {
			seenChars[rune(input[i])] = true
			count++
			if count == n {
				return i + 1
			}
		} else {
			i -= count
			seenChars = map[rune]bool{}
			count = 0
		}
	}
	return -1
}

func day5Task2() string {
	input := readInputWithSpace("assets/input5.txt")
	inputs := strings.Split(input, "\n\n")
	stacks := map[int][]string{}
	lines := strings.Split(inputs[0], "\n")
	for i := len(lines) - 2; i >= 0; i-- {
		if !strings.HasPrefix(lines[i], " 1") {
			lines[i] = strings.ReplaceAll(lines[i], "    ", "[]")
			lines[i] = strings.TrimSpace(lines[i])
			count := 0
			for j, stack := range strings.Split(lines[i], " ") {
				if _, ok := stacks[j+1]; !ok {
					stacks[j+1] = make([]string, 0)
				}
				if count > 0 {
					stack = strings.ReplaceAll(stack, "[]", "")
					stack = strings.TrimSpace(stack)
					stacks[j+1+count] = append(stacks[j+1+count], stack)
				} else {
					count = strings.Count(stack, "[]")
					stack = strings.ReplaceAll(stack, "[]", "")
					stack = strings.TrimSpace(stack)
					stacks[j+1] = append(stacks[j+1], stack)
				}
			}
		}
	}

	inputs[1] = strings.TrimSpace(inputs[1])
	for _, line := range strings.Split(inputs[1], "\n") {
		count, from, to := day5MoveInfo(line)
		if items, ok := stacks[from]; ok {
			if count > 0 && len(items) > 0 {
				stacks[to] = append(stacks[to], items[len(items)-count:]...)
				items = items[:len(items)-count]
			}
			stacks[from] = items
		}
	}

	var code string
	for i := 1; i <= len(stacks); i++ {
		ch := stacks[i][len(stacks[i])-1]
		ch = strings.TrimPrefix(ch, "[")
		code += strings.TrimSuffix(ch, "]")
	}

	return code
}

func day5Task1() string {
	input := readInputWithSpace("assets/input5.txt")
	inputs := strings.Split(input, "\n\n")
	stacks := map[int][]string{}
	lines := strings.Split(inputs[0], "\n")
	for i := len(lines) - 2; i >= 0; i-- {
		if !strings.HasPrefix(lines[i], " 1") {
			lines[i] = strings.ReplaceAll(lines[i], "    ", "[]")
			lines[i] = strings.TrimSpace(lines[i])
			count := 0
			for j, stack := range strings.Split(lines[i], " ") {
				if _, ok := stacks[j+1]; !ok {
					stacks[j+1] = make([]string, 0)
				}
				if count > 0 {
					stack = strings.ReplaceAll(stack, "[]", "")
					stack = strings.TrimSpace(stack)
					stacks[j+1+count] = append(stacks[j+1+count], stack)
				} else {
					count = strings.Count(stack, "[]")
					stack = strings.ReplaceAll(stack, "[]", "")
					stack = strings.TrimSpace(stack)
					stacks[j+1] = append(stacks[j+1], stack)
				}
			}
		}
	}

	inputs[1] = strings.TrimSpace(inputs[1])
	for _, line := range strings.Split(inputs[1], "\n") {
		count, from, to := day5MoveInfo(line)
		if items, ok := stacks[from]; ok {
			for count > 0 && len(items) > 0 {
				count--
				stacks[to] = append(stacks[to], items[len(items)-1])
				items = items[:len(items)-1]
			}
			stacks[from] = items
		}
	}

	var code string
	for i := 1; i <= len(stacks); i++ {
		ch := stacks[i][len(stacks[i])-1]
		ch = strings.TrimPrefix(ch, "[")
		code += strings.TrimSuffix(ch, "]")
	}

	return code
}

func day5MoveInfo(line string) (count int, from int, to int) {
	var err error
	lin := strings.TrimLeft(line, "move ")

	inst := strings.Split(lin, " from ")
	if count, err = strconv.Atoi(inst[0]); err != nil {
		log.Panic(err)
	}
	move := strings.Split(inst[1], " to ")
	if from, err = strconv.Atoi(move[0]); err != nil {
		log.Panic(err)
	}
	if to, err = strconv.Atoi(move[1]); err != nil {
		log.Panic(err)
	}
	return
}

func day4Task2() int {
	input := readInput("assets/input4.txt")
	containCount := 0
	for _, line := range strings.Split(input, "\n") {
		sects := strings.Split(line, ",")
		r1 := day4NewRange(sects[0])
		r2 := day4NewRange(sects[1])
		if day4ContainsOverlap(r1, r2) || day4ContainsOverlap(r2, r1) {
			containCount++
		}
	}

	return containCount
}

func day4Task1() int {
	input := readInput("assets/input4.txt")
	containCount := 0
	for _, line := range strings.Split(input, "\n") {
		sects := strings.Split(line, ",")
		r1 := day4NewRange(sects[0])
		r2 := day4NewRange(sects[1])
		if day4ContainsFull(r1, r2) || day4ContainsFull(r2, r1) {
			containCount++
		}
	}

	return containCount
}

func day4NewRange(section string) *Range {
	rangeStr := strings.Split(section, "-")
	if min, err := strconv.Atoi(rangeStr[0]); err == nil {
		if max, err := strconv.Atoi(rangeStr[1]); err == nil {
			return &Range{min, max}
		}
	}
	return nil
}

func day4ContainsFull(r1 *Range, r2 *Range) bool {
	return r1.Min <= r2.Min && r2.Max <= r1.Max
}

func day4ContainsOverlap(r1 *Range, r2 *Range) bool {
	return r1.Min <= r2.Min && r2.Min <= r1.Max || r1.Min <= r2.Max && r2.Max <= r1.Max
}

func day3Task2() int {
	input := readInput("assets/input3.txt")
	prioritySum := 0
	chars := []rune{}
	for g, line := range strings.Split(input, "\n") {
		if g%3 == 0 {
			chars = []rune(line)
			continue
		}

		matchChars := []rune{}
		for _, ch := range chars {
			if strings.ContainsRune(line, ch) && !strings.ContainsRune(string(matchChars), ch) {
				matchChars = append(matchChars, ch)
			}
		}
		chars = matchChars

		if len(chars) == 1 {
			prioritySum += day3CalcPriority(chars[0])
		}
	}

	return prioritySum
}

func day3Task1() int {
	input := readInput("assets/input3.txt")
	prioritySum := 0
	for _, line := range strings.Split(input, "\n") {
		compLen := len(line) / 2
	BREAK:
		for i := 0; i < compLen; i++ {
			for j := compLen; j < len(line); j++ {
				if line[i] == line[j] {
					prioritySum += day3CalcPriority(rune(line[i]))
					break BREAK
				}
			}

		}

	}

	return prioritySum
}

func day3CalcPriority(item rune) int {
	if item >= 97 { // a - 97, b - 98, ...
		return int(item) - 96
	} else { // A - 65, B - 66, ...
		return int(item) - 38
	}
}

func day2Task2() int {
	input := readInput("assets/input2.txt")
	score := 0
	for _, line := range strings.Split(input, "\n") {
		// Opponent: A for Rock, B for Paper, and C for Scissors
		// Player outcome: X - lose, Y - draw, and Z - win
		strategy := strings.Split(line, " ")
		outcome := strategy[1]
		opponent := strategy[0]
		score += day2CalcExpectScore(outcome, opponent)
	}

	return score
}

func day2CalcExpectScore(outcome, opponent string) int {
	var player string
	if outcome == "Z" { // Player wins
		if opponent == "A" {
			player = "B"
		} else if opponent == "B" {
			player = "C"
		} else {
			player = "A"
		}
	} else if outcome == "X" { // Player loses
		if opponent == "A" {
			player = "C"
		} else if opponent == "B" {
			player = "A"
		} else {
			player = "B"
		}
	} else { // Draw
		player = opponent
	}
	return day2RoundScore(player, opponent) + day2PlayerScore(player)
}

func day2Task1() int {
	input := readInput("assets/input2.txt")
	score := 0
	for _, line := range strings.Split(input, "\n") {
		// Opponent: A for Rock, B for Paper, and C for Scissors
		// Player: X for Rock, Y for Paper, and Z for Scissors
		strategy := strings.Split(line, " ")
		var player string
		if strategy[1] == "X" {
			player = "A"
		} else if strategy[1] == "Y" {
			player = "B"
		} else {
			player = "C"
		}
		opponent := strategy[0]
		score += day2RoundScore(player, opponent) + day2PlayerScore(player)
	}

	return score
}

func day2PlayerScore(player string) int {
	if player == "C" {
		return 3
	} else if player == "B" {
		return 2
	} else {
		return 1
	}
}

func day2RoundScore(p1, p2 string) int {
	if (p1 == "A" && p2 == "C") || (p1 == "B" && p2 == "A") || (p1 == "C" && p2 == "B") {
		return 6
	} else if (p1 == "C" && p2 == "A") || (p1 == "A" && p2 == "B") || (p1 == "B" && p2 == "C") {
		return 0
	} else {
		return 3
	}
}

func day1Task2() int64 {
	input := readInput("assets/input1.txt")
	calories := []int64{}

	for _, lines := range strings.Split(input, "\n\n") {
		var sum int64
		for _, line := range strings.Split(lines, "\n") {
			if n, err := strconv.ParseInt(line, 10, 64); err == nil {
				sum += n
			}
		}
		calories = append(calories, sum)
	}

	nth := 3
	indeces := map[int]bool{}
	for k := 0; k < nth; k++ {
		index := 0
		var max int64
		for i := 0; i < len(calories); i++ {
			if _, ok := indeces[i]; ok {
				continue
			}
			if calories[i] > max {
				max = calories[i]
				index = i
			}
		}
		indeces[index] = true
	}

	var sum int64
	for i := range indeces {
		sum += calories[i]
	}

	return sum
}

func day1Task1() int64 {
	input := readInput("assets/input1.txt")
	calories := make([]int64, 0)

	for _, lines := range strings.Split(input, "\n\n") {
		var sum int64
		for _, line := range strings.Split(lines, "\n") {
			if n, err := strconv.ParseInt(line, 10, 64); err == nil {
				sum += n
			}
		}
		calories = append(calories, sum)
	}

	var max int64
	for i := 0; i < len(calories); i++ {
		if calories[i] > max {
			max = calories[i]
		}
	}

	return max
}

// Useful structs
type Range struct {
	Min int
	Max int
}

// Read input file
func readInput(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(bytes.NewBuffer(file).String())
}

func readInputWithSpace(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewBuffer(file).String()
}
