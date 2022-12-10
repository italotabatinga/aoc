package aoc

import (
	"strconv"
	"strings"

	"github.com/italotabatinga/aoc/2020/src"
)

type Input12 []ShipCommand

type Runner12 struct{}

func (r Runner12) FmtInput(input string) Input12 {
	lines := strings.Split(input, "\n")
	result := []ShipCommand{}
	for _, line := range lines {
		dir := direction(line[0])
		value, _ := strconv.Atoi(line[1:])

		result = append(result, ShipCommand{Value: value, Dir: dir})
	}
	return result
}

func (r Runner12) Run1(input Input12, _ bool) int {
	ship := Ship{CurrDir: EAST}

	for _, command := range input {
		ship.ProcessCommand(command)
	}
	return src.Abs(ship.East) + src.Abs(ship.South)
}

func (r Runner12) Run2(input Input12, _ bool) int {
	ship := Ship{CurrDir: EAST, Waypoint: Waypoint{East: 10, South: -1}}

	for _, command := range input {
		switch command.Dir {
		case FORWARD:
			ship.ForwardWaypoint(command.Value)
		default:
			ship.Waypoint.ProcessCommand(command)
		}
	}
	return src.Abs(ship.East) + src.Abs(ship.South)
}

type Ship struct {
	CurrDir  direction
	Waypoint Waypoint
	East     int
	South    int
}

type Waypoint struct {
	East  int
	South int
}

type ShipCommand struct {
	Dir   direction
	Value int
}

type direction rune

const (
	EAST    = 'E'
	WEST    = 'W'
	NORTH   = 'N'
	SOUTH   = 'S'
	LEFT    = 'L'
	RIGHT   = 'R'
	FORWARD = 'F'
)

func (s *Ship) ProcessCommand(c ShipCommand) {
	switch c.Dir {
	case EAST:
		s.East += c.Value
	case WEST:
		s.East -= c.Value
	case SOUTH:
		s.South += c.Value
	case NORTH:
		s.South -= c.Value
	case RIGHT:
		turnCount := c.Value / 90
		for i := 0; i < turnCount; i++ {
			s.TurnRight()
		}
	case LEFT:
		turnCount := c.Value / 90
		for i := 0; i < turnCount; i++ {
			s.TurnLeft()
		}
	case FORWARD:
		s.ProcessCommand(ShipCommand{Dir: s.CurrDir, Value: c.Value})
	}
}

func (s *Ship) TurnRight() {
	switch s.CurrDir {
	case EAST:
		s.CurrDir = SOUTH
	case WEST:
		s.CurrDir = NORTH
	case SOUTH:
		s.CurrDir = WEST
	case NORTH:
		s.CurrDir = EAST
	}
}

func (s *Ship) TurnLeft() {
	switch s.CurrDir {
	case EAST:
		s.CurrDir = NORTH
	case WEST:
		s.CurrDir = SOUTH
	case SOUTH:
		s.CurrDir = EAST
	case NORTH:
		s.CurrDir = WEST
	}
}

func (s *Ship) ForwardWaypoint(count int) {
	for i := 0; i < count; i++ {
		s.East += s.Waypoint.East
		s.South += s.Waypoint.South
	}
}

func (w *Waypoint) ProcessCommand(c ShipCommand) {
	switch c.Dir {
	case EAST:
		w.East += c.Value
	case WEST:
		w.East -= c.Value
	case SOUTH:
		w.South += c.Value
	case NORTH:
		w.South -= c.Value
	case RIGHT:
		turnCount := c.Value / 90
		for i := 0; i < turnCount; i++ {
			w.TurnRight()
		}
	case LEFT:
		turnCount := c.Value / 90
		for i := 0; i < turnCount; i++ {
			w.TurnLeft()
		}
	}
}

func (w *Waypoint) TurnRight() {
	tmp := w.South
	w.South = w.East
	w.East = -tmp
}

func (w *Waypoint) TurnLeft() {
	tmp := w.South
	w.South = -w.East
	w.East = tmp
}
