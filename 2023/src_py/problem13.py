import logging
from src_py.problem import Problem, Part


def run(problem: Problem, input: list[str]) -> str:
    count = 0
    patterns = []
    pattern = []

    for line in input:
        if line == '':
            patterns.append(pattern)
            pattern = []
            continue

        pattern.append(line)
    if len(pattern) > 0:
        patterns.append(pattern)
        pattern = []

    smudge_tolerance = 0 if problem.part == Part.ONE else 1
    for pattern in patterns:
        logging.debug(f"pattern {pattern}")
        v_syms = set(range(0, len(pattern[0]) - 1))
        h_syms = set(range(0, len(pattern) - 1))

        for v_sym in list(v_syms):
            smudges_count = 0
            for row in pattern:
                min_to_edge = min(v_sym + 1 - 0, len(row) - 1 - v_sym)
                for i in range(0, min_to_edge):
                    if row[v_sym - i] != row[v_sym + i + 1]:
                        smudges_count += 1
                        if smudges_count > smudge_tolerance:
                            v_syms.remove(v_sym)
                            break
                if v_sym not in v_syms:
                    break
            if smudges_count < smudge_tolerance:
                v_syms.remove(v_sym)

        for h_sym in list(h_syms):
            smudges_count = 0
            for j in range(len(pattern[0])):
                min_to_edge = min(h_sym + 1 - 0, len(pattern) - 1 - h_sym)
                for i in range(0, min_to_edge):
                    if pattern[h_sym - i][j] != pattern[h_sym + i + 1][j]:
                        smudges_count += 1
                        if smudges_count > smudge_tolerance:
                            h_syms.remove(h_sym)
                            break
                if h_sym not in h_syms:
                    break
            if smudges_count < smudge_tolerance:
                h_syms.remove(h_sym)

        logging.debug(f"v_syms {v_syms}, h_syms {h_syms}")
        if len(v_syms) > 0:
            count += v_syms.pop() + 1
        if len(h_syms) > 0:
            count += (h_syms.pop() + 1) * 100

    return str(count)
