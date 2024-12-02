package aoc

Input :: struct {
	value: string,
}

ProblemPart :: enum {
	One,
	Two,
}

Problem :: struct {
	day:  int,
	part: ProblemPart,
	test: bool,
}

ProblemSolver :: struct {
	day:  int,
	part: ProblemPart,
	test: bool,
}
