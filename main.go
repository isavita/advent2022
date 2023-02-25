package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("day1\n answer1: %d\n answer2: %d", Day1Task1(), Day1Task2())
	fmt.Printf("\nday2\n answer1: %d\n answer2: %d", Day2Task1(), Day2Task2())
	fmt.Printf("\nday3\n answer1: %d\n answer2: %d", Day3Task1(), Day3Task2())
	fmt.Printf("\nday4\n answer1: %d\n answer2: %d", Day4Task1(), Day4Task2())
	fmt.Printf("\nday5\n answer1: %s\n answer2: %s", Day5Task1(), Day5Task2())
	fmt.Printf("\nday6\n answer1: %d\n answer2: %d", Day6Task1(), Day6Task2())
	fmt.Printf("\nday7\n answer1: %d\n answer2: %d", Day7Task1(), Day7Task2())
	fmt.Printf("\nday8\n answer1: %d\n answer2: %d", Day8Task1(), Day8Task2())
	fmt.Printf("\nday9\n answer1: %d\n answer2: %d", Day9Task1(), Day9Task2())
	fmt.Printf("\nday10\n answer1: %d\n answer2: %s%s", Day10Task1(), "RBPARAGF", Day10Task2())
	fmt.Printf("\nday11\n answer1: %d\n answer2: %d", Day11Task1(), Day11Task2())
	fmt.Printf("\nday12\n answer1: %d\n answer2: %d", Day12Task1(), Day12Task2())
	fmt.Printf("\nday13\n answer1: %d\n answer2: %d", Day13Task1(), Day13Task2())
	fmt.Printf("\nday14\n answer1: %d\n answer2: %d", Day14Task1(), Day14Task2())
}

func Day14Task2() int {
	input := readInput("assets/input14.txt")
	sc := bufio.NewScanner(strings.NewReader(input))

	cave := make(map[point]rune)

	var maxY, maxX, minX int = 0, 0, 500

	//Generate cave
	for sc.Scan() {
		points := strings.Split(sc.Text(), " -> ")
		for i := range points[:len(points)-1] {
			from := strings.Split(points[i], ",")
			to := strings.Split(points[i+1], ",")
			fromX, _ := strconv.Atoi(from[0])
			fromY, _ := strconv.Atoi(from[1])
			toX, _ := strconv.Atoi(to[0])
			toY, _ := strconv.Atoi(to[1])

			cave[point{toX, toY}] = '#'
			cave[point{fromX, fromY}] = '#'
			if toY > maxY {
				maxY = toY
			}
			if toX > maxX {
				maxX = toX
			}
			if toX < minX {
				minX = toX
			}

			for fromX != toX || fromY != toY {
				cave[point{fromX, fromY}] = '#'
				switch {
				case fromX < toX:
					fromX++
				case fromY < toY:
					fromY++
				case fromX > toX:
					fromX--
				case fromY > toY:
					fromY--
				}
			}
			if fromY > maxY {
				maxY = fromY
			}
			if fromX > maxX {
				maxX = fromX
			}
		}
	}
	for i := minX - 500; i < maxX+500; i++ {
		cave[point{i, maxY + 2}] = '#'
	}

	var sand int = 0
	for {
		newSand := point{500, 0}
		if cave[newSand] == 'o' {
			break
		}
		for {
			cave[newSand] = '"'
			if cave[point{newSand.x, newSand.y + 1}] < '#' {
				newSand.y++
			} else if cave[point{newSand.x - 1, newSand.y + 1}] < '#' {
				newSand.y++
				newSand.x--
			} else if cave[point{newSand.x + 1, newSand.y + 1}] < '#' {
				newSand.y++
				newSand.x++
			} else {
				cave[newSand] = 'o'
				sand++
				break
			}
		}
	}

	return sand
}

