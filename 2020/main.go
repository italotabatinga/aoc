package main

import (
	"fmt"
	"os"

	"github.com/italotabatinga/aoc/2020/src"
	"github.com/italotabatinga/aoc/2020/src/aoc"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("Args must not be empty")
	}
	isTest := src.Contains(args, "--test")
	problemString := args[len(args)-1]
	fmt.Printf("Advent of Code 2020 - %v\n", problemString)

	problem := src.ParseProblem(problemString, isTest)
	switch problem.Day {
	case 1:
		src.Run[aoc.Input1](problem, aoc.Runner1{})
	default:
		fmt.Printf("Problem %v not found\n", problem)
	}
}
