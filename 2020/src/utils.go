package src

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type Part int

const (
	Part1 Part = iota
	Part2
)

type Runner[R any] interface {
	Run1(v R)
	Run2(v R)
	FmtInput(s string) R
}

func findInputFile(p Problem) string {
	specificFilePath := fmt.Sprintf("%v.%v", p.Day, p.Part)
	genericFilePath := fmt.Sprintf("%v", p.Day)
	testablePaths := []string{specificFilePath, genericFilePath}
	for _, filepath := range testablePaths {
		filepath := path.Join("inputs", filepath+".txt")
		if p.Test {
			filepath = strings.Replace(filepath, ".txt", "_test.txt", 1)
		}

		if _, err := os.Stat(filepath); err == nil {
			bytes, err := os.ReadFile(filepath)
			if err != nil {
				panic(err)
			}

			return string(bytes)
		}
	}

	panic(fmt.Sprintf("Missing input file for %v", p))
}

func Run[R any](problem Problem, runner Runner[R]) {
	stringInput := findInputFile(problem)
	input := runner.FmtInput(stringInput)

	switch problem.Part {
	case Part1:
		runner.Run1(input)
	case Part2:
		runner.Run2(input)
	}
}

type Problem struct {
	Day  int
	Part Part
	Test bool
}

func ParseProblem(s string, test bool) Problem {
	split := strings.Split(s, ".")

	day, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}

	partInt, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	var part Part
	switch partInt {
	case 1:
		part = Part1
	case 2:
		part = Part2
	default:
		panic(fmt.Sprintf("Unexpected part: %v", partInt))
	}
	return Problem{Day: day, Part: part, Test: test}
}

func Contains[V comparable](s []V, v V) bool {
	for _, vs := range s {
		if v == vs {
			return true
		}
	}
	return false
}