func Day14Task1() int {
	input := readInput("assets/input14.txt")
	sc := bufio.NewScanner(strings.NewReader(input))

	cave := make(map[point]rune)

	var maxY, maxX int

	//Generate cave
	for sc.Scan() {
		points := strings.Split(sc.Text(), " -> ")
		for i := range points[:len(points)-1] {
			from := strings.Split(points[i], ",")
			to := strings.Split(points[i+1], ",")
			fromX, _ := strconv.Atoi(from[0])
			fromY, _ := strconv.Atoi(from[1])
			toX, _ := strconv.Atoi(to[0])
			toY, _ := strconv.Atoi(to[1])

			cave[point{toX, toY}] = '#'
			cave[point{fromX, fromY}] = '#'
			if toY > maxY {
				maxY = toY
			}
			if toX > maxX {
				maxX = toX
			}

			for fromX != toX || fromY != toY {
				cave[point{fromX, fromY}] = '#'
				switch {
				case fromX < toX:
					fromX++
				case fromY < toY:
					fromY++
				case fromX > toX:
					fromX--
				case fromY > toY:
					fromY--
				}
			}
			if fromY > maxY {
				maxY = fromY
			}
			if fromX > maxX {
				maxX = fromX
			}
		}
	}

	intoVoid := false
	var sand int = 0
	for !intoVoid {
		newSand := point{500, 0}
		for {
			cave[newSand] = '"'
			if newSand.y+1 > maxY {
				intoVoid = true
				break
			}
			if cave[point{newSand.x, newSand.y + 1}] < '#' {
				newSand.y++
			} else if cave[point{newSand.x - 1, newSand.y + 1}] < '#' {
				newSand.y++
				newSand.x--
			} else if cave[point{newSand.x + 1, newSand.y + 1}] < '#' {
				newSand.y++
				newSand.x++
			} else {
				cave[newSand] = 'o'
				sand++
				break
			}
		}
	}
	return sand
}

type point struct{ x, y int }

func Day13Task2() int {
	input := readInput("assets/input13.txt")
	lines := strings.Split(input, "\n\n")
	packages := make([]any, 0, len(lines))
	for _, doubleLine := range lines {
		parts := strings.Split(doubleLine, "\n")
		var p1, p2 any
		json.Unmarshal([]byte(parts[0]), &p1)
		json.Unmarshal([]byte(parts[1]), &p2)
		packages = append(packages, p1, p2)
	}

	packages = append(packages, []any{[]any{2.0}}, []any{[]any{6.0}})
	sort.Slice(packages, func(i, j int) bool {
		return Day13Compare(packages[i], packages[j]) < 0
	})

	decoder := 1
	for i, pckg := range packages {
		s := fmt.Sprint(pckg)
		if s == "[[2]]" || s == "[[6]]" {
			decoder *= i + 1
		}
	}

	return decoder
}

func Day13Task1() int {
	input := readInput("assets/input13.txt")
	lines := strings.Split(input, "\n\n")
	sum := 0
	for i, doubleLine := range lines {
		parts := strings.Split(doubleLine, "\n")
		var p1, p2 any
		json.Unmarshal([]byte(parts[0]), &p1)
		json.Unmarshal([]byte(parts[1]), &p2)
		if Day13Compare(p1, p2) <= 0 {
			sum += i + 1
		}
	}

	return sum
}

func Day13Compare(l, r any) int {
	xs, ok1 := l.([]any)
	ys, ok2 := r.([]any)
	switch {
	case !ok1 && !ok2:
		return int(l.(float64) - r.(float64))
	case !ok1:
		xs = []any{l}
	case !ok2:
		ys = []any{r}
	}

	for i := 0; i < len(xs) && i < len(ys); i++ {
		if c := Day13Compare(xs[i], ys[i]); c != 0 {
			return c
		}
	}
	return len(xs) - len(ys)
}

func Day12Task2() int {
	input := readInput("assets/input12.txt")
	lines := strings.Split(input, "\n")
	pq := make(PriorityQueue, 0)
	heightMap := make(map[string]rune)
	xLen, yLen := len(lines), len(lines[0])
	for i, row := range lines {
		for j, ch := range row {
			coord := fmt.Sprintf("%d,%d", i, j)
			if ch == 'E' {
				heap.Push(&pq, &Item{value: coord, priority: 0, index: 0})
			}
			heightMap[coord] = ch
		}
	}

	return Day12FindShortestPath2(pq, heightMap, xLen, yLen)
}

