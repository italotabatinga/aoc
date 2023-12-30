import logging
from time import sleep
from src_py.problem import Problem, Part


def run(problem: Problem, input: list[str]) -> str:
    times = list(map(int, input[0].split()[1:]))
    distances = list(map(int, input[1].split()[1:]))

    if problem.part == Part.TWO:
        times = [int(''.join(map(str, times)))]
        distances = [int(''.join(map(str, distances)))]

    logging.debug(f"times: {times}")
    logging.debug(f"distances: {distances}")

    possibilities = []

    for time, distance in zip(times, distances):
        logging.debug(f"time: {time}, distance: {distance}")

        start = 0
        end = time // 2
        while start < end:
            mid = (start + end) // 2
            if mid * (time - mid) > distance:
                end = mid
            else:
                start = mid + 1

        logging.debug(f"\tfound: {start}")
        possibilities.append(round((time / 2 - start) * 2) + 1)

    logging.debug(f"possibilities: {possibilities}")
    result = 1
    for possibility in possibilities:
        result *= possibility
    return str(result)
