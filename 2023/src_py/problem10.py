import logging
from src_py.problem import Problem, Part
from src_py.vec2 import Vec2


def connections(input: list[str], pos: Vec2) -> list[Vec2]:
    pipe = input[pos.x][pos.y]

    if pipe == 'S':
        return s_connections(input, pos)

    paths = {
        '|': [pos + Vec2(-1, 0), pos + Vec2(1, 0)],
        '-': [pos + Vec2(0, -1), pos + Vec2(0, 1)],
        'L': [pos + Vec2(-1, 0), pos + Vec2(0, 1)],
        'J': [pos + Vec2(-1, 0), pos + Vec2(0, -1)],
        '7': [pos + Vec2(0, -1), pos + Vec2(1, 0)],
        'F': [pos + Vec2(0, 1), pos + Vec2(1, 0)],
        '.': [pos + Vec2(0, 0), pos + Vec2(0, 0)],
    }[pipe]

    return list(filter(lambda p: p.x in range(0, len(input)) and p.y in range(0, len(input[pos.x])), paths))


def next_position(input: list[str], prev: Vec2, curr: Vec2) -> Vec2:
    possibilities = connections(input, curr)

    if prev == possibilities[0]:
        return possibilities[1]

    return possibilities[0]


def s_connections(input: list[str], s: Vec2) -> (str, list[Vec2]):
    directions = [Vec2(1, 0), Vec2(-1, 0),
                  Vec2(0, 1), Vec2(0, -1)]
    possibilities = list(filter(lambda p: p.x in range(0, len(
        input)) and p.y in range(0, len(input)), map(lambda d: s + d, directions)))
    possibilities = list(
        filter(lambda p: s in connections(input, p), possibilities))

    pipe = '.'
    if possibilities[0].x == s.x and possibilities[1].x == s.x:
        pipe = '-'
    elif possibilities[0].y == s.y and possibilities[1].y == s.y:
        pipe = '|'
    elif possibilities[0].x == s.x and possibilities[1].y == s.y:
        pipe = 'L'
    elif possibilities[0].x == s.x and possibilities[1].y == s.y:
        pipe = 'J'
    elif possibilities[0].y == s.y and possibilities[1].x == s.x:
        pipe = '7'
    elif possibilities[0].y == s.y and possibilities[1].x == s.x:
        pipe = 'F'

    return (pipe, possibilities)


def run(problem: Problem, input: list[str]) -> str:
    i, j = 0, 0
    for x in range(len(input)):
        line = input[x]
        if 'S' in line:
            i = x
            j = line.index('S')
            break

    s_pos = Vec2(i, j)
    logging.debug(f"S {s_pos}")
    s_pipe, s_nexts = connections(input, s_pos)
    input[i] = input[i][:j] + s_pipe + input[i][j+1:]
    logging.debug(f"input[][] {input[i][j]} s_nexts: {s_nexts}")

    loop_size = 0
    curr = s_nexts[0]
    prev = s_pos
    closed_loop = False
    loop_tiles: set[Vec2] = set([s_pos, curr])
    while not closed_loop:
        next = next_position(input, prev, curr)
        logging.debug(f"prev: {prev}, curr: {curr}, next: {next}")
        prev = curr
        curr = next
        loop_size += 1
        loop_tiles.add(curr)
        if curr == s_pos:
            closed_loop = True

    if problem.part == Part.ONE:
        return str(loop_size // 2 + 1)

    logging.debug(f"calculating...")
    loop_mapping = []
    for i in range(len(input)):
        loop_crossings = 0
        from_up = False
        from_bottom = False
        for j in range(len(input[i])):
            if Vec2(i, j) in loop_tiles:
                loop_mapping.append(0)
                if input[i][j] == '|':
                    loop_crossings += 1
                elif input[i][j] == 'L':
                    from_up = True
                elif input[i][j] == 'F':
                    from_bottom = True
                elif input[i][j] == 'J':
                    if from_bottom:
                        loop_crossings += 1
                    from_up = from_bottom = False
                elif input[i][j] == '7':
                    if from_up:
                        loop_crossings += 1
                    from_up = from_bottom = False

            else:
                loop_mapping.append(loop_crossings)

    enclosed_tiles = 0
    for tile in loop_mapping:
        if tile % 2 == 1:
            enclosed_tiles += 1
    return str(enclosed_tiles)
