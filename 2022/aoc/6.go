package aoc

type Input6 []rune

type Runner6 struct{}

func (r Runner6) FmtInput(input string) Input6 {
	result := []rune(input)
	return result
}

func (r Runner6) Run1(input Input6, _ bool) int {
	return findStarterOfPacket(input, 4)
}

func (r Runner6) Run2(input Input6, _ bool) int {
	return findStarterOfPacket(input, 14)
}

func findStarterOfPacket(runes []rune, length int) int {
	for i := length - 1; i < len(runes); i++ {
		anyEqual := false
		for j := i - length + 1; j <= i && !anyEqual; j++ {
			for k := j + 1; k <= i && !anyEqual; k++ {
				if runes[j] == runes[k] {
					anyEqual = true
				}
			}
		}

		if !anyEqual {
			return i + 1
		}
	}

	return -1
}
