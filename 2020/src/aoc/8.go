package aoc

import (
	"fmt"
	"strconv"
	"strings"

	c "github.com/italotabatinga/aoc/2020/src/collections"
)

type Input8 []c.Tuple[string, int]

type Runner8 struct{}

func (r Runner8) FmtInput(input string) Input8 {
	lines := strings.Split(input, "\n")
	result := make(Input8, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		cmd := fields[0]
		num, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(fmt.Errorf("unexpected num: %v", fields[1]))
		}
		result[i] = c.Tuple[string, int]{First: cmd, Second: num}
	}

	return result
}

func (r Runner8) Run1(cmds Input8) int {
	cmdRun := make([]int, len(cmds))
	acc := 0
	currCmd := 0
	for cmdRun[currCmd] == 0 {
		cmd, val := cmds[currCmd].Values()
		cmdRun[currCmd]++
		switch cmd {
		case "nop":
			currCmd++
		case "jmp":
			currCmd += val
		case "acc":
			acc += val
			currCmd++
		default:
			panic(fmt.Errorf("unexpected cmd: %v", cmd))
		}
	}
	fmt.Printf("acc: %v\n", acc)
	return acc
}

func (r Runner8) Run2(cmds Input8) int {
	stack := make([]int, len(cmds))
	acc := 0
	currCmd := 0
	corruptedLine := findCorruptedLine(cmds, stack, 0, -1)
	fmt.Printf("corrupted: %v\n", corruptedLine)
	stack = make([]int, len(cmds))
	for currCmd < len(cmds) && stack[currCmd] == 0 {
		cmd, val := cmds[currCmd].Values()
		stack[currCmd]++
		if cmd == "acc" {
			acc += val
		}
		currCmd = getNextLine(cmds, currCmd, corruptedLine)
	}
	fmt.Printf("acc: %v\n", acc)
	return acc
}

func findCorruptedLine(cmds Input8, stack []int, currLine int, corruptedLine int) int {
	fmt.Printf("running: %v %v %v\n", currLine+1, corruptedLine+1, stack)
	if currLine >= len(cmds) {
		return corruptedLine
	}
	if stack[currLine] > 0 {
		return -1
	}

	stack[currLine]++
	cmd, _ := cmds[currLine].Values()
	nextLine := getNextLine(cmds, currLine, corruptedLine)
	res := findCorruptedLine(cmds, stack, nextLine, corruptedLine)
	if res > 0 {
		return res
	} else {
		stack[nextLine]--
		if cmd == "nop" || cmd == "jmp" {
			nextLine := getNextLine(cmds, currLine, currLine)
			return findCorruptedLine(cmds, stack, nextLine, currLine)
		} else {
			return -1
		}
	}
}

func getNextLine(cmds Input8, currLine int, corruptedLine int) int {
	nextLine := currLine
	cmd, val := cmds[currLine].Values()
	if currLine == corruptedLine {
		switch cmd {
		case "nop":
			cmd = "jmp"
		case "jmp":
			cmd = "nop"
		}
	}
	switch cmd {
	case "nop":
		nextLine++
	case "jmp":
		nextLine += val
	case "acc":
		nextLine++
	default:
		panic(fmt.Errorf("unexpected cmd: %v", cmd))
	}
	return nextLine
}
