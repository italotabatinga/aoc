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
	case 2:
		src.Run[aoc.Input2](problem, aoc.Runner2{})
	case 3:
		src.Run[aoc.Input3](problem, aoc.Runner3{})
	case 4:
		src.Run[aoc.Input4](problem, aoc.Runner4{})
	case 5:
		src.Run[aoc.Input5](problem, aoc.Runner5{})
	case 6:
		src.Run[aoc.Input6](problem, aoc.Runner6{})
	case 7:
		src.Run[aoc.Input7](problem, aoc.Runner7{})
	case 8:
		src.Run[aoc.Input8](problem, aoc.Runner8{})
	case 9:
		src.Run[aoc.Input9](problem, aoc.Runner9{})
	case 10:
		src.Run[aoc.Input10](problem, aoc.Runner10{})
	case 11:
		src.Run[aoc.Input11](problem, aoc.Runner11{})
	default:
		fmt.Printf("Problem %v not found\n", problem)
	}
}
