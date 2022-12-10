package aoc

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Input10 []int

type Runner10 struct{}

func (r Runner10) FmtInput(input string) Input10 {
	lines := strings.Split(input, "\n")
	result := make(Input10, len(lines))
	for i, line := range lines {
		if val, err := strconv.Atoi(line); err == nil {
			result[i] = val
		} else {
			panic(fmt.Errorf("unexpected atoi %v", line))
		}
	}

	return result
}

func (r Runner10) Run1(adapters Input10, test bool) int {
	sort.Ints(adapters)
	lastJoltage := 0
	diffs := [3]int{}
	for _, joltage := range adapters {
		if diff := joltage - lastJoltage; diff <= 3 {
			diffs[diff-1]++
		}
		lastJoltage = joltage
	}
	diffs[2]++
	result := diffs[0] * diffs[2]
	return result
}

func (r Runner10) Run2(adapters Input10, test bool) int {
	return -1
}
