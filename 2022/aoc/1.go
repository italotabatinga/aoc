package aoc

import (
	"strconv"
	"strings"
)

type Input1 []Elf

type Runner1 struct{}

func (r Runner1) FmtInput(input string) Input1 {
	elvesGroupString := strings.Split(input, "\n\n")
	result := make(Input1, 0)
	for _, elfString := range elvesGroupString {
		foodsString := strings.Split(elfString, "\n")
		foods := make([]int, 0)
		totalCalories := 0
		for _, foodString := range foodsString {
			food, err := strconv.Atoi(foodString)
			if err != nil {
				panic(err)
			}
			totalCalories += food
			foods = append(foods, food)
		}

		result = append(result, Elf{foods, totalCalories})
	}
	return result
}

func (r Runner1) Run1(input Input1, _ bool) int {
	if len(input) == 0 {
		return 0
	}
	max := input[0].totalCalories
	for i := 1; i < len(input); i++ {
		elf := input[i]

		totalCalories := elf.totalCalories
		if totalCalories > max {
			max = totalCalories
		}
	}

	return max
}

func (r Runner1) Run2(input Input1, _ bool) int {
	Sort(input, func(e Elf) int { return e.totalCalories })

	topThree := 0
	for i := 0; i < len(input) && i < 3; i++ {
		index := len(input) - 1 - i
		topThree += input[index].totalCalories
	}

	return topThree
}

type Elf struct {
	foods         []int
	totalCalories int
}
