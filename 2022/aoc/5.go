package aoc

import (
	"fmt"
	"strconv"
	"strings"

	c "github.com/italotabatinga/aoc/2022/aoc/collections"
)

type Input5 struct {
	stacks    []c.Stack[rune]
	movements []CrateMovement
}

type Runner5 struct{}

func (r Runner5) FmtInput(input string) Input5 {
	inputString := strings.Split(input, "\n\n")
	stacksStrings := strings.Split(inputString[0], "\n")

	result := Input5{}
	result.stacks = make([]c.Stack[rune], 0)
	numStacks := (len(stacksStrings[0]) + 1) / 4

	for i := 0; i < numStacks; i++ {
		column := i*4 + 1
		stack := c.NewStack[rune]()

		for j := len(stacksStrings) - 2; j >= 0; j-- {
			line := stacksStrings[j]
			r := []rune(line)[column]

			if r != ' ' {
				stack.Push(r)
			} else {
				break
			}
		}
		result.stacks = append(result.stacks, stack)
	}

	movementsString := strings.Split(inputString[1], "\n")
	result.movements = make([]CrateMovement, 0)
	for _, line := range movementsString {
		elems := strings.Split(line, " ")
		count, _ := strconv.Atoi(elems[1])
		src, _ := strconv.Atoi(elems[3])
		dst, _ := strconv.Atoi(elems[5])
		result.movements = append(result.movements, CrateMovement{count, src, dst})
	}

	return result
}

func (r Runner5) Run1(input Input5, _ bool) int {
	for _, mov := range input.movements {
		src := &input.stacks[mov.src-1]
		dst := &input.stacks[mov.dst-1]

		for i := 0; i < mov.count && src.Len() > 0; i++ {
			dst.Push(src.Pop())
		}
	}

	var sb strings.Builder

	for _, stack := range input.stacks {
		if stack.Len() > 0 {
			sb.WriteRune(stack.Top())
		}
	}

	return StringToInt(sb.String())
}

func (r Runner5) Run2(input Input5, _ bool) int {
	for _, mov := range input.movements {
		src := &input.stacks[mov.src-1]
		dst := &input.stacks[mov.dst-1]

		dst.MultPush(src.MultPop(mov.count)...)
	}

	var sb strings.Builder

	for _, stack := range input.stacks {
		if stack.Len() > 0 {
			sb.WriteRune(stack.Top())
		}
	}

	return StringToInt(sb.String())
}

type CrateMovement struct {
	count, src, dst int
}

func (i Input5) String() string {
	return fmt.Sprintf("Input5(\n%v\n%v\n)", i.stacks, i.movements)
}
