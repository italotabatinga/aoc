import logging
from enum import Enum
from typing import Dict
from src_py.vec2 import Vec2
from src_py.problem import Problem, Part

ROUNDED_ROCK = 'O'
EMPTY_SPACE = '.'
FIXED_ROCK = '#'


class Direction(Enum):
    UP = 1
    DOWN = 2
    LEFT = 3
    RIGHT = 4


class Platform:
    def __init__(self, platform: list[list[str]]) -> None:
        self._platform = platform
        self._cycle_cache: Dict[str, (int, list[list[str]])] = {}

    @classmethod
    def from_input(cls, input: list[str]) -> 'Platform':
        return Platform(list(map(lambda line: list(line), input)))

    def cycle(self, n: int = 0) -> None:
        cached = self.get_cycle_from_cached()
        if cached is not None:
            self._platform = cached[1]
            return

        cache_key = str(self)
        self.tilt(Direction.UP)
        self.tilt(Direction.LEFT)
        self.tilt(Direction.DOWN)
        self.tilt(Direction.RIGHT)
        self.cache_platform(cache_key, n)

    def cycle_n(self, n: int) -> None:
        self._cycle_cache.clear()
        next_cached_cycle = None
        count = 0
        while next_cached_cycle is None and count < n:
            count += 1
            self.cycle(count)
            next_cached_cycle = self.get_cycle_from_cached()

        logging.debug(f"Found cycle at {count} from {n}, cached {len(self._cycle_cache)}")
        logging.debug(f"\tnext cycle index {next_cached_cycle[0]}")
        remaining_cycles = (n - count) % (count - next_cached_cycle[0] + 1)
        for i in range(0, remaining_cycles):
            self.cycle()

    def cache_platform(self, key: str, it: int) -> None:
        self._cycle_cache[key] = (it, self._platform)

    def get_cycle_from_cached(self) -> (int, list[list[str]]):
        return self._cycle_cache.get(str(self))

    def tilt(self, dir: Direction) -> None:
        new_platform = self._platform.copy()
        for i in range(0, len(self._platform)):
            new_platform[i] = self._platform[i].copy()
        self._platform = new_platform

        mov_dir, init_positions = self._calc_tilt_initial_pos(dir)
        for pos in init_positions:
            first_empty_space = Vec2(-1, -1)
            empty_trail_len = 0
            curr_pos = pos
            while curr_pos.x >= 0 and curr_pos.x < len(self._platform) and curr_pos.y >= 0 and curr_pos.y < len(self._platform[0]):
                curr = self._platform[curr_pos.x][curr_pos.y]
                if curr == EMPTY_SPACE:
                    if first_empty_space.x < 0:
                        first_empty_space = curr_pos
                    empty_trail_len += 1
                elif curr == FIXED_ROCK:
                    if first_empty_space.x >= 0:
                        first_empty_space = Vec2(-1, -1)
                        empty_trail_len = 0
                elif curr == ROUNDED_ROCK:
                    if first_empty_space.x >= 0:
                        self._platform[first_empty_space.x][first_empty_space.y] = ROUNDED_ROCK
                        self._platform[curr_pos.x][curr_pos.y] = EMPTY_SPACE
                        if empty_trail_len > 1:
                            first_empty_space = first_empty_space + mov_dir
                        else:
                            first_empty_space = curr_pos

                curr_pos += mov_dir

    def _calc_tilt_initial_pos(self, dir: Direction) -> (Vec2, list[Vec2]):
        if dir == Direction.UP:
            return (Vec2(1, 0), list(map(lambda j: Vec2(0, j), range(0, len(self._platform[0])))))
        elif dir == Direction.DOWN:
            return (Vec2(-1, 0), list(map(lambda j: Vec2(len(self._platform) - 1, j), range(0, len(self._platform[0])))))
        elif dir == Direction.LEFT:
            return (Vec2(0, 1), list(map(lambda i: Vec2(i, 0), range(0, len(self._platform)))))
        elif dir == Direction.RIGHT:
            return (Vec2(0, -1), list(map(lambda i: Vec2(i, len(self._platform[0]) - 1), range(0, len(self._platform)))))

    def load_north(self):
        total_load = 0
        for j in range(0, len(self._platform[0])):
            i = 0
            while i < len(self._platform):
                curr = self._platform[i][j]
                if curr == ROUNDED_ROCK:
                    total_load += len(self._platform) - i
                i += 1
        return total_load

    def __str__(self) -> str:
        return '\n'.join(map(lambda row: ''.join(row), self._platform))


def run(problem: Problem, input: list[str]) -> str:
    count = 0
    platform = Platform.from_input(input)

    logging.debug(f"platform before tilt\n{platform}")
    if problem.part == Part.TWO:
       platform.cycle_n(1000000000)
    else:
        platform.tilt(Direction.UP)
    logging.debug(f"final platform\n{platform}")

    load_north = platform.load_north()
    return str(load_north)
