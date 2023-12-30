import logging
from src_py.problem import Problem, Part
from src_py.vec2 import Vec2

GALAXY = '#'
EMPTY = '.'


class Universe:
    def __init__(self, image: list[list[str]]) -> None:
        self.width = len(image[0])
        self.height = len(image)
        self.image = image
        self.rows_distance = [1] * self.height
        self.cols_distance = [1] * self.width
        self._galaxies = None

    def from_input(input: list[str]) -> 'Universe':
        return Universe(list(map(lambda line: list(line), input)))

    def expand(self, expanded_distance=1) -> None:
        empty_rows = list(range(0, self.height))
        empty_cols = list(range(0, self.width))

        for i, row in enumerate(self.image):
            for j, pixel in enumerate(row):
                if pixel == GALAXY:
                    if i in empty_rows:
                        empty_rows.remove(i)
                    if j in empty_cols:
                        empty_cols.remove(j)

        for i in empty_rows:
            self.rows_distance[i] *= expanded_distance

        for j in empty_cols:
            self.cols_distance[j] *= expanded_distance

    def distance(self, g1: Vec2, g2: Vec2) -> int:
        src_y = min(g1.y, g2.y)
        dst_y = max(g1.y, g2.y)
        src_x = min(g1.x, g2.x)
        dst_x = max(g1.x, g2.x)

        distance = 0
        for i in range(src_x, dst_x):
            distance += self.rows_distance[i]

        for i in range(src_y, dst_y):
            distance += self.cols_distance[i]

        return distance

    @property
    def galaxies(self) -> list[Vec2]:
        if self._galaxies is not None:
            return self._galaxies

        self._galaxies = []
        for i, row in enumerate(self.image):
            for j, pixel in enumerate(row):
                if pixel == GALAXY:
                    self._galaxies.append(Vec2(i, j))
        return self._galaxies

    def __str__(self) -> str:
        strings = [f"Universe ({self.height}, {self.width})",]
        for row in self.image:
            strings.append(''.join(row))
        return '\n'.join(strings)


def run(problem: Problem, input: list[str]) -> str:
    universe: Universe = Universe.from_input(input)
    universe.expand(expanded_distance=2 if problem.part ==
                    Part.ONE else 1000000)

    distances = []
    for i, gal1 in enumerate(universe.galaxies):
        for j in range(i + 1, len(universe.galaxies)):
            gal2 = universe.galaxies[j]
            distances.append(universe.distance(gal1, gal2))

    return str(sum(distances))
