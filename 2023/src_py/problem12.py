import logging
from typing import Any
from src_py.problem import Problem, Part


class SpringGroup:
    def __init__(self, springs: list[str]) -> 'SpringGroup':
        self.springs = springs
        self.unknown_count = self.springs.count('?')
        self.broken_count = self.springs.count('#')

    def __str__(self) -> str:
        return f"SG({self.springs})"

    def __repr__(self) -> str:
        return self.__str__()

    def __len__(self) -> int:
        return len(self.springs)


def run(problem: Problem, input: list[str]) -> str:
    count = 0
    for line in input:
        groups = []
        line, validation_str = line.split()
        groups_validation = list(map(int, validation_str.split(',')))
        first_not_working = -1
        for i, c in enumerate(line):
            if c == '.':
                if first_not_working >= 0:
                    groups.append(SpringGroup(line[first_not_working:i]))

                first_not_working = -1
            elif c != '.':
                if first_not_working < 0:
                    first_not_working = i
        if first_not_working >= 0:
            groups.append(SpringGroup(line[first_not_working:]))

        logging.debug(f"g {groups}, v {groups_validation}")

    return str(count)
