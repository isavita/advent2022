package main

import "testing"

func TestDay12Task1(t *testing.T) {
	ans := Day12Task1()
	if ans != 423 {
		t.Errorf("Day12Task1() = %d; expected 423", ans)
	}
}

func BenchmarkDay12Task1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day12Task1()
	}
}

func TestDay12Task2(t *testing.T) {
	ans := Day12Task2()
	if ans != 416 {
		t.Errorf("Day12Task1() = %d; expected 423", ans)
	}
}

func BenchmarkDay12Task2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day12Task2()
	}
}
