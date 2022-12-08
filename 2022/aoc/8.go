package aoc

import (
	"strings"

	c "github.com/italotabatinga/aoc/2022/aoc/structures"
)

type Input8 [][]int

type Runner8 struct{}

func (r Runner8) FmtInput(input string) Input8 {
	lines := strings.Split(input, "\n")
	result := make([][]int, len(lines))
	for i, line := range lines {
		result[i] = make([]int, len(line))
		for j, r := range line {
			result[i][j] = int(r - '0')
		}
	}
	return result
}

func (r Runner8) Run1(input Input8, _ bool) int {
	count := 0
	for i, row := range input {
		for j, _ := range row {
			if input.treeIsVisible(i, j) {
				count++
			}
		}
	}

	return count
}

func (r Runner8) Run2(input Input8, _ bool) int {
	highestScenicScore := 0
	for i, row := range input {
		for j := range row {
			if scenicScore := input.ScenicScore(i, j); scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	return highestScenicScore
}

func (t Input8) treeIsVisible(i int, j int) bool {
	if t.isEdgePosition(i, j) {
		return true
	}

	visibleDirections := 4
	t.iterateDirections(i, j, func(x int, y int, _ c.Pair[int]) bool {
		if t[x][y] >= t[i][j] {
			visibleDirections--
			return true
		}

		return false
	})

	return visibleDirections > 0
}

func (t Input8) ScenicScore(i int, j int) int {
	if t.isEdgePosition(i, j) {
		return 0
	}

	viewingDistances := make(map[c.Pair[int]]int, 0)
	t.iterateDirections(i, j, func(x int, y int, dir c.Pair[int]) bool {
		if _, ok := viewingDistances[dir]; !ok {
			viewingDistances[dir] = 0
		}
		viewingDistances[dir]++

		return t[x][y] >= t[i][j]
	})

	result := 1
	for _, viewingDistance := range viewingDistances {
		result *= viewingDistance
	}

	return result
}

func (t Input8) isEdgePosition(i int, j int) bool {
	return i == 0 || j == 0 || i == len(t)-1 || j == len(t[0])-1
}

func (t Input8) iterateDirections(i int, j int, callback func(i int, j int, dir c.Pair[int]) bool) {
	directions := []c.Pair[int]{{First: 0, Second: 1}, {First: 0, Second: -1}, {First: 1, Second: 0}, {First: -1, Second: 0}}
	MAX_I := len(t) - 1
	MAX_J := len(t[0]) - 1
	for _, dir := range directions {
		x, y := i+dir.First, j+dir.Second

		for x >= 0 && y >= 0 && x <= MAX_I && y <= MAX_J {
			if callback(x, y, dir) {
				break
			}

			x += dir.First
			y += dir.Second
		}
	}
}
