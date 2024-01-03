import logging
from enum import Enum
import time
from typing import Set
from time import sleep
from src_py.problem import Problem, Part
from src_py.vec2 import Vec2


class Tile(Enum):
    EMPTY = 0
    H_SPLITTER = 1
    V_SPLITTER = 2
    FWD_MIRROR = 3
    BWD_MIRROR = 4

    def __str__(self) -> str:
        if self == Tile.EMPTY:
            return '.'
        elif self == Tile.H_SPLITTER:
            return '-'
        elif self == Tile.V_SPLITTER:
            return '|'
        elif self == Tile.FWD_MIRROR:
            return '/'
        elif self == Tile.BWD_MIRROR:
            return '\\'
        else:
            raise ValueError(f"Unknown tile {self}")


class Beam:
    def __init__(self, pos: Vec2, dir: Vec2):
        self.pos = pos
        self.dir = dir
        self.hash_str = f"({self.pos.x},{self.pos.y})_({self.dir.x},{self.dir.y})"

    def __str__(self) -> str:
        return f"B({self.pos}, {self.dir})"

    def __repr__(self) -> str:
        return self.__str__()


class Contraption:
    def __init__(self, input: list[str], beam: Beam = Beam(Vec2(0, -1), Vec2(0, 1))):
        self.layout: list[list[Tile]] = []
        self.visited: list[list[int]] = []
        self.visited_beams: Set[str] = set()
        for line in input:
            row = []
            for c in line:
                if c == '.':
                    row.append(Tile.EMPTY)
                elif c == '-':
                    row.append(Tile.H_SPLITTER)
                elif c == '|':
                    row.append(Tile.V_SPLITTER)
                elif c == '/':
                    row.append(Tile.FWD_MIRROR)
                elif c == '\\':
                    row.append(Tile.BWD_MIRROR)
                else:
                    raise ValueError(f"Unknown tile {c}")
            self.layout.append(row)
            self.visited.append([0] * len(row))
        self.beams: list[Beam] = [beam]

    def reset(self, beams: list[Beam]) -> None:
        self.beams = beams
        self.visited_beams = set()
        self.visited = [[0] * len(row) for row in self.layout]

    def move_beam(self, beam: Beam) -> list[Beam]:
        pos = beam.pos + beam.dir
        if pos.x < 0 or pos.y < 0 or pos.x >= len(self.layout[0]) or pos.y >= len(self.layout):
            return []
        tile = self.layout[pos.x][pos.y]
        # logging.debug(f"Moving {beam} to {pos} -> {tile}")
        if tile == Tile.EMPTY:
            return [Beam(pos, beam.dir)]
        elif tile == Tile.H_SPLITTER:
            if beam.dir.x == 0:
                return [Beam(pos, beam.dir)]
            return [Beam(pos, Vec2(0, 1)), Beam(pos, Vec2(0, -1))]
        elif tile == Tile.V_SPLITTER:
            if beam.dir.y == 0:
                return [Beam(pos, beam.dir)]
            return [Beam(pos, Vec2(1, 0)), Beam(pos, Vec2(-1, 0))]
        elif tile == Tile.FWD_MIRROR:  # /
            # Vec2(1, 0) -> Vec2(0, -1)
            # Vec2(0, 1) -> Vec2(-1, 0)
            # Vec2(-1, 0) -> Vec2(0, 1)
            # Vec2(0, -1) -> Vec2(-1, 0)
            return [Beam(pos, Vec2(-beam.dir.y, -beam.dir.x))]
        elif tile == Tile.BWD_MIRROR:  # \
            # Vec2(1, 0) -> Vec2(0, 1)
            # Vec2(0, 1) -> Vec2(1, 0)
            # Vec2(-1, 0) -> Vec2(0, -1)
            # Vec2(0, -1) -> Vec2(-1, 0)
            return [Beam(pos, Vec2(beam.dir.y, beam.dir.x))]
        else:
            raise ValueError(f"Unknown tile {tile}")

    def run(self, iterations: int = -1) -> int:
        no_limit_iterations = iterations <= 0
        count_it = 0
        previous_set_len = len(self.visited_beams)
        set_changed = True
        while len(self.beams) > 0 and (no_limit_iterations or count_it < iterations) and set_changed:
            new_beams = []
            for beam in self.beams:
                new_beams.extend(self.move_beam(beam))
            new_beams = list(
                filter(lambda b: b.hash_str not in self.visited_beams, new_beams))
            for beam in new_beams:
                self.visited[beam.pos.x][beam.pos.y] += 1
                self.visited_beams.add(beam.hash_str)

            set_changed = len(self.visited_beams) > previous_set_len
            previous_set_len = len(self.visited_beams)
            self.beams = new_beams
            count_it += 1

        # logging.debug(f"set_changed {set_changed} previous_set_len: {previous_set_len}")
        return count_it
        # logging.debug(f"{count_it:04}Beams: {self.beams}\nContraption:\n{self}\n")
        # sleep(0.5)

    def __str__(self) -> str:
        str_builder = []
        for row in self.layout:
            str_row = []
            for c in row:
                str_row.append(str(c))
            str_builder.append(str_row)

        for beam in self.beams:
            c = 'X'
            if beam.dir.y == 1:
                c = '>'
            elif beam.dir.y == -1:
                c = '<'
            elif beam.dir.x == 1:
                c = 'v'
            elif beam.dir.x == -1:
                c = '^'
            str_builder[beam.pos.x][beam.pos.y] = c

        return "\n".join(["".join(row) for row in str_builder])

    def energy(self) -> int:
        sum = 0
        for row in self.visited:
            for val in row:
                if val > 0:
                    sum += 1
        return sum

    def energized_string(self) -> str:
        str_builder = []
        for row in self.visited:
            str_row = []
            for v in row:
                str_row.append('#' if v > 0 else '.')
            str_builder.append(str_row)

        return "\n".join(["".join(row) for row in str_builder])


def run(problem: Problem, input: list[str]) -> str:
    sum = 0
    contraption = Contraption(input)
    if problem.part == Part.ONE:
        iterations = contraption.run()
        logging.debug(
            f"visited {list(sorted(list(contraption.visited_beams)))}")
        logging.debug(f"it: {iterations: 4} Contraption:\n{contraption}\n")
        logging.debug(f"visited:\n{contraption.energized_string()}\n")
        for row in contraption.visited:
            for val in row:
                if val > 0:
                    sum += 1
    else:
        beams = []
        for i in range(len(input[0])):
            beams.append(Beam(Vec2(-1, i), Vec2(1, 0)))
            beams.append(Beam(Vec2(len(input) - 1, i), Vec2(-1, 0)))
        for i in range(len(input)):
            beams.append(Beam(Vec2(i, -1), Vec2(0, 1)))
            beams.append(Beam(Vec2(i, len(input[0])-1), Vec2(0, -1)))
        max_beam = None
        max_energy = -1
        start_time = time.time()

        for beam in beams:
            contraption.reset([beam])
            iterations = contraption.run()
            energy = contraption.energy()
            if energy > max_energy:
                max_energy = energy
                max_beam = beam
        sum = max_energy
        logging.debug(f"Time: {time.time() - start_time}")
        logging.debug(f"Max beam: {max_beam}, max_energy: {max_energy}")
    return str(sum)
