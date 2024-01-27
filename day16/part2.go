package main

import (
	"fmt"
	"os"
	"strings"
)

type valve struct {
	id      string
	flow    int
	tunnels map[string]int
}

func main() {
	valves := map[string]valve{}

	input := readAll("day16/input.txt")
	for _, line := range strings.Split(input, "\n") {
		sp := strings.Split(line, "; ")
		var v valve
		fmt.Sscanf(sp[0], "Valve %s has flow rate=%d", &v.id, &v.flow)
		sp[1] = sp[1][len("tunnel leads to valve"):]
		if strings.HasPrefix(sp[1], "s") {
			sp[1] = sp[1][2:]
		} else {
			sp[1] = sp[1][1:]
		}
		v.tunnels = map[string]int{v.id: 0}
		for _, t := range strings.Split(sp[1], ", ") {
			v.tunnels[t] = 1
		}
		valves[v.id] = v
	}
	for k := range valves {
		for i := range valves {
			for j := range valves {
				dik, okik := valves[i].tunnels[k]
				dkj, okkj := valves[k].tunnels[j]
				if okik && okkj {
					dij, okij := valves[i].tunnels[j]
					if !okij || dij > dik+dkj {
						valves[i].tunnels[j] = dik + dkj
					}
				}
			}
		}
	}
	open := []string{}
	for _, v := range valves {
		if v.flow > 0 {
			open = append(open, v.id)
		}
	}

	max := 0
	for _, d := range divide(len(open)) {
		if len(d[0]) == 0 || len(d[1]) == 0 {
			continue
		}
		mine := make([]string, len(d[0]))
		for i := range d[0] {
			mine[i] = open[d[0][i]]
		}
		elephant := make([]string, len(d[1]))
		for i := range d[1] {
			elephant[i] = open[d[1][i]]
		}

		x := maxPressure(valves, "AA", 26, 0, mine, 0) + maxPressure(valves, "AA", 26, 0, elephant, 0)
		max = Max(max, x)
	}
	fmt.Println(max)
}

func maxPressure(valves map[string]valve, curr string, minute int, pressure int, open []string, d int) int {
	max := pressure
	for _, next := range open {
		newopen := make([]string, 0, len(open)-1)
		for _, v := range open {
			if v != next {
				newopen = append(newopen, v)
			}
		}
		timeLeft := minute - valves[curr].tunnels[next] - 1
		if timeLeft > 0 {
			max = Max(max, maxPressure(valves, next, timeLeft, timeLeft*valves[next].flow+pressure, newopen, d+1))
		}
	}
	return max
}

func divide(l int) [][2][]int {
	if l == 1 {
		return [][2][]int{
			{{}, {0}},
			{{0}, {}},
		}
	}
	d := divide(l - 1)
	r := make([][2][]int, 2*len(d))
	for i := range d {
		r[2*i] = [2][]int{append([]int{l - 1}, d[i][0]...), d[i][1]}
		r[2*i+1] = [2][]int{d[i][0], append([]int{l - 1}, d[i][1]...)}
	}
	return r
}

func readAll(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(file))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
