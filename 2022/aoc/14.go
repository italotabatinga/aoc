package aoc

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	c "github.com/italotabatinga/aoc/2022/aoc/collections"
)

type Input14 Cavern

type Runner14 struct{}

func (r Runner14) FmtInput(input string) Input14 {
	paths := strings.Split(input, "\n")
	result := Cavern{
		Blocks: make(map[string]int, 0),
		xBound: c.Pair[int]{First: math.MaxInt, Second: math.MinInt},
		yBound: c.Pair[int]{First: math.MaxInt, Second: math.MinInt},
	}
	for _, path := range paths {
		coordStrings := strings.Split(path, " -> ")
		packetStrings := strings.Split(coordStrings[0], ",")
		prevX, _ := strconv.Atoi(packetStrings[0])
		prevY, _ := strconv.Atoi(packetStrings[1])

		for _, coordString := range coordStrings {
			packetStrings := strings.Split(coordString, ",")
			x, _ := strconv.Atoi(packetStrings[0])
			y, _ := strconv.Atoi(packetStrings[1])

			diffX, diffY := (x - prevX), (y - prevY)
			if diffX != 0 {
				diffX /= Abs(diffX)
			}
			if diffY != 0 {
				diffY /= Abs(diffY)
			}
			result.Blocks[result.coordString(x, y)] = C_ROCK
			for i, j := prevX, prevY; i != x || j != y; i, j = i+diffX, j+diffY {
				result.Blocks[result.coordString(i, j)] = C_ROCK
			}

			if y < result.yBound.First {
				result.yBound.First = y
			}
			if y > result.yBound.Second {
				result.yBound.Second = y
			}
			if x < result.xBound.First {
				result.xBound.First = x
			}
			if x > result.xBound.Second {
				result.xBound.Second = x
			}

			prevX = x
			prevY = y
		}
	}

	return Input14(result)
}

func (r Runner14) Run1(input Input14, _ bool) int {
	cavern := Cavern(input)
	count := 0
	for cavern.DropSand() != 0 {
		count++
	}
	return count
}

func (r Runner14) Run2(input Input14, _ bool) int {
	cavern := Cavern(input)
	cavern.hasFloor = true
	cavern.Floor = cavern.yBound.Second + 2
	count := 0
	for ; cavern.GetBlock(500, 0) != SAND; cavern.DropSand() {
		count++
	}
	return count
}

type Cavern struct {
	Blocks   map[string]int
	xBound   c.Pair[int]
	yBound   c.Pair[int]
	Floor    int
	hasFloor bool
}

func (c *Cavern) DropSand() int {
	x, y := 500, 0
	for c.hasFloor || y < c.yBound.Second {
		if c.GetBlock(x, y+1) == AIR {
			y++
		} else if c.GetBlock(x-1, y+1) == AIR {
			x--
			y++
		} else if c.GetBlock(x+1, y+1) == AIR {
			x++
			y++
		} else {
			c.Blocks[c.coordString(x, y)] = SAND
			return 1
		}
	}

	return 0
}

func (c Cavern) GetBlock(x int, y int) int {
	if c.hasFloor && c.Floor == y {
		return ROCK
	} else if val, ok := c.Blocks[c.coordString(x, y)]; ok {
		return val
	}

	return AIR
}

func (c Cavern) coordString(x int, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}

func (c Cavern) String() string {
	var sb strings.Builder
	minY := Min(c.yBound.First, 0)
	for y := minY; y <= c.yBound.Second; y++ {
		for x := c.xBound.First; x <= c.xBound.Second; x++ {
			if x == 500 && y == 0 {
				sb.WriteRune('+')
			} else {
				switch c.GetBlock(x, y) {
				case AIR:
					sb.WriteRune('.')
				case C_ROCK:
					sb.WriteRune('#')
				case SAND:
					sb.WriteRune('o')
				default:
					sb.WriteRune('?')
				}
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

const (
	AIR = iota
	C_ROCK
	SAND
)
