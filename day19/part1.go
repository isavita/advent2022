package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("day19/input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.TrimSpace(string(file))

	fmt.Println(part1(input))
}

func part1(input string) int {
	blueprints := parseInput(input)

	sum := 0
	for _, bp := range blueprints {
		st := newState(bp)
		geodesMade := st.calcMostGeodes(0, map[string]int{}, 24, 24)

		sum += st.blueprint.id * geodesMade
	}

	return sum
}

type blueprint struct {
	id                                        int
	oreForOreRobot                            int
	oreForClayRobot                           int
	oreForObsidianRobot, clayForObsidianRobot int
	oreForGeodeRobot, obsidianForGeodeRobot   int
}

type state struct {
	blueprint
	ore, clay, obsidian, geode                         int
	oreRobots, clayRobots, obsidianRobots, geodeRobots int
}

func newState(blueprint blueprint) state {
	return state{
		blueprint: blueprint,
		oreRobots: 1,
	}
}

func (s *state) farm() {
	s.ore += s.oreRobots
	s.clay += s.clayRobots
	s.obsidian += s.obsidianRobots
	s.geode += s.geodeRobots
}

func (s *state) hash(time int) string {
	return fmt.Sprint(time, s.ore, s.clay, s.obsidian,
		s.geode, s.oreRobots, s.clayRobots, s.obsidianRobots, s.geodeRobots)
}

func (s state) copy() state {
	return s
}

func (s *state) calcMostGeodes(time int, memo map[string]int, totalTime int, earliestGeode int) int {
	if time == totalTime {
		return s.geode
	}

	h := s.hash(time)
	if v, ok := memo[h]; ok {
		return v
	}

	if s.geode == 0 && time > earliestGeode {
		return 0
	}

	mostGeodes := s.geode

	if s.ore >= s.oreForGeodeRobot &&
		s.obsidian >= s.obsidianForGeodeRobot {
		cp := s.copy()

		cp.farm()

		cp.ore -= cp.oreForGeodeRobot
		cp.obsidian -= cp.obsidianForGeodeRobot
		cp.geodeRobots++
		if cp.geodeRobots == 1 {
			earliestGeode = minInt(earliestGeode, time+1)
		}
		mostGeodes = maxInt(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))

		memo[h] = mostGeodes
		return mostGeodes
	}

	if time <= totalTime-16 &&
		s.oreRobots < s.oreForObsidianRobot*2 &&
		s.ore >= s.oreForOreRobot {
		cp := s.copy()
		cp.ore -= cp.oreForOreRobot

		cp.farm()

		cp.oreRobots++
		mostGeodes = maxInt(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))
	}
	if time <= totalTime-8 &&
		s.clayRobots < s.clayForObsidianRobot &&
		s.ore >= s.oreForClayRobot {
		cp := s.copy()
		cp.ore -= cp.oreForClayRobot

		cp.farm()

		cp.clayRobots++
		mostGeodes = maxInt(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))
	}
	if time <= totalTime-4 &&
		s.obsidianRobots < s.obsidianForGeodeRobot &&
		s.ore >= s.oreForObsidianRobot && s.clay >= s.clayForObsidianRobot {

		cp := s.copy()
		cp.ore -= cp.oreForObsidianRobot
		cp.clay -= cp.clayForObsidianRobot
		cp.farm()

		cp.obsidianRobots++
		mostGeodes = maxInt(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))
	}

	cp := s.copy()
	cp.ore += cp.oreRobots
	cp.clay += cp.clayRobots
	cp.obsidian += cp.obsidianRobots
	cp.geode += cp.geodeRobots
	mostGeodes = maxInt(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))

	memo[h] = mostGeodes
	return mostGeodes
}

func parseInput(input string) (ans []blueprint) {

	for _, line := range strings.Split(input, "\n") {
		bp := blueprint{}
		_, err := fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&bp.id, &bp.oreForOreRobot, &bp.oreForClayRobot, &bp.oreForObsidianRobot,
			&bp.clayForObsidianRobot, &bp.oreForGeodeRobot, &bp.obsidianForGeodeRobot)
		if err != nil {
			panic("parsing: " + err.Error())
		}
		ans = append(ans, bp)
	}
	return ans
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
