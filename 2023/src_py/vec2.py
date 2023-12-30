class Vec2:
    def __init__(self, *args) -> None:
        if len(args) == 1:
            self.x, self.y = args[0]
        elif len(args) == 2:
            self.x, self.y = args

    def __str__(self) -> str:
        return f"({self.x}, {self.y})"

    def __repr__(self) -> str:
        return self.__str__()

    def __eq__(self, other: object) -> bool:
        if not isinstance(other, Vec2):
            return NotImplemented
        return self.x == other.x and self.y == other.y

    def __add__(self, other: 'Vec2') -> 'Vec2':
        return Vec2(self.x + other.x, self.y + other.y)

    def __hash__(self) -> int:
        return hash((self.x, self.y))