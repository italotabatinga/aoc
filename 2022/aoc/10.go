package aoc

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
)

type Input10 []CpuInst

type Runner10 struct{}

func (r Runner10) FmtInput(input string) Input10 {
	lines := strings.Split(input, "\n")
	result := make(Input10, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")
		var kind, cycleCount, value int
		switch split[0] {
		case "noop":
			kind = NOOP
			cycleCount = 1
		case "addx":
			kind = ADDX
			cycleCount = 2
			value, _ = strconv.Atoi(split[1])
		}
		result[i] = CpuInst{kind, value, cycleCount}
	}
	return result
}

func (r Runner10) Run1(input Input10, _ bool) int {
	maxCycles := 220
	clock := NewClockCircuit()
	sum := 0
	nextInstruction := 0
	for cycle := 0; cycle < maxCycles; cycle++ {
		if (clock.Cycle+20)%40 == 0 {
			sum += clock.SignalStrength()
		}
		if clock.EmptyIntruction() {
			clock.AddInstruction(input[nextInstruction])
			nextInstruction++
		}

		clock.Tick()
	}
	return sum
}

func (r Runner10) Run2(input Input10, _ bool) int {
	maxCycles := 240
	clock := NewClockCircuit()
	sum := 0
	nextInstruction := 0
	for cycle := 0; cycle < maxCycles; cycle++ {
		if (clock.Cycle+20)%40 == 0 {
			sum += clock.SignalStrength()
		}
		if clock.EmptyIntruction() {
			clock.AddInstruction(input[nextInstruction])
			nextInstruction++
		}

		clock.Tick()
	}

	// fmt.Printf("Displayyy\n%v\n", clock.PrintDisplay())
	h := fnv.New32a()
	h.Write([]byte(clock.PrintDisplay()))
	return int(h.Sum32())
}

type ClockCircuit struct {
	register        int
	display         [240]bool
	Cycle           int
	currInstruction *CpuInst
}

type CpuInst struct {
	kind       int
	value      int
	cycleCount int
}

func NewClockCircuit() ClockCircuit {
	return ClockCircuit{
		Cycle:    1,
		register: 1,
	}
}

func (c *ClockCircuit) Tick() {
	c.WriteToDisplay()
	c.Cycle += 1

	if c.currInstruction != nil {
		inst := c.currInstruction
		inst.cycleCount -= 1

		if inst.cycleCount == 0 {
			switch inst.kind {
			case ADDX:
				c.register += inst.value
			}
			c.currInstruction = nil
		}
	}
}

func (c *ClockCircuit) WriteToDisplay() {
	xPos := (c.Cycle - 1) % 40
	spriteDiff := c.register - xPos
	displayPos := (c.Cycle - 1) % 240
	if Abs(spriteDiff) <= 1 {
		c.display[displayPos] = true
	} else {
		c.display[displayPos] = false
	}
}

func (c ClockCircuit) EmptyIntruction() bool {
	return c.currInstruction == nil
}

func (c *ClockCircuit) AddInstruction(i CpuInst) {
	if c.currInstruction == nil {
		c.currInstruction = &i
	}
}

func (c ClockCircuit) SignalStrength() int {
	return c.Cycle * c.register
}

func (c ClockCircuit) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("{cycle: %d, X: %d, i: ", c.Cycle, c.register))
	if c.EmptyIntruction() {
		sb.WriteString("nil")
	} else {
		sb.WriteString(fmt.Sprintf("%v", *c.currInstruction))
	}
	sb.WriteRune('}')
	return sb.String()
}

func (c ClockCircuit) PrintDisplay() string {
	var sb strings.Builder
	for i, val := range c.display {
		if i > 0 && i%40 == 0 {
			sb.WriteRune('\n')
		}
		if val {
			sb.WriteRune('#')
		} else {
			sb.WriteRune('.')
		}
	}
	return sb.String()
}

const (
	NOOP int = iota
	ADDX
)
