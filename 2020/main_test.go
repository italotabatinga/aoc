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

func TestAoc31Test(t *testing.T) {
	runTest(
		Test{problem: src.Problem{Day: 3, Part: src.Part1, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input3](p, aoc.Runner3{}) }, expected: 7},
		t,
	)
}
func TestAoc31(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 3, Part: src.Part1, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input3](p, aoc.Runner3{}) }, expected: 184}, t)
}

func TestAoc32Test(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 3, Part: src.Part2, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input3](p, aoc.Runner3{}) }, expected: 336}, t)
}
func TestAoc32(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 3, Part: src.Part2, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input3](p, aoc.Runner3{}) }, expected: 2431272960}, t)
}

func TestAoc41Test(t *testing.T) {
	runTest(
		Test{problem: src.Problem{Day: 4, Part: src.Part1, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input4](p, aoc.Runner4{}) }, expected: 2},
		t,
	)
}
func TestAoc41(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 4, Part: src.Part1, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input4](p, aoc.Runner4{}) }, expected: 202}, t)
}

func TestAoc42Test(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 4, Part: src.Part2, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input4](p, aoc.Runner4{}) }, expected: 4}, t)
}
func TestAoc42(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 4, Part: src.Part2, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input4](p, aoc.Runner4{}) }, expected: 137}, t)
}

func TestAoc61Test(t *testing.T) {
	runTest(
		Test{problem: src.Problem{Day: 6, Part: src.Part1, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input6](p, aoc.Runner6{}) }, expected: 11},
		t,
	)
}
func TestAoc61(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 6, Part: src.Part1, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input6](p, aoc.Runner6{}) }, expected: 6161}, t)
}

func TestAoc62Test(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 6, Part: src.Part2, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input6](p, aoc.Runner6{}) }, expected: 6}, t)
}
func TestAoc62(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 6, Part: src.Part2, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input6](p, aoc.Runner6{}) }, expected: 2971}, t)
}

func TestAoc71Test(t *testing.T) {
	runTest(
		Test{problem: src.Problem{Day: 7, Part: src.Part1, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input7](p, aoc.Runner7{}) }, expected: 4},
		t,
	)
}
func TestAoc71(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 7, Part: src.Part1, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input7](p, aoc.Runner7{}) }, expected: 208}, t)
}

func TestAoc72Test(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 7, Part: src.Part2, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input7](p, aoc.Runner7{}) }, expected: 126}, t)
}
func TestAoc72(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 7, Part: src.Part2, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input7](p, aoc.Runner7{}) }, expected: 1664}, t)
}

func TestAoc81Test(t *testing.T) {
	runTest(
		Test{problem: src.Problem{Day: 8, Part: src.Part1, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input8](p, aoc.Runner8{}) }, expected: 5},
		t,
	)
}
func TestAoc81(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 8, Part: src.Part1, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input8](p, aoc.Runner8{}) }, expected: 1614}, t)
}

func TestAoc82Test(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 8, Part: src.Part2, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input8](p, aoc.Runner8{}) }, expected: 8}, t)
}
func TestAoc82(t *testing.T) {
	runTest(Test{problem: src.Problem{Day: 8, Part: src.Part2, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input8](p, aoc.Runner8{}) }, expected: 1260}, t)
}

func runTest(test Test, t *testing.T) {
	got :=
		test.run(test.problem)

	if got != test.expected {
		t.Errorf("Problem %v - got: %v; expected: %v", test.problem, got, test.expected)
	}
}
