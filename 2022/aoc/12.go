package aoc

import (
	"math"
	"strings"

	c "github.com/italotabatinga/aoc/2022/aoc/collections"
)

type Input12 SurroundingArea

type Runner12 struct{}

func (r Runner12) FmtInput(input string) Input12 {
	lines := strings.Split(input, "\n")
	result := SurroundingArea{}
	result.hmap = make([][]rune, len(lines))
	result.srcs = []c.Pair[int]{}
	result.size = c.Pair[int]{First: len(lines), Second: len(lines[0])}
	for i := range lines {
		result.hmap[i] = []rune(lines[i])

	}

	for i, line := range result.hmap {
		for j, r := range line {
			if r == 'S' {
				result.hmap[i][j] = 'a'
				result.src = c.Pair[int]{First: i, Second: j}
				result.srcs = append(result.srcs, c.Pair[int]{First: i, Second: j})
			} else if r == 'E' {
				result.hmap[i][j] = 'z'
				result.dst = c.Pair[int]{First: i, Second: j}
			} else if r == 'a' {
				result.srcs = append(result.srcs, c.Pair[int]{First: i, Second: j})
			}
		}
	}

	return Input12(result)
}

func (r Runner12) Run1(input Input12, _ bool) int {
	area := SurroundingArea(input)
	cost := area.newCost(area.src.First, area.src.Second)
	area.calculateCosts(area.src.First, area.src.Second, &cost)

	return cost[area.dst.First][area.dst.Second]
}

func (r Runner12) Run2(input Input12, _ bool) int {
	area := SurroundingArea(input)
	costs := []int{}
	for _, src := range area.srcs {
		cost := area.newCost(src.First, src.Second)
		area.calculateCosts(src.First, src.Second, &cost)
		costs = append(costs, cost[area.dst.First][area.dst.Second])
	}

	return Min(costs...)
}

type SurroundingArea struct {
	size c.Pair[int]
	hmap [][]rune
	src  c.Pair[int]
	srcs []c.Pair[int]
	dst  c.Pair[int]
}

func (s SurroundingArea) possibleNextSteps(x int, y int) []c.Pair[int] {
	curr := s.hmap[x][y]
	possibilities := []c.Pair[int]{}
	if x > 0 {
		if next := s.hmap[x-1][y]; int(next-curr) <= 1 {
			possibilities = append(possibilities, c.Pair[int]{First: x - 1, Second: y})
		}
	}
	if x < s.size.First-1 {
		if next := s.hmap[x+1][y]; int(next-curr) <= 1 {
			possibilities = append(possibilities, c.Pair[int]{First: x + 1, Second: y})
		}
	}
	if y > 0 {
		if next := s.hmap[x][y-1]; int(next-curr) <= 1 {
			possibilities = append(possibilities, c.Pair[int]{First: x, Second: y - 1})
		}
	}
	if y < s.size.Second-1 {
		if next := s.hmap[x][y+1]; int(next-curr) <= 1 {
			possibilities = append(possibilities, c.Pair[int]{First: x, Second: y + 1})
		}
	}

	return possibilities
}

func (s SurroundingArea) calculateCosts(x int, y int, cost *[][]int) {
	possibilitites := s.possibleNextSteps(x, y)
	for _, possibility := range possibilitites {
		currCost := (*cost)[x][y]
		currPossibilityCost := (*cost)[possibility.First][possibility.Second]
		newPossibilityCost := currCost + 1
		if newPossibilityCost < currPossibilityCost {
			(*cost)[possibility.First][possibility.Second] = newPossibilityCost
			s.calculateCosts(possibility.First, possibility.Second, cost)
		}
	}
}

func (s SurroundingArea) newCost(initX int, initY int) [][]int {
	cost := make([][]int, s.size.First)
	for i := range cost {
		cost[i] = make([]int, s.size.Second)

		for j := 0; j < s.size.Second; j++ {
			if i == initX && j == initY {
				cost[i][j] = 0
			} else {
				cost[i][j] = math.MaxInt
			}
		}
	}

	return cost
}
