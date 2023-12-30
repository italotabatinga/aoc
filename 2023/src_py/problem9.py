import logging
from src_py.problem import Problem, Part


def part_one(input: list[str]) -> int:
    nexts = []

    for line in input:
        seq = list(map(int, line.split()))
        all_zeros = False
        iteration = 0
        logging.debug(f"seq: {seq}")
        while not all_zeros:
            all_zeros = True
            for i in range(len(seq) - iteration - 1):
                seq[i] = seq[i + 1] - seq[i]
                if seq[i] != 0:
                    all_zeros = False

            iteration += 1

            logging.debug(f"\titeration: {iteration}, seq: {seq}")
        nexts.append(sum(seq[len(seq) - iteration:]))

    logging.debug(f"\tnexts: {nexts}")
    return sum(nexts)


def part_two(input: list[str]) -> int:
    prevs = []

    for line in input:
        seq = list(map(int, line.split()))
        all_zeros = False
        iteration = 0
        logging.debug(f"seq: {seq}")
        while not all_zeros:
            all_zeros = True
            for i in range(len(seq)-1, iteration, -1):
                seq[i] = seq[i] - seq[i-1]
                if seq[i] != 0:
                    all_zeros = False

            iteration += 1

            logging.debug(f"\titeration: {iteration}, seq: {seq}")
        prev = seq[iteration]
        logging.debug(f"\t\titeration: {iteration}, prev: {prev}")
        for i in range(iteration+1):
            prev = seq[iteration - i] - prev
            logging.debug(
                f"\t\titeration: {iteration}, i: {i}, prev: {prev}, seq: {seq[iteration - i]}")
        prevs.append(prev)

    logging.debug(f"\tprevs: {prevs}")
    return sum(prevs)


def run(problem: Problem, input: list[str]) -> str:
    if problem.part == Part.ONE:
        return str(part_one(input))

    return str(part_two(input))
