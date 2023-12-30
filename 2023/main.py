import argparse
import logging

from src_py.problem import parseProblem, read_input
import src_py.problem5 as p5


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Advent of Code 2023")
    parser.add_argument('problem')
    parser.add_argument('-l', '--log', default='INFO', dest="log",
                        help='Set the logging level')
    parser.add_argument('-t', '--test', action='store_true', dest="test",
                        default=False, help='Set the logging level')
    return parser.parse_args()


def setup_logging(level: str) -> None:
    levels = {'DEBUG': logging.DEBUG, 'INFO': logging.INFO, 'WARNING': logging.WARNING,
              'ERROR': logging.ERROR, 'CRITICAL': logging.CRITICAL}
    level = levels.get(args.log.upper(), logging.WARNING)
    logging.basicConfig(level=level)


if __name__ == "__main__":
    args = parse_args()
    setup_logging(args.log)

    problem = parseProblem(args.problem, args.test)
    input = read_input(problem)

    logging.debug(
        f"Running {problem.day}.{problem.part}{'-test' if problem.test else ''}")

    result = ""
    if problem.day == 5:
        result = p5.run(problem, input)
    else:
        logging.error("Problem not implemented")
        exit(1)

    print(result)