func Day12Task1() int {
	input := readInput("assets/input12.txt")
	lines := strings.Split(input, "\n")
	pq := make(PriorityQueue, 0)
	heightMap := make(map[string]rune)
	xLen, yLen := len(lines), len(lines[0])
	for i, row := range lines {
		for j, ch := range row {
			coord := fmt.Sprintf("%d,%d", i, j)
			if ch == 'S' {
				heap.Push(&pq, &Item{value: coord, priority: 0, index: 0})
			}
			heightMap[coord] = ch
		}
	}

	return Day12FindShortestPath1(pq, heightMap, xLen, yLen)
}

func Day12FindShortestPath2(pq PriorityQueue, heightMap map[string]rune, xLen, yLen int) int {
	visited := make(map[string]bool)
	for pq.Len() > 0 {
		front := heap.Pop(&pq).(*Item)
		coord := front.value.(string)
		if visited[coord] {
			continue
		} else {
			visited[coord] = true
		}

		if heightMap[coord] == 'a' {
			return front.priority
		}

		for _, neighbor := range Day12GetNeighbors(coord, xLen, yLen) {
			distance := Day12Weight(heightMap[coord], heightMap[neighbor])
			if distance >= -1 {
				heap.Push(&pq, &Item{value: neighbor, priority: front.priority + 1, index: 0})
			}
		}

	}

	return -1
}

func Day12FindShortestPath1(pq PriorityQueue, heightMap map[string]rune, xLen, yLen int) int {
	visited := make(map[string]bool)
	for pq.Len() > 0 {
		front := heap.Pop(&pq).(*Item)
		coord := front.value.(string)
		if visited[coord] {
			continue
		} else {
			visited[coord] = true
		}

		if heightMap[coord] == 'E' {
			return front.priority
		}

		for _, neighbor := range Day12GetNeighbors(coord, xLen, yLen) {
			distance := Day12Weight(heightMap[coord], heightMap[neighbor])
			if distance <= 1 {
				heap.Push(&pq, &Item{value: neighbor, priority: front.priority + 1, index: 0})
			}
		}

	}

	return -1
}

func Day12GetNeighbors(coord string, xLen, yLen int) []string {
	x, y := Day12ParseCoord(coord)

	neighbors := make([]string, 0)

	if x > 0 {
		neighbors = append(neighbors, fmt.Sprintf("%d,%d", x-1, y))
	}
	if x < xLen-1 {
		neighbors = append(neighbors, fmt.Sprintf("%d,%d", x+1, y))
	}
	if y > 0 {
		neighbors = append(neighbors, fmt.Sprintf("%d,%d", x, y-1))
	}
	if y < yLen-1 {
		neighbors = append(neighbors, fmt.Sprintf("%d,%d", x, y+1))
	}

	return neighbors
}

func Day12Weight(h1, h2 rune) int {
	if h1 == 'S' {
		h1 = 'a'
	}

	if h2 == 'S' {
		h2 = 'a'
	}

	if h1 == 'E' {
		h1 = 'z'
	}

	if h2 == 'E' {
		h2 = 'z'
	}

	return int(h2 - h1)
}

