package aoc

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/italotabatinga/aoc/2020/src"
	"github.com/italotabatinga/aoc/2020/src/collections"
)

func Run1(part src.Part, isTest bool) {
	filepath := path.Join("inputs", "1.txt")
	if isTest {
		filepath = strings.Replace(filepath, ".txt", "_test.txt", 1)
	}
	bytes, err := os.ReadFile(filepath)
	input := string(bytes)
	if err != nil {
		panic(err)
	}

	switch part {
	case src.Part1:
		runner_1(formatInput(input))
	case src.Part2:
		runner_2(formatInput(input))
	}
}

func formatInput(input string) []int {
	inputSlice := strings.Split(input, "\n")
	var result []int
	for _, s := range inputSlice {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		result = append(result, v)
	}
	return result
}

func runner_1(input []int) {
	set := make(collections.Set[int])
	for _, v := range input {
		comp := 2020 - v
		if set.Contains(comp) {
			result := comp * v
			fmt.Printf("%v\n", result)
			return
		}
		set.Add(v)
	}
}

func runner_2(input []int) {
	values := make(collections.Set[int])
	sums := make(map[int]collections.Tuple[int, int])
	for _, v := range input {
		if values.Contains(v) {
			continue
		}
		comp := 2020 - v
		if tup, ok := sums[comp]; ok {
			result := v * tup.First * tup.Second
			fmt.Printf("%v * %v * %v = %v\n", v, tup.First, tup.Second, result)
			return
		}
		for val := range values {
			sums[val+v] = collections.Tuple[int, int]{First: val, Second: v}
		}
		values.Add(v)
	}
}
