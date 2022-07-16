package aoc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/italotabatinga/aoc/2020/src/collections"
)

type Input2 []Test

type Runner2 struct{}

func (r Runner2) FmtInput(input string) Input2 {
	lines := strings.Split(input, "\n")
	var result Input2
	reg := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`) // 2-9 c: ccccccccc
	for _, s := range lines {
		groups := reg.FindStringSubmatch(s)
		minString, maxString, letter, password := groups[1], groups[2], groups[3], groups[4]
		min, err := strconv.Atoi(minString)
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(maxString)
		if err != nil {
			panic(err)
		}

		result = append(result, Test{
			constraint: Constraint{
				limits: collections.Tuple[int, int]{First: min, Second: max},
				letter: []rune(letter)[0],
			},
			password: password,
		})
	}

	return result
}

func (r Runner2) Run1(input Input2) {
	validCount := 0
	for _, test := range input {
		if test.isValidRange() {
			validCount++
		}
	}
	fmt.Printf("%v\n", validCount)
}

func (r Runner2) Run2(input Input2) {
	validCount := 0
	for _, test := range input {
		if test.isValidPosition() {
			validCount++
		}
	}
	fmt.Printf("%v\n", validCount)
}

type Test struct {
	constraint Constraint
	password   string
}

type Constraint struct {
	limits collections.Tuple[int, int]
	letter rune
}

func (t Test) isValidRange() bool {
	letterCount := 0
	for _, char := range t.password {
		if char == t.constraint.letter {
			letterCount++

			if letterCount > t.constraint.limits.Second {
				break
			}
		}
	}
	return letterCount <= t.constraint.limits.Second && letterCount >= t.constraint.limits.First
}

func (t Test) isValidPosition() bool {
	runes := []rune(t.password)
	valids := 0
	for _, pos := range []int{t.constraint.limits.First, t.constraint.limits.Second} {
		if val, err := safeGet(runes, pos-1); err == nil {
			if val == t.constraint.letter {
				valids++
			}
		}
	}
	return valids == 1
}

func safeGet[T any](s []T, pos int) (T, error) {
	var res T
	if pos < 0 || (pos+1) > len(s) {
		return res, fmt.Errorf("out of range(%v): %v", len(s), pos)
	}

	return s[pos], nil
}
