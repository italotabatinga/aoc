package aoc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Input7 map[string]map[string]int

type Runner7 struct{}

func (r Runner7) FmtInput(input string) Input7 {
	lines := strings.Split(input, "\n")
	result := make(Input7)
	reg := regexp.MustCompile(`(\d)* ?\w+ \w+ bag`) // light red bags contain 1 bright white bag, 2 muted yellow bags.
	for _, line := range lines {
		groups := reg.FindAllString(line, 20)
		if groups == nil {
			panic(fmt.Errorf("match not found: %v", line))
		}
		parent := strings.Join(strings.Fields(groups[0])[:2], " ")
		children := make(map[string]int)
		for _, node := range groups[1:] {
			if node[:13] == " no other bag" {
				break
			}
			fields := strings.Fields(node)
			bag := strings.Join(fields[1:3], " ")
			number, err := strconv.Atoi(fields[0])
			if err != nil {
				panic(fmt.Errorf("cannot atoi: %v", fields[0]))
			}
			children[bag] = number
		}
		result[parent] = children
	}

	return result
}

func (r Runner7) Run1(bags Input7, _ bool) int {
	count := 0
	goalBag := "shiny gold"
	for bag := range bags {
		if bag != goalBag && fitsInBag(bags, bag, goalBag) {
			count++
		}
	}
	return count
}

func (r Runner7) Run2(bags Input7, _ bool) int {
	count := countInBag(bags, "shiny gold")
	return count
}

func fitsInBag(bags Input7, searchBag string, bag string) bool {
	for contentBag := range bags[searchBag] {
		if contentBag == bag {
			return true
		}
	}
	for contentBag := range bags[searchBag] {
		if fitsInBag(bags, contentBag, bag) {
			return true
		}
	}
	return false
}

func countInBag(bags Input7, bag string) int {
	count := 0
	for contentBag, value := range bags[bag] {
		res := value + value*countInBag(bags, contentBag)
		count += res
	}
	return count
}
