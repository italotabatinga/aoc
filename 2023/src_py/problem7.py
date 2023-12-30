import logging
from enum import Enum
from functools import cmp_to_key
from src_py.problem import Problem, Part


def _compare_games(game1: 'Game', game2: 'Game') -> int:
    return game1.hand.compare(game2.hand)


def run(problem: Problem, input: list[str]) -> str:
    games = []
    for line in input:
        hand_str, bid_str = line.split()
        cards = list(map(lambda x: Card.from_string(
            x, problem.part == Part.TWO), list(hand_str)))
        games.append(Game(Hand(cards), int(bid_str)))
    games.sort(key=cmp_to_key(_compare_games))
    logging.debug(f"games: {games}")

    total_winnings = 0
    for i, game in enumerate(games):
        total_winnings += game.bid * (i + 1)

    return str(total_winnings)


class Card(Enum):
    A = 14
    K = 13
    Q = 12
    J = 11
    T = 10
    NINE = 9
    EIGHT = 8
    SEVEN = 7
    SIX = 6
    FIVE = 5
    FOUR = 4
    THREE = 3
    TWO = 2
    JOKER = 1

    @classmethod
    def from_string(cls, card_str: str, joker: bool = False) -> 'Card':
        upper = card_str.upper()
        if upper == 'J':
            return cls.JOKER if joker else cls.J

        return {
            'A': cls.A,
            'K': cls.K,
            'Q': cls.Q,
            'T': cls.T,
            '9': cls.NINE,
            '8': cls.EIGHT,
            '7': cls.SEVEN,
            '6': cls.SIX,
            '5': cls.FIVE,
            '4': cls.FOUR,
            '3': cls.THREE,
            '2': cls.TWO,
        }[upper]

    def __str__(self) -> str:
        return self.name

    def __repr__(self) -> str:
        return self.__str__()


class HandType(Enum):
    FIVE_OF_A_KIND = 6
    FOUR_OF_A_KIND = 5
    FULL_HOUSE = 4
    THREE_OF_A_KIND = 3
    TWO_PAIR = 2
    ONE_PAIR = 1
    HIGH_CARD = 0


class Hand:
    def __init__(self, cards: list[int]) -> None:
        self.cards = cards

        self._map = [0] * 15
        for card in cards:
            self._map[card.value] += 1

        self._type = None

    def __str__(self) -> str:
        return str(self.cards)

    def __repr__(self) -> str:
        return self.__str__()

    def _calculate_type(self) -> HandType:
        counters = [0] * 6
        counters[0] = self._map[Card.JOKER.value]
        _map = self._map.copy()
        counters[0] = _map.pop(Card.JOKER.value)
        _map.sort(reverse=True)
        # Given the order of the hand types, we can assume that the best strategy
        # is to put all jokers as the card with higher frequency
        _map[0] += counters[0]

        for count in _map:
            if count == 1:
                counters[1] += 1
            if count == 2:
                counters[2] += 1
            elif count == 3:
                counters[3] += 1
            elif count == 4:
                counters[4] += 1
            elif count == 5:
                counters[5] += 1

        if counters[5] == 1:
            return HandType.FIVE_OF_A_KIND
        elif counters[4] == 1:
            return HandType.FOUR_OF_A_KIND
        elif counters[3] == 1 and counters[2] == 1:
            return HandType.FULL_HOUSE
        elif counters[3] == 1:
            return HandType.THREE_OF_A_KIND
        elif counters[2] == 2:
            return HandType.TWO_PAIR
        elif counters[2] == 1:
            return HandType.ONE_PAIR

        return HandType.HIGH_CARD

    @property
    def type(self) -> HandType:
        if self._type is None:
            self._type = self._calculate_type()
        return self._type

    def compare(self, other: 'Hand') -> int:
        if self.type.value > other.type.value:
            return 1
        elif self.type.value < other.type.value:
            return -1

        for card_1, card_2 in zip(self.cards, other.cards):
            if card_1.value > card_2.value:
                return 1
            elif card_1.value < card_2.value:
                return -1

        return 0


class Game:
    def __init__(self, hand: Hand, bid: int) -> None:
        self.hand = hand
        self.bid = bid

    def __str__(self) -> str:
        return f"{{{self.hand}, {self.bid}}}"

    def __repr__(self) -> str:
        return self.__str__()
