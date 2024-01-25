package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := strings.Split(readAll("day5/input.txt"), "\n\n")
	input := strings.Split(s[0], "\n")
	stacks := make([][]byte, (len(input[0])+1)/4)
	for _, line := range input {
		for i, b := range []byte(line) {
			if b >= 'A' && b <= 'Z' {
				stacks[(i-1)/4] = append(stacks[(i-1)/4], b)
			}
		}
	}

	steps := strings.Split(s[1], "\n")
	fmt.Println(move(stacks, steps))
}

func move(st [][]byte, steps []string) string {
	stacks := make([][]byte, len(st))
	// Reverse stacks, make them bottom to top.
	for i := range st {
		stacks[i] = make([]byte, len(st[i]))
		for j := range st[i] {
			stacks[i][j] = st[i][len(st[i])-j-1]
		}
	}

	for _, step := range steps {
		var n, from, to int
		fmt.Sscanf(step, "move %d from %d to %d", &n, &from, &to)
		from--
		to--
		for i := 0; i < n; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	b := make([]byte, len(stacks))
	for i := range stacks {
		b[i] = stacks[i][len(stacks[i])-1]
	}
	return string(b)
}

func readAll(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(file))
}
