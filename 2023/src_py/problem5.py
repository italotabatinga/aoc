from enum import Enum
import logging
import sys
from src_py.problem import Problem, Part


def run(problem: Problem, input: list[str]) -> str:

    seeds = []
    if problem.part == Part.ONE:
        seeds.extend(map(int, input[0][7:].split(" ")))
    else:
        ranges = list(map(int, input[0][7:].split(" ")))
        for i in range(0, len(ranges) - 1, 2):
            for i in range(ranges[i], ranges[i] + ranges[i + 1]):
                seeds.append(i)
    logging.debug(f"seeds: {len(seeds)}")

    almanac_type = AlmanacType.SEED_TO_SOIL

    almanac = Almanac()

    for line in input[2:]:
        if line.startswith("\n"):
            continue

        if line.startswith("seed-to-soil"):
            almanac_type = AlmanacType.SEED_TO_SOIL
        elif line.startswith("soil-to-fertilizer"):
            almanac_type = AlmanacType.SOIL_TO_FERTILIZER
        elif line.startswith("fertilizer-to-water"):
            almanac_type = AlmanacType.FERTILIZER_TO_WATER
        elif line.startswith("water-to-light"):
            almanac_type = AlmanacType.WATER_TO_LIGHT
        elif line.startswith("light-to-temperature"):
            almanac_type = AlmanacType.LIGHT_TO_TEMPERATURE
        elif line.startswith("temperature-to-humidity"):
            almanac_type = AlmanacType.TEMPERATURE_TO_HUMIDITY
        elif line.startswith("humidity-to-location"):
            almanac_type = AlmanacType.HUMIDITY_TO_LOCATION
        else:
            ranges = list(map(int, line.split(" ")))
            item = AlmanacItem(ranges[0], ranges[1], ranges[2])

            almanac.add(item, almanac_type)

    almanac.build_seed_to_location_map()
    print(almanac)

    sequence = [
        AlmanacType.SEED_TO_SOIL,
        AlmanacType.SOIL_TO_FERTILIZER,
        AlmanacType.FERTILIZER_TO_WATER,
        AlmanacType.WATER_TO_LIGHT,
        AlmanacType.LIGHT_TO_TEMPERATURE,
        AlmanacType.TEMPERATURE_TO_HUMIDITY,
        AlmanacType.HUMIDITY_TO_LOCATION,
    ]

    locations = []
    for seed in seeds:
        value = seed
        for almanac_type in sequence:
            value = almanac.map(value, almanac_type)

        locations.append(value)

    return str(min(locations))


class AlmanacItem:
    def __init__(self, dst_start: int, src_start: int, length: int) -> None:
        self.dst_start = dst_start
        self.src_start = src_start
        self.length = length

    @property
    def dst_end(self) -> int:
        return self.dst_start + self.length - 1

    @property
    def src_end(self) -> int:
        return self.src_start + self.length - 1

    def map(self, src: int) -> int:
        return self.dst_start + (src - self.src_start)

    def __str__(self) -> str:
        return f"{self.src_start}..{self.src_end} maps to {self.dst_start}..{self.dst_end}"

    def __repr__(self) -> str:
        return self.__str__()

    def __contains__(self, item: int) -> bool:
        return self.src_start <= item and item < self.src_start + self.length


class AlmanacType(Enum):
    SEED_TO_SOIL = 1,
    SOIL_TO_FERTILIZER = 2,
    FERTILIZER_TO_WATER = 3,
    WATER_TO_LIGHT = 4,
    LIGHT_TO_TEMPERATURE = 5,
    TEMPERATURE_TO_HUMIDITY = 6,
    HUMIDITY_TO_LOCATION = 7,


