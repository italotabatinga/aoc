package main

import (
	"testing"
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

// func TestAoc11Test(t *testing.T) {
// 	runTest(
// 		Test{problem: src.Problem{Day: 1, Part: src.Part1, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 514579},
// 		t,
// 	)
// }
// func TestAoc11(t *testing.T) {
// 	runTest(Test{problem: src.Problem{Day: 1, Part: src.Part1, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 211899}, t)
// }

// func TestAoc12Test(t *testing.T) {
// 	runTest(Test{problem: src.Problem{Day: 1, Part: src.Part2, Test: true}, run: func(p src.Problem) int { return src.Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 241861950}, t)
// }
// func TestAoc12(t *testing.T) {
// 	runTest(Test{problem: src.Problem{Day: 1, Part: src.Part2, Test: false}, run: func(p src.Problem) int { return src.Run[aoc.Input1](p, aoc.Runner1{}) }, expected: 275765682}, t)
// }
