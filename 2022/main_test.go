package main

import (
	"testing"

	"github.com/italotabatinga/aoc/2022/aoc"
)

type Test struct {
	problem  Problem
	run      func(p Problem) int
	expected int
}

func runTest(test Test, t *testing.T) {
	got :=
		test.run(test.problem)

	if got != test.expected {
		t.Errorf("Problem %v - got: %v; expected: %v", test.problem, got, test.expected)
	}
}

func TestAoc11Test(t *testing.T) {
	runTest(
		Test{problem: Problem{Day: 1, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 24000},
		t,
	)
}

func TestAoc11(t *testing.T) {
	runTest(Test{problem: Problem{Day: 1, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 70296}, t)
}

func TestAoc12Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 1, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 45000}, t)
}

func TestAoc12(t *testing.T) {
	runTest(Test{problem: Problem{Day: 1, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 205381}, t)
}
