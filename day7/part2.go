package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	root := []string{""}
	dirs := map[string]int{}
	files := map[string]int{}
	var curr []string
	s := scanAll()
	for s.Scan() {
		txt := strings.Fields(s.Text())
		if txt[0] == "$" {
			if txt[1] == "cd" {
				if txt[2] == "/" {
					curr = root
				} else if txt[2] == ".." {
					curr = curr[:len(curr)-1]
				} else {
					curr = append(curr, txt[2])
				}
				dirs[strings.Join(curr, "/")] = 0
			}
		} else {
			if txt[0] != "dir" {
				files[strings.Join(append(curr, txt[1]), "/")] = toInt(txt[0])
			}
		}
	}

	for f, s := range files {
		path := strings.Split(f, "/")
		for i := 1; i < len(path); i++ {
			dirs[strings.Join(path[:i], "/")] += s
		}
	}

	var sortedSizes []int
	for _, s := range dirs {
		sortedSizes = append(sortedSizes, s)
	}

	sort.Ints(sortedSizes)
	total, want := 70000000, 30000000
	available := total - dirs[""]
	fmt.Println(sortedSizes[sort.SearchInts(sortedSizes, want-available)])
}

func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

func scanAll() *bufio.Scanner {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}
