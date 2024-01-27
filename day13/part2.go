package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	s := readAll("day13/input.txt")
	var packets []any
	for _, pair := range strings.Split(s, "\n\n") {
		sp := strings.Split(pair, "\n")
		var first, second any
		json.Unmarshal([]byte(sp[0]), &first)
		json.Unmarshal([]byte(sp[1]), &second)
		packets = append(packets, first, second)
	}

	var divider1, divider2 any
	json.Unmarshal([]byte("[[2]]"), &divider1)
	json.Unmarshal([]byte("[[6]]"), &divider2)
	packets = append(packets, divider1, divider2)
	sort.Slice(packets, func(i, j int) bool { return compare(packets[i], packets[j]) < 0 })
	divider1Pos := sort.Search(len(packets), func(i int) bool { return compare(packets[i], divider1) >= 0 })
	divider2Pos := sort.Search(len(packets), func(i int) bool { return compare(packets[i], divider2) >= 0 })
	fmt.Println((divider1Pos + 1) * (divider2Pos + 1))
}

func readAll(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func compare(a, b any) int {
	_, anum := a.(float64)
	_, bnum := b.(float64)
	switch {
	case anum && bnum:
		return sign(int(a.(float64)) - int(b.(float64)))
	case anum:
		return compare([]any{a}, b)
	case bnum:
		return compare(a, []any{b})
	default:
		aa, bb := a.([]any), b.([]any)
		for i := 0; i < len(aa) && i < len(bb); i++ {
			if c := compare(aa[i], bb[i]); c != 0 {
				return c
			}
		}
		return sign(len(aa) - len(bb))
	}
}

func sign(n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	return 1
}
