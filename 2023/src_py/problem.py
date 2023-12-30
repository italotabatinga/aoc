from enum import Enum
import os


class Part(Enum):
    ONE = 1
    TWO = 2

    def __str__(self) -> str:
        if self == Part.ONE:
            return "1"
        elif self == Part.TWO:
            return "2"
        else:
            return "unknown"


class Problem:
    def __init__(self, day: int, part: Part, test: bool):
        self.day = day
        self.part = part
        self.test = test


def parseProblem(problem: str, test: bool) -> Problem:
    part = Part.ONE

    dayStr, partStr = problem.split(".")  # 3.1
    if partStr == "1":
        part = Part.ONE
    elif partStr == "2":
        part = Part.TWO

    return Problem(int(dayStr), part, test)


def read_input(problem: Problem) -> list[str]:
    compact_file_string = "files/" + str(problem.day)
    complete_file_string = "files/" + \
        str(problem.day) + "." + str(problem.part)
    if problem.test:
        compact_file_string += "_test"
        complete_file_string += "_test"
    compact_file_string += ".txt"
    complete_file_string += ".txt"

    if os.path.exists(complete_file_string):
        file = open(complete_file_string, "r")
    elif os.path.exists(compact_file_string):
        file = open(compact_file_string, "r")
    else:
        print(f"Could not open file for problem {problem}")
        exit(1)

    return file.read().splitlines()
