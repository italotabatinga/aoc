import logging
from enum import Enum
import math
from functools import reduce
from src_py.problem import Problem, Part

Network = dict[str, tuple[str, str]]


def run(problem: Problem, input: list[str]) -> str:
    instructions: list[Instruction] = []
    for inst in input[0].strip():
        instructions.append(Instruction.from_string(inst))
    logging.debug(f"instructions: {instructions}")

    network: Network = {}
    for line in input[2:]:
        curr_position, nodes = line.strip().split(' = (')
        left_pos, right_pos = nodes[:-1].split(', ')
        network[curr_position] = (left_pos, right_pos)
    logging.debug(f"network: {network}")

    positions = []
    for position in network.keys():
        if problem.part == Part.ONE:
            if position == 'AAA':
                positions.append(position)
        else:
            if position.endswith('A'):
                positions.append(position)

    length = len(positions)
    min_steps = [0] * length
    logging.debug(
            f"initial positions: {positions}")

    for i in range(length):
        moves = 0
        reached = False
        curr = positions[i]
        while reached is False:
            move = instructions[moves % len(instructions)]

            next_position = network[curr][move.value]

            if problem.part == Part.ONE:
                reached = next_position == 'ZZZ'
            else:
                reached = next_position.endswith('Z')

            logging.debug(
                f"moves: {moves}, curr: {curr}, move: {move}, next: {next_position}, reached: {reached}")

            curr = next_position
            moves += 1

        min_steps[i] = moves

    logging.debug(f"min_steps: {min_steps}")

    def lcm(a: int, b: int) -> int: # least common multiple
        return abs(a*b) // math.gcd(a, b)

    total_moves = reduce(lcm, min_steps)

    return str(total_moves)


class Instruction(Enum):
    LEFT = 0
    RIGHT = 1

    @classmethod
    def from_string(cls, instruction_str: str) -> 'Instruction':
        return {
            'L': cls.LEFT,
            'R': cls.RIGHT,
        }[instruction_str.upper()]
