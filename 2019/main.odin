package aoc

import "core:fmt"
import "core:os"
import "core:path/filepath"
import "core:strconv"
import "core:strings"
import "core:slice"

main :: proc() {
	if len(os.args) <= 1 {
		error("Expected problem (e.g. 1.1 [--test]) as argument\n")
	}

  problem := os.args[1]
  test := slice.contains(os.args, "--test")
  solver := solvers[problem]
  input := read_input(problem, test)
  result := solver(input, test)

  fmt.printf("result: %v\n", result)
}

solvers := map[string](proc(input: string, test: bool) -> string) {
 "1.1" = run_1_1,
 "1.2" = run_1_2
}

error :: proc(message: string, args: ..any) {
	fmt.printf(message, ..args)
	os.exit(1)
}

read_input :: proc(problem: string, test: bool) -> string {
  filenameBuilder : strings.Builder
  fmt.sbprintf(&filenameBuilder, problem)
  if test { fmt.sbprintf(&filenameBuilder, "_test") }
  filename := fmt.sbprintf(&filenameBuilder, ".txt")
  path := filepath.join({"files", filename})

  data, ok := os.read_entire_file(path)
  if !ok { error("Couldn't read file: %v\n", path) }
  defer delete(data, context.allocator)

  return strings.clone_from(data)
}
