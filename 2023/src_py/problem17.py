import heapq
import logging
import sys
import time
from src_py.problem import Problem, Part
from src_py.vec2 import Vec2


class HeatLossMap:
    def __init__(self, input: list[str]):
        self.grid = []
        for row in input:
            self.grid.append(list(map(int, list(row))))

    def __str__(self) -> str:
        return "\n".join([str(row) for row in self.grid])

    def calc_min_path(self, start: Vec2, end: Vec2) -> int:
        frontier: list[Vec2] = []
        heapq.heappush(frontier, (0, start))
        came_from: list[list[list[Vec2]]] = [[[] for _ in range(
            len(self.grid[0]))] for _ in range(len(self.grid))]
        cost_so_far = [[sys.maxsize for _ in range(
            len(self.grid[0]))] for _ in range(len(self.grid))]
        cost_so_far[start.y][start.x] = 0

        logging.debug(f"cost_so_far: {cost_so_far}")
        iterations = 0
        while len(frontier) > 0:
            iterations += 1
            _, current = heapq.heappop(frontier)
            logging.debug(f"visiting current: {current}, frontier: {frontier}")
            if current == end:
                logging.debug(f"\t reached end: {end}")
                break

            movements = self.possible_movements(
                current, came_from[current.y][current.x])
            logging.debug(f"\t possible movements: {movements}")
            if logging.getLogger().isEnabledFor(logging.DEBUG):
                time.sleep(0.5)

            for movement in movements:
                next_mov = movement + current
                new_cost = cost_so_far[current.y][current.x] + \
                    self.grid[next_mov.y][next_mov.x]
                logging.debug(
                    f"\t next: {next_mov}, new_cost: {new_cost}, cost_so_far_next: {cost_so_far[next_mov.y][next_mov.x]}")
                if cost_so_far[next_mov.y][next_mov.x] == sys.maxsize or new_cost < cost_so_far[current.y][current.x]:
                    cost_so_far[next_mov.y][next_mov.x] = new_cost
                    heapq.heappush(frontier, (new_cost, next_mov))
                    came_from[next_mov.y][next_mov.x] = came_from[current.y][current.x] + [current]
                    # logging.debug(
                    #     f"\t setting cost of {next_mov} to {new_cost}")

                    # logging.debug(
                    #     f"\t came_from2@\n{came_from[current.y][current.x]}, {came_from[next_mov.y][next_mov.x + 1]}")
                    # logging.debug(f"\t came_from\n{came_from}\n")
                    # logging.debug(f"\t came_from")
                    if next_mov == end:
                        breakpoint()
                    for i, row in enumerate(came_from):
                        for j, col in enumerate(row):
                            if len(col) > 0:
                                logging.debug(f"\t\t({j: 2},{i: 2}): {col}")
                    if len(came_from[current.y][current.x]) > 0 and (came_from[current.y][current.x][-1] - current).sq_mag() > 1:
                        raise Exception(
                            f'Error on camefrom {came_from[current.y][current.x]} and current {current} and diff {(came_from[current.y][current.x][-1] - current).sq_mag()}')

        logging.info(f"iterations: {iterations}")
        logging.info(f"cost_so_far: {cost_so_far}")
        logging.info(f"came_from[]: {came_from[end.y][end.x]}")

        # print the path from came_from using the grid
        for row in range(len(self.grid)):
            for col in range(len(self.grid[0])):
                if Vec2(col, row) == start:
                    print("S", end="")
                elif Vec2(col, row) in came_from[end.y][end.x]:
                    i = came_from[end.y][end.x].index(Vec2(col, row))
                    diff = came_from[end.y][end.x][i] - \
                        came_from[end.y][end.x][i-1]
                    if diff == Vec2(1, 0):
                        print(">", end="")
                    elif diff == Vec2(-1, 0):
                        print("<", end="")
                    elif diff == Vec2(0, 1):
                        print("v", end="")
                    elif diff == Vec2(0, -1):
                        print("^", end="")
                    else:
                        print("X", end="")
                else:
                    print(self.grid[row][col], end="")
            print()

        return cost_so_far[end.y][end.x]

    def possible_movements(self, pos: Vec2, tail: list[Vec2] = []) -> list[Vec2]:
        movements = []
        if pos.x > 0:
            movements.append(Vec2(-1, 0))
        if pos.x < len(self.grid[0]) - 1:
            movements.append(Vec2(1, 0))
        if pos.y > 0:
            movements.append(Vec2(0, -1))
        if pos.y < len(self.grid) - 1:
            movements.append(Vec2(0, 1))

        if len(tail) > 0:
            last_move_opposite = (pos - tail[-1]) * -1
            logging.debug(f"\t\tlast_move_opposite: {last_move_opposite}")
            if last_move_opposite in movements:
                movements.remove(last_move_opposite)

        if len(tail) >= 3:
            dir1 = pos - tail[-1]
            dir2 = tail[-1] - tail[-2]
            dir3 = tail[-2] - tail[-3]
            if dir1 == dir2 and dir2 == dir3 and dir1 in movements:
                movements.remove(dir1)

        return movements


def run(problem: Problem, input: list[str]) -> str:
    sum = 0
    heat_loss = HeatLossMap(input)
    logging.debug(f"heat_loss:\n{heat_loss}\n")
    sum = heat_loss.calc_min_path(
        Vec2(0, 0), Vec2(len(input[0]) - 1, len(input) - 1))

    return str(sum)
