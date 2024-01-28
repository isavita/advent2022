package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ X, Y int }

type Dir int8

const (
	N Dir = iota
	E
	S
	W
)

func (dir Dir) Rotate(direction rune) Dir {
	switch direction {
	case 'R':
		return (dir + 1) % 4
	case 'L':
		return (dir - 1 + 4) % 4
	}
	return dir
}

func (dir Dir) Points() int {
	return (int(dir) + 3) % 4
}

type Movement struct {
	Steps  int
	Rotate rune
}

var (
	Map       = map[P]bool{}
	Size      int
	Movements []Movement
	Dirs      = [4]P{
		{-1, 0}, // N
		{0, 1},  // E
		{1, 0},  // S
		{0, -1}, // W
	}
)

type Human struct {
	Curr   P
	Facing Dir
}

func main() {
	parse()

	human := Human{
		Curr:   P{0, Size},
		Facing: E,
	}

	for _, mov := range Movements {
		human.Facing = human.Facing.Rotate(mov.Rotate)
		for i := 0; i < mov.Steps; i++ {
			if moved := human.Walk(); !moved {
				break
			}
		}
	}

	fmt.Println(1000*(human.Curr.X+1) + 4*(human.Curr.Y+1) + human.Facing.Points())
}

func parse() {
	file, err := os.Open("day22/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()

		if line == "" {
			break
		}

		if r == 0 {
			Size = len(line) / 3
		}

		for c, char := range line {

			switch char {
			case ' ':
				continue
			case '#':
				Map[P{r, c}] = true
			case '.':
				Map[P{r, c}] = false
			}
		}
	}
	scanner.Scan()
	Movements = parsePath(scanner.Text())
}

func parsePath(path string) []Movement {
	movements := []Movement{}
	acc := 0
	for _, char := range path {

		switch char {

		case 'R':
			movements = append(movements, Movement{Steps: acc})
			acc = 0
			movements = append(movements, Movement{Rotate: 'R'})

		case 'L':
			movements = append(movements, Movement{Steps: acc})
			acc = 0
			movements = append(movements, Movement{Rotate: 'L'})

		default:
			acc = 10*acc + int(char-'0')

		}
	}
	movements = append(movements, Movement{Steps: acc})
	return movements
}

func (h *Human) Walk() bool {
	dirDelta := Dirs[h.Facing]
	next := P{h.Curr.X + dirDelta.X, h.Curr.Y + dirDelta.Y}
	if wall, ok := Map[next]; ok {
		if wall {
			return false
		}
		h.Curr = next
		return true
	}

	next, nextFacing := crossBorder(next, h.Facing)

	if Map[next] {
		return false
	}

	h.Curr = next
	h.Facing = nextFacing
	return true
}

func crossBorder(n P, dir Dir) (P, Dir) {
	x, y := n.X, n.Y

	switch {

	// A --> H
	case x == -1 && y < 2*Size:
		return P{y + 2*Size, x + 1}, E

	// B --> F
	case x == -1 && y >= 2*Size:
		return P{x + 4*Size, y - 2*Size}, N

	// C --> L
	case x == Size && dir == S:
		return P{y - Size, x + Size - 1}, W

	// D --> J
	case x == 2*Size-1 && dir == N:
		return P{y + Size, x - Size + 1}, E

	// E --> K
	case x == 3*Size && dir == S:
		return P{y + 2*Size, x - 2*Size - 1}, W

	// F --> B
	case x == 4*Size:
		return P{x - 4*Size, y + 2*Size}, S

	// G --> I
	case y == -1 && x < 3*Size:
		return P{3*Size - 1 - x, y + Size + 1}, E

	// H --> A
	case y == -1 && x >= 3*Size:
		return P{y + 1, x - 2*Size}, S

	// I --> G
	case y == Size-1 && x < Size:
		return P{3*Size - 1 - x, y - Size + 1}, E

	// J --> D
	case y == Size-1 && x >= Size && dir == W:
		return P{y + Size + 1, x - Size}, S

	// K --> E
	case y == Size && dir == E:
		return P{y + 2*Size - 1, x - 2*Size}, N

	// L --> C
	case y == 2*Size && x < 2*Size && dir == E:
		return P{y - Size - 1, x + Size}, N

	// M --> N
	case y == 2*Size && x >= 2*Size:
		return P{3*Size - 1 - x, y + Size - 1}, W

	// N --> M
	case y == 3*Size:
		return P{3*Size - 1 - x, y - Size - 1}, W
	}

	panic("not a border crossing")
}
