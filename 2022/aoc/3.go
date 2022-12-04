package aoc

import (
	"fmt"
	"strings"
)

type Input3 []rucksack

type Runner3 struct{}

func (r Runner3) FmtInput(input string) Input3 {
	rucksacksString := strings.Split(input, "\n")
	result := make(Input3, 0)
	for _, rucksackString := range rucksacksString {
		chars := []rune(rucksackString)
		pivot := len(chars) / 2
		rucksack := rucksack{}
		var priority int
		for i := 0; i < pivot; i++ {
			char := chars[i]
			if char >= 97 {
				priority = int(char - 'a')
			} else {
				priority = int(char-'A') + 26
			}
			rucksack.firstCompartment[priority]++
		}
		for i := pivot; i < len(chars); i++ {
			char := chars[i]
			if char >= 97 {
				priority = int(char - 'a')
			} else {
				priority = int(char-'A') + 26
			}
			rucksack.secondCompartment[priority]++
		}
		result = append(result, rucksack)
	}

	return result
}

func (r Runner3) Run1(input Input3, _ bool) int {
	sum := 0
	for _, rucksack := range input {
		for priority := 0; priority < 52; priority++ {
			if rucksack.firstCompartment[priority] > 0 && rucksack.secondCompartment[priority] > 0 {
				sum += priority + 1
			}
		}
	}

	return sum
}

func (r Runner3) Run2(input Input3, _ bool) int {
	sum := 0

	for i := 0; i < len(input); i += 3 {
		firstRucksack := input[i]
		secondRucksack := input[i+1]
		thirdRucksack := input[i+2]
		for priority := 0; priority < 52; priority++ {
			if (firstRucksack.firstCompartment[priority] > 0 || firstRucksack.secondCompartment[priority] > 0) &&
				(secondRucksack.firstCompartment[priority] > 0 || secondRucksack.secondCompartment[priority] > 0) &&
				(thirdRucksack.firstCompartment[priority] > 0 || thirdRucksack.secondCompartment[priority] > 0) {
				sum += priority + 1
			}
		}
	}
	return sum
}

type rucksack struct {
	firstCompartment  [52]int
	secondCompartment [52]int
}

func (r rucksack) String() string {
	return fmt.Sprintf("rucksack\n  f: %v\n  s: %v\n", r.firstCompartment, r.secondCompartment)
}
