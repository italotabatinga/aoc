package main

import (
	"testing"

	"github.com/italotabatinga/aoc/2020/src"
	"github.com/italotabatinga/aoc/2020/src/aoc"
)

type Test struct {
	problem  src.Problem
	run      func(p src.Problem) int
	expected int
}

func TestAoc11Test(t *testing.T) {
	runTest(
		Test{problem: src.Problem{Day: 1, Part: src.Part1, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 514579},
		t,
	)
}
func TestAoc11(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 1, Part: src.Part1, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 211899}, t)
}

func TestAoc12Test(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 1, Part: src.Part2, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 241861950}, t)
}
func TestAoc12(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 1, Part: src.Part2, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 275765682}, t)
}

func TestAoc21Test(t *testing.T) {
	runTest(
		Test{problem: src.Problem{Day: 2, Part: src.Part1, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input2](p, aoc.Runner2{}) }, expected: 2},
		t,
	)
}
func TestAoc21(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 2, Part: src.Part1, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input2](p, aoc.Runner2{}) }, expected: 638}, t)
}

func TestAoc22Test(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 2, Part: src.Part2, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input2](p, aoc.Runner2{}) }, expected: 1}, t)
}
func TestAoc22(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 2, Part: src.Part2, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input2](p, aoc.Runner2{}) }, expected: 699}, t)
}

func runTest(test Test, t *testing.T) {
	got :=
		test.run(test.problem)

	if got != test.expected {
		t.Errorf("Problem %v - got: %v; expected: %v", test.problem, got, test.expected)
	}
}
