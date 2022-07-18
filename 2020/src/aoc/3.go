package aoc

import (
	"fmt"
	"strings"

	c "github.com/italotabatinga/aoc/2020/src/collections"
)

type Input3 [][]bool

type Runner3 struct{}

func (r Runner3) FmtInput(input string) Input3 {
	lines := strings.Split(input, "\n")
	var result Input3
	height := len(lines)
	result = make(Input3, height)
	for i, s := range lines {
		row := &result[i]
		for _, c := range s {
			switch c {
			case '.':
				*row = append(*row, false)
			case '#':
				*row = append(*row, true)
			default:
				panic(fmt.Errorf("unexpected rune: %v", c))
			}
		}
	}

	return result
}

func (r Runner3) Run1(input Input3, _ bool) int {
	count := 0
	for y, x := 1, 3; y < len(input); y, x = y+1, x+3 {
		row := input[y]
		mod_x := x % len(row)

		if row[mod_x] {
			count++
		}
	}
	fmt.Printf("Total trees: %v\n", count)
	return count
}

func (r Runner3) Run2(input Input3, _ bool) int {
	slopes := []c.Pair[int]{{First: 1, Second: 1}, {First: 1, Second: 3}, {First: 1, Second: 5}, {First: 1, Second: 7}, {First: 2, Second: 1}}
	treesFound := make([]int, len(slopes))

	for i, slope := range slopes {
		dy, dx := slope.First, slope.Second
		for y, x := dy, dx; y < len(input); y, x = y+dy, x+dx {
			row := input[y]
			mod_x := x % len(row)

			if row[mod_x] {
				treesFound[i]++
			}
		}
	}
	result := 1
	for _, count := range treesFound {
		result *= count
	}
	fmt.Printf("Trees: %v\n", treesFound)
	fmt.Printf("Result: %v\n", result)
	return result
}
