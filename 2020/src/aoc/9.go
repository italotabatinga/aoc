package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

type Input9 []int

type Runner9 struct{}

func (r Runner9) FmtInput(input string) Input9 {
	lines := strings.Split(input, "\n")
	result := make(Input9, len(lines))
	for i, line := range lines {
		if val, err := strconv.Atoi(line); err == nil {
			result[i] = val
		} else {
			panic(fmt.Errorf("unexpected atoi %v", line))
		}
	}

	return result
}

func (r Runner9) Run1(nums Input9, test bool) int {
	preamble := 25
	if test {
		preamble = 5
	}
	if preamble > len(nums) {
		panic(fmt.Errorf("input should be larger than %v", preamble))
	}

	for i := preamble; i < len(nums); i++ {
		slice := nums[i-preamble : i]
		if val := nums[i]; !hasTwoSum(slice, val) {
			fmt.Printf("Found %v!\n", val)
			return val
		}
	}

	return -1
}

func (r Runner9) Run2(nums Input9, test bool) int {
	preamble := 25
	if test {
		preamble = 5
	}
	if preamble > len(nums) {
		panic(fmt.Errorf("input should be larger than %v", preamble))
	}

	invalid := 0
	for i := preamble; i < len(nums); i++ {
		slice := nums[i-preamble : i]
		if val := nums[i]; !hasTwoSum(slice, val) {
			invalid = val
			break
		}
	}

	for i := 0; i < len(nums); i++ {
		list := []int{nums[i]}
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			list = append(list, nums[j])
			sum += nums[j]

			if sum == invalid {
				x := smallest(list) + largest(list)
				fmt.Printf("Found %v!\n", x)
				return x
			}
		}
	}
	return -1
}

func hasTwoSum(slice []int, sum int) bool {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i]+slice[j] == sum {
				return true
			}
		}
	}

	return false
}

func smallest(slice []int) int {
	min := slice[0]
	for _, val := range slice {
		if min < val {
			min = val
		}
	}

	return min
}
func largest(slice []int) int {
	max := slice[0]
	for _, val := range slice {
		if max > val {
			max = val
		}
	}

	return max
}
