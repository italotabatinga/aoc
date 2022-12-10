package aoc

import (
	"strings"
)

type Input6 []FormGroup

type Runner6 struct{}

func (r Runner6) FmtInput(input string) Input6 {
	groups := strings.Split(input, "\n\n")
	var result Input6
	count := len(groups)
	result = make(Input6, count)
	for i, group := range groups {
		entries := strings.Split(group, "\n")
		chosen := [26]int{}
		for _, entry := range entries {
			for _, r := range entry {
				chosen[int(r-'a')]++
			}
		}
		result[i] = FormGroup{chosen: chosen, group: entries}
	}

	return result
}

func (r Runner6) Run1(input Input6, _ bool) int {
	sum := 0
	for _, group := range input {
		sum += group.anyYesCount()
	}
	return sum
}

func (r Runner6) Run2(input Input6, _ bool) int {
	sum := 0
	for _, group := range input {
		sum += group.everyYesCount()
	}
	return sum
}

type FormGroup struct {
	group  []string
	chosen [26]int
}

func (f FormGroup) anyYesCount() int {
	sum := 0
	for _, yes := range f.chosen {
		if yes > 0 {
			sum++
		}
	}
	return sum
}

func (f FormGroup) everyYesCount() int {
	sum := 0
	for _, yes := range f.chosen {
		if yes == len(f.group) {
			sum++
		}
	}
	return sum
}
