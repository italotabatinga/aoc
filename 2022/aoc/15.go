package aoc

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Input15 SensorField

type Runner15 struct{}

func (r Runner15) FmtInput(input string) Input15 {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	result := SensorField{Sensors: []Sensor{}, elements: make(map[string]int)}
	for _, line := range lines {
		submatches := re.FindStringSubmatch(line)
		sX, _ := strconv.Atoi(submatches[1])
		sY, _ := strconv.Atoi(submatches[2])
		bX, _ := strconv.Atoi(submatches[3])
		bY, _ := strconv.Atoi(submatches[4])

		manhattan := ManhattanDistance(sX, bX, sY, bY)
		result.elements[result.coordString(sX, sY)] = SENSOR
		result.elements[result.coordString(bX, bY)] = BEACON
		result.Sensors = append(result.Sensors, Sensor{x: sX, y: sY, closestManhattan: manhattan})
	}
	return Input15(result)
}

func (r Runner15) Run1(input Input15, isTest bool) int {
	field := SensorField(input)
	var goalLine int
	if isTest {
		goalLine = 10
	} else {
		goalLine = 2000000
	}
	minX, maxX := math.MaxInt, math.MinInt

	for _, sensor := range field.Sensors {
		minSX := sensor.x - sensor.closestManhattan
		maxSX := sensor.x + sensor.closestManhattan
		if minSX < minX {
			minX = minSX
		}
		if maxSX > maxX {
			maxX = maxSX
		}
	}
	count := 0
	for x := minX; x <= maxX; x++ {
		if field.GetBlock(x, goalLine) != EMPTY {
			continue
		}

		for _, sensor := range field.Sensors {
			manhattan := ManhattanDistance(x, sensor.x, goalLine, sensor.y)

			if manhattan <= sensor.closestManhattan {
				count++
				break
			}
		}
	}

	return count
}

func (r Runner15) Run2(input Input15, _ bool) int {
	return 0
}

type SensorField struct {
	Sensors  []Sensor
	elements map[string]int
}

type Sensor struct {
	x, y             int
	closestManhattan int
}

func ManhattanDistance(x1, x2, y1, y2 int) int {
	return Abs(x2-x1) + Abs(y2-y1)
}

func (c SensorField) GetBlock(x int, y int) int {
	if val, ok := c.elements[c.coordString(x, y)]; ok {
		return val
	}

	return EMPTY
}

func (s SensorField) coordString(x int, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}

const (
	SENSOR = iota
	BEACON
	EMPTY
)
