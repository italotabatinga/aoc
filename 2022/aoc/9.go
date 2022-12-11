package aoc

import (
	"strconv"
	"strings"

	c "github.com/italotabatinga/aoc/2022/aoc/collections"
)

type Input9 []HeadMotion

type Runner9 struct{}

func (r Runner9) FmtInput(input string) Input9 {
	lines := strings.Split(input, "\n")
	result := make(Input9, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")
		value, _ := strconv.Atoi(split[1])
		result[i] = HeadMotion{direction: rune(split[0][0]), value: value}
	}
	return result
}

func (r Runner9) Run1(input Input9, _ bool) int {
	rope := NewRope(2)
	for _, motion := range input {

		rope.MoveHead(motion)
	}
	return rope.tailVisited.Size()
}

func (r Runner9) Run2(input Input9, _ bool) int {
	rope := NewRope(10)
	for _, motion := range input {

		rope.MoveHead(motion)
	}
	return rope.tailVisited.Size()
}

type HeadMotion struct {
	direction rune
	value     int
}

type Rope struct {
	head        *c.Pair[int]
	tail        *c.Pair[int]
	knots       []c.Pair[int]
	tailVisited c.Set[string]
}

func NewRope(knotsLength int) Rope {

	tailVisited := c.NewSet[string]()
	tailVisited.Add("0_0")

	knots := make([]c.Pair[int], knotsLength)
	head := &knots[0]
	tail := &knots[len(knots)-1]
	return Rope{
		head:        head,
		tail:        tail,
		knots:       knots,
		tailVisited: tailVisited,
	}
}

func (r *Rope) MoveHead(motion HeadMotion) {
	var direction c.Pair[int]
	switch motion.direction {
	case RIGHT:
		direction.First = 1
	case LEFT:
		direction.First = -1
	case UP:
		direction.Second = 1
	case DOWN:
		direction.Second = -1
	}
	for i := 0; i < motion.value; i++ {
		r.head.First += direction.First
		r.head.Second += direction.Second

		for j := 1; j < len(r.knots); j++ {
			r.MoveKnot(j)
		}
		r.RecordTail()
	}
}

func (r *Rope) MoveKnot(index int) {
	headKnot := &r.knots[index-1]
	tailKnot := &r.knots[index]
	xDiff := headKnot.First - tailKnot.First
	yDiff := headKnot.Second - tailKnot.Second

	if Abs(xDiff) <= 1 && Abs(yDiff) <= 1 {
		return
	} else if xDiff == 0 {
		tailKnot.Second += yDiff / Abs(yDiff)
	} else if yDiff == 0 {
		tailKnot.First += xDiff / Abs(xDiff)
	} else {
		tailKnot.Second += yDiff / Abs(yDiff)
		tailKnot.First += xDiff / Abs(xDiff)
	}
}

func (r *Rope) RecordTail() {
	r.tailVisited.Add(strconv.Itoa(r.tail.First) + "_" + strconv.Itoa(r.tail.Second))
}

const (
	RIGHT rune = 'R'
	LEFT  rune = 'L'
	DOWN  rune = 'D'
	UP    rune = 'U'
)
