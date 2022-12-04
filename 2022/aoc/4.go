package aoc

import (
	"strconv"
	"strings"
)

type Input4 []assignment

type Runner4 struct{}

func (r Runner4) FmtInput(input string) Input4 {
	assignmentsStr := strings.Split(input, "\n")
	result := make(Input4, 0)
	for _, assignmentString := range assignmentsStr {
		elfsStr := strings.Split(assignmentString, ",")
		elf1Str := strings.Split(elfsStr[0], "-")
		elf2Str := strings.Split(elfsStr[1], "-")

		l, _ := strconv.Atoi(elf1Str[0])
		r, _ := strconv.Atoi(elf1Str[1])
		elf1 := sectionPair{l, r}

		l, _ = strconv.Atoi(elf2Str[0])
		r, _ = strconv.Atoi(elf2Str[1])
		elf2 := sectionPair{l, r}

		result = append(result, assignment{elf1, elf2})
	}

	return result
}

func (r Runner4) Run1(input Input4, _ bool) int {
	sum := 0

	for _, assignment := range input {
		if assignment.elf1.contains(assignment.elf2) || assignment.elf2.contains(assignment.elf1) {
			sum += 1
		}
	}

	return sum
}

func (r Runner4) Run2(input Input4, _ bool) int {
	sum := 0

	for _, assignment := range input {
		if assignment.elf1.overlap(assignment.elf2) || assignment.elf2.overlap(assignment.elf1) {
			sum += 1
		}
	}

	return sum
}

type assignment struct {
	elf1, elf2 sectionPair
}

type sectionPair struct {
	l, r int
}

func (s sectionPair) contains(other sectionPair) bool {
	return s.l <= other.l && s.r >= other.r
}

func (s sectionPair) overlap(other sectionPair) bool {
	return (s.l <= other.l && s.r >= other.l) ||
		(s.l <= other.r && s.r >= other.r)
}
