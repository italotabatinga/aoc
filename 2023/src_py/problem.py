from enum import Enum


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


def parseProblem(argv: list[str]) -> Problem:
    if len(argv) == 0:
        print("No problem specified, use format: [--test] 3.1")
        exit(1)

    part = Part.ONE
    test = False
    if "--test" in argv:
        test = True

    dayStr, partStr = argv[-1].split(".")  # 3.1
    if partStr == "1":
        part = Part.ONE
    elif partStr == "2":
        part = Part.TWO

    return Problem(int(dayStr), part, test)


def readInput(problem: Problem) -> list[str]:
    fileString = "files/" + str(problem.day) + "." + str(problem.part)
    if problem.test:
        fileString += "_test"
    fileString += ".txt"
    file = open(fileString, "r")
    if file == None:
        print("Could not open file " + fileString)
        exit(1)

    return file.readlines()
