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
	default:
		fmt.Printf("Problem %v not found\n", problem)
	}
}
