package main

import (
	"fmt"
	"os"

	"github.com/italotabatinga/aoc/2022/aoc"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("Args must not be empty")
	}
	isTest := Contains(args, "--test")
	problemString := args[len(args)-1]
	fmt.Printf("Advent of Code 2020 - %v\n", problemString)

	problem := ParseProblem(problemString, isTest)
	switch problem.Day {
	case 1:
		Run[aoc.Input1](problem, aoc.Runner1{})
	case 2:
		Run[aoc.Input2](problem, aoc.Runner2{})
	case 3:
		Run[aoc.Input3](problem, aoc.Runner3{})
	case 4:
		Run[aoc.Input4](problem, aoc.Runner4{})
	case 5:
		Run[aoc.Input5](problem, aoc.Runner5{})
	case 6:
		Run[aoc.Input6](problem, aoc.Runner6{})
	case 7:
		Run[aoc.Input7](problem, aoc.Runner7{})
	case 8:
		Run[aoc.Input8](problem, aoc.Runner8{})
	case 9:
		Run[aoc.Input9](problem, aoc.Runner9{})
	default:
		fmt.Printf("Problem %v not found\n", problem)
	}
}
