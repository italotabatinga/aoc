import sys

from src_py.problem import parseProblem, readInput


if __name__ == "__main__":
    argv = sys.argv

    problem = parseProblem(argv[1:])
    input = readInput(problem)

    print(
        f"Running {problem.day}.{problem.part}{'-test' if problem.test else ''}")
