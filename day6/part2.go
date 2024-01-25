package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := readAll("day6/input.txt")
	fmt.Println(firstNUnique(s, 14))
}

func firstNUnique(s string, n int) int {
	for i := n; i < len(s); i++ {
		b := []byte(s[i-n : i])
		if len(b) == len(SetOf(b)) {
			return i
		}
	}
	return -1
}

func readAll(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(file))
}

func SetOf(b []byte) []byte {
	m := make(map[byte]bool)
	for _, c := range b {
		m[c] = true
	}
	b = make([]byte, 0, len(m))
	for c := range m {
		b = append(b, c)
	}
	return b
}
