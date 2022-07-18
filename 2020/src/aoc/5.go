package aoc

import (
	"fmt"
	"strings"
)

type Input5 []BoardingPass

type Runner5 struct{}

func (r Runner5) FmtInput(input string) Input5 {
	lines := strings.Split(input, "\n")
	var result Input5
	count := len(lines)
	result = make(Input5, count)
	for i, s := range lines {
		boardingPass := parseRawBoardingPass(s)
		result[i] = boardingPass
	}

	return result
}

func (r Runner5) Run1(input Input5) int {
	max := -1
	for _, bpass := range input {
		if bpass.id > max {
			max = bpass.id
		}
	}
	fmt.Printf("max %v\n", max)
	return max
}

func (r Runner5) Run2(input Input5) int {
	seats := [128][8]bool{}
	for _, bpass := range input {
		seats[bpass.row][bpass.col] = true
	}
	fmt.Printf("    01234567\n")
	for i, row := range seats {
		fmt.Printf("%3v ", i)
		for _, occupied := range row {
			if occupied {
				fmt.Printf("_")
			} else {
				fmt.Printf("O")
			}
		}
		fmt.Printf("\n")
	}
	return 0
}

type BoardingPass struct {
	id  int
	row int
	col int
	raw string
}

func parseRawBoardingPass(raw string) BoardingPass {
	row, col := 0, 0
	runes := []rune(raw)
	rowRunes := runes[:7]
	colRunes := runes[7:]
	for _, rune := range rowRunes {
		row <<= 1
		switch rune {
		case 'F':
		case 'B':
			row += 1
		default:
			panic(fmt.Errorf("unexpected rune: %v", rune))
		}
	}

	for _, rune := range colRunes {
		col <<= 1
		switch rune {
		case 'R':
			col += 1
		case 'L':
		default:
			panic(fmt.Errorf("unexpected rune: %v", rune))
		}
	}

	id := row*8 + col
	return BoardingPass{id, row, col, raw}
}
