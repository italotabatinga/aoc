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
	problem := args[len(args)-1]
	fmt.Printf("Advent of Code 2020 - %v\n", problem)

	switch problem {
	case "1.1":
		aoc.Run1(src.Part1, isTest)
	case "1.2":
		aoc.Run1(src.Part2, isTest)
	default:
		fmt.Printf("Problem %v not found\n", problem)
	}
}
