package aoc

import (
	"math"
	"strings"

	. "github.com/italotabatinga/aoc/2020/src/collections"
)

type Input11 SeatGrid

type Runner11 struct{}

func (r Runner11) FmtInput(input string) Input11 {
	lines := strings.Split(input, "\n")
	result := make(Input11, len(lines))
	for i, line := range lines {
		runes := []rune(line)
		result[i] = runes
	}
	return result
}

func (r Runner11) Run1(input Input11, _ bool) int {
	var grid1, grid2 SeatGrid
	grid1 = SeatGrid(input)
	grid2 = make(SeatGrid, len(grid1))
	for i, runes := range grid1 {
		grid2[i] = make([]rune, len(runes))
	}
	var old, new *SeatGrid = &grid1, &grid2

	hasChanged := true

	for count := 0; hasChanged; count++ {
		hasChanged = false

		for i, line := range *old {
			for j, state := range line {
				nextState := nextSeatState(*old, i, j, SeatStateOptions{NeighboorsToLeave: 4, RangeOfSight: 1})

				(*new)[i][j] = nextState
				if nextState != state {
					hasChanged = true
				}
			}
		}

		tmp := new
		new = old
		old = tmp
	}

	return countOccupiedSeats(*old)
}

func (r Runner11) Run2(input Input11, _ bool) int {
	var grid1, grid2 SeatGrid
	grid1 = SeatGrid(input)
	grid2 = make(SeatGrid, len(grid1))
	for i, runes := range grid1 {
		grid2[i] = make([]rune, len(runes))
	}
	var old, new *SeatGrid = &grid1, &grid2

	hasChanged := true

	for count := 0; hasChanged; count++ {
		hasChanged = false

		for i, line := range *old {
			for j, state := range line {
				nextState := nextSeatState(*old, i, j, SeatStateOptions{NeighboorsToLeave: 5, RangeOfSight: math.MaxInt})

				(*new)[i][j] = nextState
				if nextState != state {
					hasChanged = true
				}
			}
		}

		tmp := new
		new = old
		old = tmp
	}

	return countOccupiedSeats(*old)
}

type SeatGrid [][]rune

type SeatStateOptions struct {
	NeighboorsToLeave int
	RangeOfSight      int
}

func nextSeatState(grid SeatGrid, i int, j int, options SeatStateOptions) rune {
	state := grid[i][j]
	if state == NONE {
		return state
	}

	directions := []Pair[int]{{First: -1, Second: -1}, {First: -1, Second: 0}, {First: -1, Second: 1}, {First: 0, Second: 1}, {First: 1, Second: 1}, {First: 1, Second: 0}, {First: 1, Second: -1}, {First: 0, Second: -1}}
	occupiedNeighboors := 0
	MAX_I := len(grid) - 1
	MAX_J := len(grid[0]) - 1
	for _, dir := range directions {
		x, y := i+dir.First, j+dir.Second

		countSteps := 0
		for x >= 0 && y >= 0 && x <= MAX_I && y <= MAX_J && countSteps < options.RangeOfSight {
			if grid[x][y] == OCCUPIED {
				occupiedNeighboors++
			}
			if grid[x][y] == OCCUPIED || grid[x][y] == EMPTY {
				break
			}

			x += dir.First
			y += dir.Second
			countSteps++
		}
	}

	if state == EMPTY && occupiedNeighboors == 0 {
		return OCCUPIED
	} else if state == OCCUPIED && occupiedNeighboors >= options.NeighboorsToLeave {
		return EMPTY
	} else {
		return state
	}
}

func countOccupiedSeats(grid SeatGrid) int {
	count := 0
	for _, line := range grid {
		for _, state := range line {
			if state == OCCUPIED {
				count += 1
			}
		}
	}

	return count
}

func (g SeatGrid) String() string {
	var sb strings.Builder

	for _, lines := range g {
		for _, state := range lines {
			sb.WriteRune(state)
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

const (
	OCCUPIED = '#'
	EMPTY    = 'L'
	NONE     = '.'
)
