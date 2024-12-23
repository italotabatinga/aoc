import argparse
import logging

from src_py.problem import parseProblem, read_input
import src_py.problem5 as p5
import src_py.problem6 as p6
import src_py.problem7 as p7
import src_py.problem8 as p8
import src_py.problem9 as p9
import src_py.problem10 as p10
import src_py.problem11 as p11
import src_py.problem13 as p13
import src_py.problem14 as p14
import src_py.problem15 as p15
import src_py.problem16 as p16
import src_py.problem17 as p17


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
    elif problem.day == 6:
        result = p6.run(problem, input)
    elif problem.day == 7:
        result = p7.run(problem, input)
    elif problem.day == 8:
        result = p8.run(problem, input)
    elif problem.day == 9:
        result = p9.run(problem, input)
    elif problem.day == 10:
        result = p10.run(problem, input)
    elif problem.day == 11:
        result = p11.run(problem, input)
    elif problem.day == 13:
        result = p13.run(problem, input)
    elif problem.day == 14:
        result = p14.run(problem, input)
    elif problem.day == 15:
        result = p15.run(problem, input)
    elif problem.day == 16:
        result = p16.run(problem, input)
    elif problem.day == 17:
        result = p17.run(problem, input)
    else:
        logging.error("Problem not implemented")
        exit(1)

    print(result)
