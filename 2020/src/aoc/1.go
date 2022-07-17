package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/italotabatinga/aoc/2020/src/collections"
)

type Input1 []int

type Runner1 struct{}

func (r Runner1) FmtInput(input string) Input1 {
	inputSlice := strings.Split(input, "\n")
	var result Input1
	for _, s := range inputSlice {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		result = append(result, v)
	}
	return result
}

func (r Runner1) Run1(input Input1) int {
	set := make(collections.Set[int])
	for _, v := range input {
		comp := 2020 - v
		if set.Contains(comp) {
			result := comp * v
			fmt.Printf("%v\n", result)
			return result
		}
		set.Add(v)
	}
	return 0
}

func (r Runner1) Run2(input Input1) int {
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
			return result
		}
		for val := range values {
			sums[val+v] = collections.Tuple[int, int]{First: val, Second: v}
		}
		values.Add(v)
	}
	return 0
}
