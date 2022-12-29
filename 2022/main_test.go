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

func TestAoc41Test(t *testing.T) {
	runTest(
		Test{problem: Problem{Day: 4, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input4](p, aoc.Runner4{}) }, expected: 2},
		t,
	)
}

func TestAoc41(t *testing.T) {
	runTest(Test{problem: Problem{Day: 4, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input4](p, aoc.Runner4{}) }, expected: 528}, t)
}

func TestAoc42Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 4, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input4](p, aoc.Runner4{}) }, expected: 4}, t)
}

func TestAoc42(t *testing.T) {
	runTest(Test{problem: Problem{Day: 4, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input4](p, aoc.Runner4{}) }, expected: 881}, t)
}

func TestAoc51Test(t *testing.T) {
	runTest(
		Test{problem: Problem{Day: 5, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input5](p, aoc.Runner5{}) }, expected: 2272061909},
		t,
	)
}

func TestAoc51(t *testing.T) {
	runTest(Test{problem: Problem{Day: 5, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input5](p, aoc.Runner5{}) }, expected: 2680388803}, t)
}

func TestAoc52Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 5, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input5](p, aoc.Runner5{}) }, expected: 73263583}, t)
}

func TestAoc52(t *testing.T) {
	runTest(Test{problem: Problem{Day: 5, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input5](p, aoc.Runner5{}) }, expected: 3097473586}, t)
}

func TestAoc61Test(t *testing.T) {
	runTest(
		Test{problem: Problem{Day: 6, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input6](p, aoc.Runner6{}) }, expected: 11},
		t,
	)
}

func TestAoc61(t *testing.T) {
	runTest(Test{problem: Problem{Day: 6, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input6](p, aoc.Runner6{}) }, expected: 1343}, t)
}

func TestAoc62Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 6, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input6](p, aoc.Runner6{}) }, expected: 29}, t)
}

func TestAoc62(t *testing.T) {
	runTest(Test{problem: Problem{Day: 6, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input6](p, aoc.Runner6{}) }, expected: 2193}, t)
}

func TestAoc71Test(t *testing.T) {
	runTest(
		Test{problem: Problem{Day: 7, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input7](p, aoc.Runner7{}) }, expected: 95437},
		t,
	)
}

func TestAoc71(t *testing.T) {
	runTest(Test{problem: Problem{Day: 7, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input7](p, aoc.Runner7{}) }, expected: 1315285}, t)
}

func TestAoc72Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 7, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input7](p, aoc.Runner7{}) }, expected: 24933642}, t)
}

func TestAoc72(t *testing.T) {
	runTest(Test{problem: Problem{Day: 7, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input7](p, aoc.Runner7{}) }, expected: 9847279}, t)
}

func TestAoc81Test(t *testing.T) {
	runTest(
		Test{problem: Problem{Day: 8, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input8](p, aoc.Runner8{}) }, expected: 21},
		t,
	)
}

func TestAoc81(t *testing.T) {
	runTest(Test{problem: Problem{Day: 8, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input8](p, aoc.Runner8{}) }, expected: 1679}, t)
}

func TestAoc82Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 8, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input8](p, aoc.Runner8{}) }, expected: 8}, t)
}

func TestAoc82(t *testing.T) {
	runTest(Test{problem: Problem{Day: 8, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input8](p, aoc.Runner8{}) }, expected: 536625}, t)
}

func TestAoc91Test(t *testing.T) {
	runTest(
		Test{problem: Problem{Day: 9, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input9](p, aoc.Runner9{}) }, expected: 13},
		t,
	)
}

func TestAoc91(t *testing.T) {
	runTest(Test{problem: Problem{Day: 9, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input9](p, aoc.Runner9{}) }, expected: 5878}, t)
}

func TestAoc92Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 9, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input9](p, aoc.Runner9{}) }, expected: 36}, t)
}

func TestAoc92(t *testing.T) {
	runTest(Test{problem: Problem{Day: 9, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input9](p, aoc.Runner9{}) }, expected: 2405}, t)
}

func TestAoc101Test(t *testing.T) {
	runTest(
		Test{problem: Problem{Day: 10, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input10](p, aoc.Runner10{}) }, expected: 13140},
		t,
	)
}

func TestAoc101(t *testing.T) {
	runTest(Test{problem: Problem{Day: 10, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input10](p, aoc.Runner10{}) }, expected: 17020}, t)
}

func TestAoc102Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 10, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input10](p, aoc.Runner10{}) }, expected: 2151778615}, t)
}

func TestAoc102(t *testing.T) {
	runTest(Test{problem: Problem{Day: 10, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input10](p, aoc.Runner10{}) }, expected: 3012784333}, t)
}

func TestAoc111Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 11, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input11](p, aoc.Runner11{}) }, expected: 10605}, t)
}

func TestAoc111(t *testing.T) {
	runTest(Test{problem: Problem{Day: 11, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input11](p, aoc.Runner11{}) }, expected: 61005}, t)
}

func TestAoc112Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 11, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input11](p, aoc.Runner11{}) }, expected: 2713310158}, t)
}

func TestAoc112(t *testing.T) {
	runTest(Test{problem: Problem{Day: 11, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input11](p, aoc.Runner11{}) }, expected: 20567144694}, t)
}

func TestAoc121Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 12, Part: Part1, Test: true}, run: func(p Problem) int { return Run[aoc.Input12](p, aoc.Runner12{}) }, expected: 31}, t)
}

func TestAoc121(t *testing.T) {
	runTest(Test{problem: Problem{Day: 12, Part: Part1, Test: false}, run: func(p Problem) int { return Run[aoc.Input12](p, aoc.Runner12{}) }, expected: 425}, t)
}

func TestAoc122Test(t *testing.T) {
	runTest(Test{problem: Problem{Day: 12, Part: Part2, Test: true}, run: func(p Problem) int { return Run[aoc.Input12](p, aoc.Runner12{}) }, expected: 29}, t)
}

func TestAoc122(t *testing.T) {
	runTest(Test{problem: Problem{Day: 12, Part: Part2, Test: false}, run: func(p Problem) int { return Run[aoc.Input12](p, aoc.Runner12{}) }, expected: 418}, t)
}
