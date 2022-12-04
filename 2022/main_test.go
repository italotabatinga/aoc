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

func TestAoc21Test(t *testing.T) {
	runTest(
		Test{problem: Problem{Day: 2, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input2](p, aoc.Runner2{}) }, expected: 15},
		t,
	)
}

func TestAoc21(t *testing.T) {
	runTest(Test{problem: Problem{Day: 2, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input2](p, aoc.Runner2{}) }, expected: 11841}, t)
}

func TestAoc22Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 2, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input2](p, aoc.Runner2{}) }, expected: 12}, t)
}

func TestAoc22(t *testing.T) {
	runTest(Test{problem: Problem{Day: 2, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input2](p, aoc.Runner2{}) }, expected: 13022}, t)
}

func TestAoc31Test(t *testing.T) {
	runTest(
		Test{problem: Problem{Day: 3, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input3](p, aoc.Runner3{}) }, expected: 157},
		t,
	)
}

func TestAoc31(t *testing.T) {
	runTest(Test{problem: Problem{Day: 3, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input3](p, aoc.Runner3{}) }, expected: 8109}, t)
}

func TestAoc32Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 3, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input3](p, aoc.Runner3{}) }, expected: 70}, t)
}

func TestAoc32(t *testing.T) {
	runTest(Test{problem: Problem{Day: 3, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input3](p, aoc.Runner3{}) }, expected: 2738}, t)
}
