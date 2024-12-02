package aoc

import "core:strconv"
import "core:fmt"
import "core:strings"

run_1_1 :: proc(input: string, test: bool) -> string {
  input := input
  sum := 0
  for line in strings.split_lines_iterator(&input) {
    mass := strconv.atoi(line)
    fuel := (mass / 3) - 2
    sum += fuel
  }

	buf: [8]byte
	return strings.clone(strconv.itoa(buf[:], sum))
}

run_1_2 :: proc(input: string, test: bool) -> string {
  input := input
  sum := 0
  for line in strings.split_lines_iterator(&input) {
    mass := strconv.atoi(line)
    for mass > 6 {
      fuel := (mass / 3) - 2
      mass = fuel
      sum += fuel
    }
  }

	buf: [8]byte
	return strings.clone(strconv.itoa(buf[:], sum))
}
