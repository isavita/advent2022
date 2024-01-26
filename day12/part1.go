package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"image"
	"os"
)

func main() {
	grid := map[image.Point]byte{}
	s := scanAll()
	var start, end image.Point
	var as []image.Point
	y := 0
	for s.Scan() {
		for x, b := range s.Bytes() {
			p := image.Pt(x, y)
			grid[p] = b
			if b == 'S' {
				start = p
			} else if b == 'E' {
				end = p
			} else if b == 'a' {
				as = append(as, p)
			}
		}
		y++
	}
	grid[start], grid[end] = 'a', 'z'

	dists := djikstra(grid, end)

	l := dists[start]
	fmt.Println(l)
}

type Item struct {
	Obj      interface{}
	Priority int
}

type PQ []*Item

func (pq PQ) Len() int            { return len(pq) }
func (pq PQ) Less(i, j int) bool  { return pq[i].Priority > pq[j].Priority }
func (pq PQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) { *pq = append(*pq, x.(*Item)) }
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

var Neighbors4 = []image.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func djikstra(grid map[image.Point]byte, end image.Point) map[image.Point]int {
	pq := PQ{&Item{end, 0}}
	dist := map[image.Point]int{end: 0}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item).Obj.(image.Point)
		for _, n := range Neighbors4 {
			next := curr.Add(n)
			if _, ok := grid[next]; !ok {
				continue
			}
			if int(grid[curr])-int(grid[next]) > 1 {
				continue
			}
			nextdist := dist[curr] + 1
			if d, ok := dist[next]; !ok || nextdist < d {
				dist[next] = nextdist
				heap.Push(&pq, &Item{next, nextdist})
			}
		}
	}
	return dist
}

func scanAll() *bufio.Scanner {
	file, err := os.Open("day12/input.txt")
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
