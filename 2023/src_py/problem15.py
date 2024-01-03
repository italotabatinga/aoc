import logging
from src_py.problem import Problem, Part

def hash(s: str) -> int:
    num = 0
    for char in s:
        num += ord(char)
        num *= 17
        num %= 256
    return num

class Lens:
    def __init__(self, label: str, focal_length: int) -> None:
        self.label = label
        self.focal_length = focal_length

    def __str__(self) -> str:
        return f"({self.label} {self.focal_length})"

    def __repr__(self) -> str:
        return self.__str__()

class Hashmap:
    def __init__(self) -> None:
        self.boxes: list[list[Lens]] = [[] for i in range(0, 256)]

    def add(self, lens: Lens) -> None:
        hash_value = hash(lens.label)
        box = self.boxes[hash_value]

        for elem in box:
            if elem.label == lens.label:
                elem.focal_length = lens.focal_length
                return

        box.append(lens)

    def remove(self, label: str) -> None:
        hash_value = hash(label)
        box = self.boxes[hash_value]
        found_index = -1
        for i, elem in enumerate(box):
            if elem.label == label:
                found_index = i
                break

        if found_index >= 0:
            box.pop(found_index)

    def focusing_power(self) -> int:
        value = 0
        for i, box in enumerate(self.boxes):
            for j, lens in enumerate(box):
                value += (i + 1) * (j + 1) * lens.focal_length
        return value

    def __str__(self) -> str:
        str_builder = []
        for i, box in enumerate(self.boxes):
            if len(box) == 0:
                continue
            str_builder.append(f"Box {i}: {box}")

        return "\n".join(str_builder)

def run(problem: Problem, input: list[str]) -> str:
    sum = 0
    commands = input[0].split(',')
    if problem.part == Part.ONE:
        for cmd in commands:
            sum += hash(cmd)
    else:
        hashmap = Hashmap()
        for cmd in commands:
            has_equal = cmd.find("=") > 0
            if has_equal:
                label, value_str = cmd.split("=")
                hashmap.add(Lens(label, int(value_str)))
            else:
                label = cmd[:-1]
                hashmap.remove(label)

            logging.debug(f"After \"{cmd}\":\n{hashmap}\n")
        sum += hashmap.focusing_power()


    return str(sum)
