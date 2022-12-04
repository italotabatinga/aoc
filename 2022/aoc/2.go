package aoc

import (
	"strings"
)

type Input2 [][]string

type Runner2 struct{}

func (r Runner2) FmtInput(input string) Input2 {
	gamesString := strings.Split(input, "\n")
	result := make(Input2, 0)
	for _, gameString := range gamesString {
		picksString := strings.Split(gameString, " ")
		result = append(result, picksString)
	}

	return result
}

func (r Runner2) Run1(input Input2, _ bool) int {
	sum := 0
	for _, picksString := range input {

		var leftPick, rightPick int
		switch picksString[0] {
		case "A":
			leftPick = ROCK
		case "B":
			leftPick = PAPER
		case "C":
			leftPick = SCISSORS
		}

		switch picksString[1] {
		case "X":
			rightPick = ROCK
		case "Y":
			rightPick = PAPER
		case "Z":
			rightPick = SCISSORS
		}
		game := Game{leftPick: leftPick, rightPick: rightPick}
		sum += game.score()
	}

	return sum
}

func (r Runner2) Run2(input Input2, _ bool) int {
	sum := 0
	for _, picksString := range input {

		var leftPick, rightPick int
		switch picksString[0] {
		case "A":
			leftPick = ROCK
		case "B":
			leftPick = PAPER
		case "C":
			leftPick = SCISSORS
		}

		switch picksString[1] {
		case "X":
			switch leftPick {
			case ROCK:
				rightPick = SCISSORS
			case PAPER:
				rightPick = ROCK
			case SCISSORS:
				rightPick = PAPER
			}
		case "Y":
			rightPick = leftPick
		case "Z":
			switch leftPick {
			case ROCK:
				rightPick = PAPER
			case PAPER:
				rightPick = SCISSORS
			case SCISSORS:
				rightPick = ROCK
			}
		}
		game := Game{leftPick: leftPick, rightPick: rightPick}
		sum += game.score()
	}

	return sum
}

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
	VICTORY  = 6
	DRAW     = 3
	LOSE     = 0
)

type Game struct {
	leftPick  int
	rightPick int
}

func (g Game) result() int {
	if (g.rightPick == ROCK && g.leftPick == SCISSORS) || (g.rightPick == PAPER && g.leftPick == ROCK) || (g.rightPick == SCISSORS && g.leftPick == PAPER) {
		return VICTORY
	} else if g.leftPick == g.rightPick {
		return DRAW
	}

	return LOSE
}

func (g Game) score() int {
	return g.rightPick + g.result()
}