func Day12ParseCoord(coord string) (int, int) {
	parts := strings.Split(coord, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}

type Item struct {
	value    interface{}
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func Day11Task2() int64 {
	input := readInput("assets/input11.txt")
	monkeysInfo := strings.Split(input, "\n\n")
	monkeys := make([]Monkey, 0, len(monkeysInfo))
	limit := 1
	for _, monkeyInfo := range monkeysInfo {
		monkey := Day11PrepMonkey(monkeyInfo)
		monkeys = append(monkeys, monkey)
		limit *= monkey.divisor
	}

	inspected := make(map[int]int, len(monkeys))
	for round := 0; round < 10000; round++ {
		for i := 0; i < len(monkeys); i++ {
			for _, item := range monkeys[i].items {
				inspected[monkeys[i].id]++
				monkeys[i].items = monkeys[i].items[1:]
				val := Day11Eval(monkeys[i].operation, item) % limit
				if val%monkeys[i].divisor == 0 {
					j := monkeys[i].trueThrowTo
					monkeys[j].items = append(monkeys[j].items, val)
				} else {
					j := monkeys[i].falseThrowTo
					monkeys[j].items = append(monkeys[j].items, val)
				}
			}
		}
	}

	max1, max2 := -1, -1
	for _, count := range inspected {
		if max2 < count {
			max2 = count
		}
		if max1 < max2 {
			max1, max2 = max2, max1
		}
	}

	return int64(max1) * int64(max2)
}

func Day11Task1() int {
	input := readInput("assets/input11.txt")
	monkeysInfo := strings.Split(input, "\n\n")
	monkeys := make([]Monkey, 0, len(monkeysInfo))
	for _, monkeyInfo := range monkeysInfo {
		monkeys = append(monkeys, Day11PrepMonkey(monkeyInfo))
	}

	inspected := make(map[int]int, len(monkeys))
	for round := 0; round < 20; round++ {
		for i := 0; i < len(monkeys); i++ {
			for _, item := range monkeys[i].items {
				inspected[monkeys[i].id]++
				monkeys[i].items = monkeys[i].items[1:]
				val := Day11Eval(monkeys[i].operation, item) / 3
				if val%monkeys[i].divisor == 0 {
					j := monkeys[i].trueThrowTo
					monkeys[j].items = append(monkeys[j].items, val)
				} else {
					j := monkeys[i].falseThrowTo
					monkeys[j].items = append(monkeys[j].items, val)
				}
			}
		}
	}

	max1, max2 := -1, -1
	for _, count := range inspected {
		if max2 < count {
			max2 = count
		}
		if max1 < max2 {
			max1, max2 = max2, max1
		}
	}

	return max1 * max2
}

func Day11Eval(eq string, x int) int {
	parts := strings.Split(eq, " ")
	var lhs, rhs int
	var err error
	if lhs, err = strconv.Atoi(parts[0]); err != nil {
		lhs = x
	}

	if rhs, err = strconv.Atoi(parts[2]); err != nil {
		rhs = x
	}

	switch parts[1] {
	case "*":
		return lhs * rhs
	case "+":
		return lhs + rhs
	default:
		log.Fatal(parts[1])
	}

	return 0
}

type Monkey struct {
	id           int
	items        []int
	operation    string
	divisor      int
	trueThrowTo  int
	falseThrowTo int
}

func Day11PrepMonkey(monkeyInfo string) Monkey {
	monkey := Monkey{}
	info := strings.Split(monkeyInfo, "\n")

	num := strings.TrimSuffix(strings.TrimPrefix(info[0], "Monkey "), ":")
	if id, err := strconv.Atoi(num); err == nil {
		monkey.id = id
	}

	startingItems := strings.TrimPrefix(strings.TrimSpace(info[1]), "Starting items: ")
	items := []int{}
	for _, item := range strings.Split(startingItems, ", ") {
		if n, err := strconv.Atoi(item); err == nil {
			items = append(items, n)
		}
	}
	monkey.items = items

	monkey.operation = strings.TrimPrefix(strings.TrimSpace(info[2]), "Operation: new = ")

	cond := strings.TrimPrefix(strings.TrimSpace(info[3]), "Test: divisible by ")
	if n, err := strconv.Atoi(cond); err == nil {
		monkey.divisor = n
	}

	trueThrowTo := strings.TrimPrefix(strings.TrimSpace(info[4]), "If true: throw to monkey ")
	if n, err := strconv.Atoi(trueThrowTo); err == nil {
		monkey.trueThrowTo = n
	}

	falseThrowTo := strings.TrimPrefix(strings.TrimSpace(info[5]), "If false: throw to monkey ")
	if n, err := strconv.Atoi(falseThrowTo); err == nil {
		monkey.falseThrowTo = n
	}

	return monkey
}

func Day10Task2() string {
	input := readInput("assets/input10.txt")
	regX, cycle := 1, 0
	var sb strings.Builder
	for _, ins := range strings.Split(input, "\n") {
		Day10DrawSpriteAndLine(cycle, regX, &sb)
		cycle++
		if ins != "noop" {
			parts := strings.Split(ins, " ")
			Day10DrawSpriteAndLine(cycle, regX, &sb)
			cycle++
			if n, err := strconv.Atoi(parts[1]); err == nil {
				regX += n
			}
		}
	}

	return sb.String()
}

func Day10DrawSpriteAndLine(cycle, regX int, sb *strings.Builder) {
	if cycle%40 == 0 {
		sb.WriteString("\n")
	}
	if cycle%40 == regX-1 || cycle%40 == regX || cycle%40 == regX+1 {
		sb.WriteString("#")
	} else {
		sb.WriteString(".")
	}
}

func Day10Task1() int {
	input := readInput("assets/input10.txt")
	strengths := map[int]int{
		20:  1,
		60:  1,
		100: 1,
		140: 1,
		180: 1,
		220: 1,
	}
	regX := 1
	cycle := 0
	for _, ins := range strings.Split(input, "\n") {
		cycle++
		if _, ok := strengths[cycle]; ok {
			strengths[cycle] = cycle * regX
		}
		if ins != "noop" {
			cycle++
			if _, ok := strengths[cycle]; ok {
				strengths[cycle] = cycle * regX
			}
			parts := strings.Split(ins, " ")
			if n, err := strconv.Atoi(parts[1]); err == nil {
				regX += n
			}
		}
	}

	sum := 0
	for _, val := range strengths {
		sum += val
	}

	return sum
}

func Day9Task2() int {
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
		direction := Day9GetDirection(line)
		for ; direction.steps > 0; direction.steps -= 1 {
			headSegment := snake.Front()
			snake.Remove(headSegment)
			headPos := headSegment.Value.(Position)
			newX := headPos.x + direction.x
			newY := headPos.y + direction.y
			newTail := Day9UpdateTail(newX, newY, snake, &positions)
			snake = list.New()
			snake.PushFrontList(&newTail)
			snake.PushFront(Position{x: newX, y: newY, segmentType: Head})
		}
	}

	return len(positions)
}

func Day9UpdateTail(newX, newY int, snake *list.List, positions *map[Position]int) list.List {
	newTail := list.New()
	for item := snake.Front(); item != nil; item = item.Next() {
		tailPos := Position{item.Value.(Position).x, item.Value.(Position).y, item.Value.(Position).segmentType}
		if Day9IsValidPosition(newX, newY, &tailPos) {
			newX, newY = tailPos.x, tailPos.y
			newTail.PushBack(tailPos)
			if tailPos.segmentType == 9 {
				(*positions)[Position{x: tailPos.x, y: tailPos.y, segmentType: tailPos.segmentType}] += 1
			}
		} else {
			deltX, deltY := Day9CalcTailUpdate(newX, newY, &tailPos)
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

func Day9Task1() int {
	input := readInput("assets/input9.txt")

	positions := make(map[Position]int)
	snake := list.New()
	snake.PushFront(Position{x: 5, y: 5, segmentType: Tail})
	snake.PushFront(Position{x: 5, y: 5, segmentType: Head})
	positions[Position{x: 5, y: 5, segmentType: Tail}] = 1

	for _, line := range strings.Split(input, "\n") {
		direction := Day9GetDirection(line)
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
			if Day9IsValidPosition(newX, newY, &tailPos) {
				snake.PushBack(tailPos)
				positions[Position{x: tailPos.x, y: tailPos.y, segmentType: Tail}] += 1
			} else {
				deltX, deltY := Day9CalcTailUpdate(newX, newY, &tailPos)
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

func Day9CalcTailUpdate(xHead, yHead int, tailPos *Position) (int, int) {
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

func Day9IsValidPosition(x, y int, tail *Position) bool {
	maxDistance := 1
	if abs(x-tail.x) > maxDistance || abs(y-tail.y) > maxDistance {
		return false
	}

	return true
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func Day9GetDirection(line string) Direction {
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

func Day8Task2() int {
	input := readInput("assets/input8.txt")
	trees := Day8PrepTrees(input)

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

func Day8Task1() int {
	input := readInput("assets/input8.txt")
	trees := Day8PrepTrees(input)
	rowLimit := len(trees) - 1
	colLimit := len(trees[0]) - 1

	count := 0
	for i := 1; i < rowLimit; i++ {
		for j := 1; j < colLimit; j++ {
			if Day8IsVisible(&trees, i, j) {
				count++
			}
		}
	}

	frameTrees := 2*rowLimit + 2*colLimit
	return frameTrees + count
}

func Day8PrepTrees(input string) [][]int {
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
func Day8IsVisible(rows *[][]int, row, col int) bool {
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

func Day7Task2() int {
	input := readInput("assets/input7.txt")
	dirs := Day7GetDirsInfo(input)

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

func Day7Task1() int {
	input := readInput("assets/input7.txt")
	dirs := Day7GetDirsInfo(input)

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

func Day7GetDirsInfo(input string) []Dir {
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

func Day6Task2() int {
	input := readInput("assets/input6.txt")
	return Day6FindPosNUniqChars(input, 14)
}

func Day6Task1() int {
	input := readInput("assets/input6.txt")
	return Day6FindPosNUniqChars(input, 4)
}

func Day6FindPosNUniqChars(input string, n int) int {
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

func Day5Task2() string {
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
		count, from, to := Day5MoveInfo(line)
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

func Day5Task1() string {
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
		count, from, to := Day5MoveInfo(line)
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

func Day5MoveInfo(line string) (count int, from int, to int) {
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

func Day4Task2() int {
	input := readInput("assets/input4.txt")
	containCount := 0
	for _, line := range strings.Split(input, "\n") {
		sects := strings.Split(line, ",")
		r1 := Day4NewRange(sects[0])
		r2 := Day4NewRange(sects[1])
		if Day4ContainsOverlap(r1, r2) || Day4ContainsOverlap(r2, r1) {
			containCount++
		}
	}

	return containCount
}

func Day4Task1() int {
	input := readInput("assets/input4.txt")
	containCount := 0
	for _, line := range strings.Split(input, "\n") {
		sects := strings.Split(line, ",")
		r1 := Day4NewRange(sects[0])
		r2 := Day4NewRange(sects[1])
		if Day4ContainsFull(r1, r2) || Day4ContainsFull(r2, r1) {
			containCount++
		}
	}

	return containCount
}

func Day4NewRange(section string) *Range {
	rangeStr := strings.Split(section, "-")
	if min, err := strconv.Atoi(rangeStr[0]); err == nil {
		if max, err := strconv.Atoi(rangeStr[1]); err == nil {
			return &Range{min, max}
		}
	}
	return nil
}

func Day4ContainsFull(r1 *Range, r2 *Range) bool {
	return r1.Min <= r2.Min && r2.Max <= r1.Max
}

func Day4ContainsOverlap(r1 *Range, r2 *Range) bool {
	return r1.Min <= r2.Min && r2.Min <= r1.Max || r1.Min <= r2.Max && r2.Max <= r1.Max
}

func Day3Task2() int {
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
			prioritySum += Day3CalcPriority(chars[0])
		}
	}

	return prioritySum
}

func Day3Task1() int {
	input := readInput("assets/input3.txt")
	prioritySum := 0
	for _, line := range strings.Split(input, "\n") {
		compLen := len(line) / 2
	BREAK:
		for i := 0; i < compLen; i++ {
			for j := compLen; j < len(line); j++ {
				if line[i] == line[j] {
					prioritySum += Day3CalcPriority(rune(line[i]))
					break BREAK
				}
			}

		}

	}

	return prioritySum
}

func Day3CalcPriority(item rune) int {
	if item >= 97 { // a - 97, b - 98, ...
		return int(item) - 96
	} else { // A - 65, B - 66, ...
		return int(item) - 38
	}
}

func Day2Task2() int {
	input := readInput("assets/input2.txt")
	score := 0
	for _, line := range strings.Split(input, "\n") {
		// Opponent: A for Rock, B for Paper, and C for Scissors
		// Player outcome: X - lose, Y - draw, and Z - win
		strategy := strings.Split(line, " ")
		outcome := strategy[1]
		opponent := strategy[0]
		score += Day2CalcExpectScore(outcome, opponent)
	}

	return score
}

func Day2CalcExpectScore(outcome, opponent string) int {
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
	return Day2RoundScore(player, opponent) + Day2PlayerScore(player)
}

func Day2Task1() int {
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
		score += Day2RoundScore(player, opponent) + Day2PlayerScore(player)
	}

	return score
}

func Day2PlayerScore(player string) int {
	if player == "C" {
		return 3
	} else if player == "B" {
		return 2
	} else {
		return 1
	}
}

func Day2RoundScore(p1, p2 string) int {
	if (p1 == "A" && p2 == "C") || (p1 == "B" && p2 == "A") || (p1 == "C" && p2 == "B") {
		return 6
	} else if (p1 == "C" && p2 == "A") || (p1 == "A" && p2 == "B") || (p1 == "B" && p2 == "C") {
		return 0
	} else {
		return 3
	}
}

func Day1Task2() int64 {
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

func Day1Task1() int64 {
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