class Almanac:
    def __init__(self) -> None:
        self.seed_to_soil_map: list[AlmanacItem] = []
        self.soil_to_fertilizer_map: list[AlmanacItem] = []
        self.fertilizer_to_water_map: list[AlmanacItem] = []
        self.water_to_light_map: list[AlmanacItem] = []
        self.light_to_temperature_map: list[AlmanacItem] = []
        self.temperature_to_humidity_map: list[AlmanacItem] = []
        self.humidity_to_location_map: list[AlmanacItem] = []
        self.humidity_to_location_map: list[AlmanacItem] = []
        self.seed_to_location_map: list[AlmanacItem] = []

    def add(self, item: AlmanacItem, type: AlmanacType) -> None:
        if type == AlmanacType.SEED_TO_SOIL:
            self.seed_to_soil_map.append(item)
        elif type == AlmanacType.SOIL_TO_FERTILIZER:
            self.soil_to_fertilizer_map.append(item)
        elif type == AlmanacType.FERTILIZER_TO_WATER:
            self.fertilizer_to_water_map.append(item)
        elif type == AlmanacType.WATER_TO_LIGHT:
            self.water_to_light_map.append(item)
        elif type == AlmanacType.LIGHT_TO_TEMPERATURE:
            self.light_to_temperature_map.append(item)
        elif type == AlmanacType.TEMPERATURE_TO_HUMIDITY:
            self.temperature_to_humidity_map.append(item)
        elif type == AlmanacType.HUMIDITY_TO_LOCATION:
            self.humidity_to_location_map.append(item)

    def map(self, value: int, type: AlmanacType) -> int:
        list = self.get_list(type)
        for item in list:
            if value in item:
                return item.map(value)

        return value

    def get_list(self, type: AlmanacType) -> list[AlmanacItem]:
        if type == AlmanacType.SEED_TO_SOIL:
            return self.seed_to_soil_map
        elif type == AlmanacType.SOIL_TO_FERTILIZER:
            return self.soil_to_fertilizer_map
        elif type == AlmanacType.FERTILIZER_TO_WATER:
            return self.fertilizer_to_water_map
        elif type == AlmanacType.WATER_TO_LIGHT:
            return self.water_to_light_map
        elif type == AlmanacType.LIGHT_TO_TEMPERATURE:
            return self.light_to_temperature_map
        elif type == AlmanacType.TEMPERATURE_TO_HUMIDITY:
            return self.temperature_to_humidity_map
        elif type == AlmanacType.HUMIDITY_TO_LOCATION:
            return self.humidity_to_location_map
        else:
            raise Exception("Invalid almanac type")

    def _fill_gaps(self, almanac: list[AlmanacItem]) -> list[AlmanacItem]:
        MIN = 0
        MAX = sys.maxsize
        almanac.sort(key=lambda x: x.src_start)

        curr = MIN
        result = []
        for item in almanac:
            if curr < item.src_start:
                result.append(AlmanacItem(curr, curr, item.src_start - curr))
            result.append(item)
            curr = item.src_end + 1
        if MAX > curr:
            result.append(AlmanacItem(curr, curr, MAX - curr))
        return result

    def merge_almanacs(self, almanac_1: list[AlmanacItem], almanac_2: list[AlmanacItem]) -> list[AlmanacItem]:
        almanac_1.sort(key=lambda x: x.dst_start)
        almanac_2.sort(key=lambda x: x.src_start)

        almanac = []
        src_start = 0
        i = j = 0

        while i < len(almanac_1) and j < len(almanac_2):
            item_1 = almanac_1[i]
            item_2 = almanac_2[j]

            if item_1.dst_end <= item_2.src_start:
                almanac.append(item_1)
                i += 1
            elif item_1.dst_start >= item_2.src_end:
                almanac.append(item_2)
                j += 1
            else:
                almanac.append(almanac_2[j])
                j += 1

        return almanac

    def build_seed_to_location_map(self) -> None:
        self.seed_to_soil_map = self._fill_gaps(self.seed_to_soil_map)
        self.soil_to_fertilizer_map = self._fill_gaps(
            self.soil_to_fertilizer_map)
        self.fertilizer_to_water_map = self._fill_gaps(
            self.fertilizer_to_water_map)
        self.water_to_light_map = self._fill_gaps(self.water_to_light_map)
        self.light_to_temperature_map = self._fill_gaps(
            self.light_to_temperature_map)
        self.temperature_to_humidity_map = self._fill_gaps(
            self.temperature_to_humidity_map)
        self.humidity_to_location_map = self._fill_gaps(
            self.humidity_to_location_map)
        self.humidity_to_location_map = self._fill_gaps(
            self.humidity_to_location_map)

    def __str__(self) -> str:
        return '\n'.join([
            f"seed_to_soil_map: {self.seed_to_soil_map}",
            f"soil_to_fertilizer_map: {self.soil_to_fertilizer_map}",
            f"fertilizer_to_water_map: {self.fertilizer_to_water_map}",
            f"water_to_light_map: {self.water_to_light_map}",
            f"light_to_temperature_map: {self.light_to_temperature_map}",
            f"temperature_to_humidity_map: {self.temperature_to_humidity_map}",
            f"humidity_to_location_map: {self.humidity_to_location_map}",
            f"humidity_to_location_map: {self.humidity_to_location_map}",
        ])
