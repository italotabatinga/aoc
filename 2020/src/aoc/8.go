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
	acc := 0
	currCmd := 0
	corruptedLine := findCorruptedLine(cmds)
	fmt.Printf("corrupted: %v\n", corruptedLine)
	for currCmd < len(cmds) {
		cmd, val := cmds[currCmd].Values()
		if cmd == "acc" {
			acc += val
		}
		currCmd = getNextLine(cmds, currCmd, corruptedLine)
	}
	fmt.Printf("acc: %v\n", acc)
	return acc
}

func findCorruptedLine(cmds Input8) int {
	for line := range cmds {
		cmd, _ := cmds[line].Values()
		switch cmd {
		case "jmp", "nop":
			if !cmdsHasLoop(cmds, line) {
				return line
			}
		}
	}

	return -1
}

func cmdsHasLoop(cmds Input8, corruptedLine int) bool {
	stack := make([]int, len(cmds))

	line := 0
	for line < len(cmds) {
		if stack[line] > 0 {
			return true
		}
		stack[line]++

		line = getNextLine(cmds, line, corruptedLine)
	}

	return false
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
