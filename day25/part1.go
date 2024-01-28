package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum := 0
	s := scanAll()
	for s.Scan() {
		sum += fromSnafu(s.Text())
	}
	fmt.Println(toSnafu(sum))
}

func fromSnafu(s string) int {
	var n int
	for i := range s {
		n *= 5
		switch s[i] {
		case '=':
			n -= 2
		case '-':
			n--
		default:
			n += int(s[i]) - '0'
		}
	}
	return n
}

func toSnafu(n int) string {
	var b []byte
	for ; n > 0; n /= 5 {
		switch n % 5 {
		case 3:
			n += 5
			b = append(b, '=')
		case 4:
			n += 5
			b = append(b, '-')
		default:
			b = append(b, byte('0'+n%5))
		}
	}
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-1-i], b[i]
	}
	return string(b)
}

func scanAll() *bufio.Scanner {
	file, err := os.Open("day25/input.txt")
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}
