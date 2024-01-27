package main

import (
	"bufio"
	"fmt"
	"os"
)

type blueprint struct {
	id       int
	oreCost  int
	clay     struct{ oreCost int }
	obsidian struct{ oreCost, clayCost int }
	geode    struct{ oreCost, obsidianCost int }
}

func main() {
	var blueprints []blueprint
	s := scanAll()
	for s.Scan() {
		var b blueprint
		fmt.Sscanf(s.Text(),
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&b.id, &b.oreCost, &b.clay.oreCost, &b.obsidian.oreCost, &b.obsidian.clayCost, &b.geode.oreCost, &b.geode.obsidianCost)
		blueprints = append(blueprints, b)
	}

	init := state{oreRobots: 1, timeLeft: 24}

	init.timeLeft = 32
	prod := 1
	for _, b := range blueprints[:3] {
		prod *= maxGeode(b, init)
	}
	fmt.Println(prod)
}

type state struct {
	ore, clay, obsidian, geode                         int
	oreRobots, clayRobots, obsidianRobots, geodeRobots int
	timeLeft                                           int
}

func maxGeode(b blueprint, st state) int {
	max := 0
	q := []state{st}
	visited := map[state]struct{}{}
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		max = Max(max, s.geode)
		if s.timeLeft == 0 {
			continue
		}
		o := Max(b.oreCost, b.clay.oreCost, b.obsidian.oreCost, b.geode.oreCost)
		if s.oreRobots >= o {
			s.oreRobots = o
		}
		if s.clayRobots >= b.obsidian.clayCost {
			s.clayRobots = b.obsidian.clayCost
		}
		if s.obsidianRobots >= b.geode.obsidianCost {
			s.obsidianRobots = b.geode.obsidianCost
		}
		if maxOre := s.timeLeft*o - s.oreRobots*(s.timeLeft-1); s.ore >= maxOre {
			s.ore = maxOre
		}
		if maxClay := s.timeLeft*b.obsidian.clayCost - s.clayRobots*(s.timeLeft-1); s.clay >= maxClay {
			s.clay = maxClay
		}
		if maxObsidian := s.timeLeft*b.geode.obsidianCost - s.obsidianRobots*(s.timeLeft-1); s.obsidian >= maxObsidian {
			s.obsidian = maxObsidian
		}
		if _, ok := visited[s]; ok {
			continue
		}
		visited[s] = struct{}{}

		q = append(q, state{
			ore:            s.ore + s.oreRobots,
			clay:           s.clay + s.clayRobots,
			obsidian:       s.obsidian + s.obsidianRobots,
			geode:          s.geode + s.geodeRobots,
			oreRobots:      s.oreRobots,
			clayRobots:     s.clayRobots,
			obsidianRobots: s.obsidianRobots,
			geodeRobots:    s.geodeRobots,
			timeLeft:       s.timeLeft - 1,
		})

		if s.ore >= b.oreCost {
			q = append(q, state{
				ore:            s.ore - b.oreCost + s.oreRobots,
				clay:           s.clay + s.clayRobots,
				obsidian:       s.obsidian + s.obsidianRobots,
				geode:          s.geode + s.geodeRobots,
				oreRobots:      s.oreRobots + 1,
				clayRobots:     s.clayRobots,
				obsidianRobots: s.obsidianRobots,
				geodeRobots:    s.geodeRobots,
				timeLeft:       s.timeLeft - 1,
			})
		}

		if s.ore >= b.clay.oreCost {
			q = append(q, state{
				ore:            s.ore - b.clay.oreCost + s.oreRobots,
				clay:           s.clay + s.clayRobots,
				obsidian:       s.obsidian + s.obsidianRobots,
				geode:          s.geode + s.geodeRobots,
				oreRobots:      s.oreRobots,
				clayRobots:     s.clayRobots + 1,
				obsidianRobots: s.obsidianRobots,
				geodeRobots:    s.geodeRobots,
				timeLeft:       s.timeLeft - 1,
			})
		}

		if s.ore >= b.obsidian.oreCost && s.clay >= b.obsidian.clayCost {
			q = append(q, state{
				ore:            s.ore - b.obsidian.oreCost + s.oreRobots,
				clay:           s.clay - b.obsidian.clayCost + s.clayRobots,
				obsidian:       s.obsidian + s.obsidianRobots,
				geode:          s.geode + s.geodeRobots,
				oreRobots:      s.oreRobots,
				clayRobots:     s.clayRobots,
				obsidianRobots: s.obsidianRobots + 1,
				geodeRobots:    s.geodeRobots,
				timeLeft:       s.timeLeft - 1,
			})
		}

		if s.ore >= b.geode.oreCost && s.obsidian >= b.geode.obsidianCost {
			q = append(q, state{
				ore:            s.ore - b.geode.oreCost + s.oreRobots,
				clay:           s.clay + s.clayRobots,
				obsidian:       s.obsidian - b.geode.obsidianCost + s.obsidianRobots,
				geode:          s.geode + s.geodeRobots,
				oreRobots:      s.oreRobots,
				clayRobots:     s.clayRobots,
				obsidianRobots: s.obsidianRobots,
				geodeRobots:    s.geodeRobots + 1,
				timeLeft:       s.timeLeft - 1,
			})
		}
	}
	return max
}

func scanAll() *bufio.Scanner {
	file, err := os.Open("day19/input.txt")
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}

func Max(ns ...int) int {
	max := ns[0]
	for _, n := range ns[1:] {
		if n > max {
			max = n
		}
	}
	return max
}
