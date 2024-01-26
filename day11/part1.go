package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items       []int
	operation   func(int) int
	div         int
	next        [2]int
	inspections int
}

func parse(s string) *monkey {
	m := &monkey{}
	lines := strings.Split(s, "\n")
	for _, item := range strings.Split(strings.Split(lines[1], ": ")[1], ", ") {
		m.items = append(m.items, int(toInt(item)))
	}
	f := strings.Fields(strings.Split(lines[2], "= ")[1])
	switch f[1] {
	case "+":
		switch f[2] {
		case "old":
			m.operation = func(old int) int { return old + old }
		default:
			m.operation = func(old int) int { return old + toInt(f[2]) }
		}
	case "*":
		switch f[2] {
		case "old":
			m.operation = func(old int) int { return old * old }
		default:
			m.operation = func(old int) int { return old * toInt(f[2]) }
		}
	}
	fmt.Sscanf(lines[3], " Test: divisible by %d", &m.div)
	fmt.Sscanf(lines[4], " If true: throw to monkey %d", &m.next[0])
	fmt.Sscanf(lines[5], " If false: throw to monkey %d", &m.next[1])
	return m
}

func monkeyBusiness(monkeys []*monkey, rounds int, worry bool) int {
	div := 1
	for _, m := range monkeys {
		div *= m.div
	}

	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				m.inspections++
				item := m.operation(m.items[0])
				if worry {
					item %= int(div)
				} else {
					item /= 3
				}
				if item%int(m.div) == 0 {
					monkeys[m.next[0]].items = append(monkeys[m.next[0]].items, item)
				} else {
					monkeys[m.next[1]].items = append(monkeys[m.next[1]].items, item)
				}
				m.items = m.items[1:]
			}
		}
	}
	inspections := []int{}
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	return inspections[0] * inspections[1]
}

func main() {
	var monkeys []*monkey
	s := readAll("day11/input.txt")
	for _, m := range strings.Split(s, "\n\n") {
		monkeys = append(monkeys, parse(m))
	}
	fmt.Println(monkeyBusiness(monkeys, 20, false))
}

func readAll(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
