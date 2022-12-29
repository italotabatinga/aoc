package aoc

import (
	"fmt"
	"strconv"
	"strings"

	c "github.com/italotabatinga/aoc/2022/aoc/collections"
)

type Input11 KeepAway

type Runner11 struct{}

func (r Runner11) FmtInput(input string) Input11 {
	monkeyStrings := strings.Split(input, "\n\n")
	result := KeepAway{monkeys: []*Monkey{}, divisibleBy: 1}
	for _, monkeyString := range monkeyStrings {
		lines := strings.Split(monkeyString, "\n")
		itemStrings := strings.Split(lines[1][18:], ", ")
		items := c.NewQueue[*Item]()
		for _, itemString := range itemStrings {
			if worryLevel, err := strconv.Atoi(itemString); err == nil {

				items.Push(&Item{worryLevel})
			}
		}

		// op
		eqString := strings.Split(lines[2][19:], " ")
		var operation func(int) int
		if eqString[1] == "*" && eqString[2] == "old" {
			operation = func(worryLevel int) int { return worryLevel * worryLevel }
		} else if eqString[1] == "+" && eqString[2] == "old" {
			operation = func(worryLevel int) int { return worryLevel + worryLevel }
		} else if eqString[1] == "*" && eqString[2] != "old" {
			val, _ := strconv.Atoi(eqString[2])
			operation = func(worryLevel int) int { return worryLevel * val }
		} else if eqString[1] == "+" && eqString[2] != "old" {
			val, _ := strconv.Atoi(eqString[2])
			operation = func(worryLevel int) int { return worryLevel + val }
		}

		// test
		divisibleBy, _ := strconv.Atoi(lines[3][21:])
		monkeyIfTrue, _ := strconv.Atoi(lines[4][29:])
		monkeyIfFalse, _ := strconv.Atoi(lines[5][30:])

		result.divisibleBy *= divisibleBy
		result.monkeys = append(result.monkeys, &Monkey{
			items:       items,
			operation:   operation,
			divisibleBy: divisibleBy,
			test: func(worryLevel int) int {
				if worryLevel%divisibleBy == 0 {
					return monkeyIfTrue
				} else {
					return monkeyIfFalse
				}
			}})
	}
	return Input11(result)
}

func (r Runner11) Run1(input Input11, _ bool) int {
	keepAway := KeepAway(input)
	keepAway.reliefAfterRound = true
	numRounds := 20
	for i := 0; i < numRounds; i++ {
		keepAway.TakeRound()
	}
	return keepAway.MonkeyBusiness()
}

func (r Runner11) Run2(input Input11, _ bool) int {
	keepAway := KeepAway(input)
	numRounds := 10000
	for i := 0; i < numRounds; i++ {
		keepAway.TakeRound()
	}
	return keepAway.MonkeyBusiness()
}

type KeepAway struct {
	monkeys          []*Monkey
	reliefAfterRound bool
	divisibleBy      int
}

type Monkey struct {
	items          c.Queue[*Item]
	countInspected int
	operation      func(int) int
	test           func(int) int
	divisibleBy    int
}

type Item struct {
	worryLevel int
}

func (ka *KeepAway) TakeRound() {
	for _, monkey := range ka.monkeys {
		ka.TakeTurn(monkey)
	}
}

func (ka *KeepAway) TakeTurn(m *Monkey) {
	for m.items.Len() > 0 {
		m.countInspected++
		item := m.items.Pop()
		item.worryLevel = m.operation(item.worryLevel)
		if ka.reliefAfterRound {
			item.worryLevel /= 3
		}
		nextMonkey := m.test(item.worryLevel)
		item.worryLevel %= ka.divisibleBy // modulo math for part 2

		ka.monkeys[nextMonkey].items.Push(item)
	}
}

func (ka KeepAway) MonkeyBusiness() int {
	inspectedCounts := []int{}
	for _, monkey := range ka.monkeys {
		inspectedCounts = append(inspectedCounts, monkey.countInspected)
	}

	Sort(inspectedCounts, func(a, b int) int { return Compare(b, a) })
	return inspectedCounts[len(inspectedCounts)-1] * inspectedCounts[len(inspectedCounts)-2]
}

func (ka KeepAway) String() string {
	var sb strings.Builder

	for i, monkey := range ka.monkeys {
		sb.WriteString(fmt.Sprintf("Monkey(%d) %d: ", monkey.divisibleBy, i))
		for j := 0; j < monkey.items.Len(); j++ {
			item := monkey.items.Index(j)
			sb.WriteString(strconv.Itoa(item.worryLevel))
			if j != monkey.items.Len()-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString(fmt.Sprintf(" (%d)\n", monkey.countInspected))
	}

	return sb.String()
}
