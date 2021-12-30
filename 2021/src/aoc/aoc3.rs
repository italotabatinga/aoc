use crate::Aoc;

pub struct Aoc3 {}

impl Aoc<Vec<u16>, u32> for Aoc3 {
    fn input() -> String {
        String::from(include_str!("inputs/3.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/3_test.txt"))
    }

    fn part1(input: Vec<u16>, is_test: bool) -> u32 {
        let max = if is_test { 5 } else { 12 };
        let gamma = (0..max)
            .into_iter()
            .fold(0, |acc, pos| acc + (find_most_frequent(&input, pos) << pos));
        let epsilon = !gamma & ((1 << max) - 1);

        gamma as u32 * epsilon as u32
    }

    fn part2(input: Vec<u16>, is_test: bool) -> u32 {
        let max = if is_test { 5 } else { 12 };

        let oxygen = find_by_criteria(&input, max - 1, |a, pos, arr| {
            let cmp = find_most_frequent(arr, pos) << pos;
            a & (1 << pos) == cmp
        });

        let co2 = find_by_criteria(&input, max - 1, |a, pos, arr| {
            let cmp = (!find_most_frequent(arr, pos) & 1) << pos;
            a & 1 << pos == cmp & 1 << pos
        });

        oxygen as u32 * co2 as u32
    }

    fn parse_input(s: String) -> Vec<u16> {
        s.split("\n")
            .map(|s| u16::from_str_radix(s.trim(), 2).unwrap())
            .collect()
    }
}

fn find_most_frequent(arr: &Vec<u16>, pos: u16) -> u16 {
    let (count0, count1) =
        arr.iter()
            .map(|bin| (bin & (1 << pos)) >> pos)
            .fold((0, 0), |mut map, bit| {
                match bit {
                    0 => map.0 += 1,
                    1 => map.1 += 1,
                    _ => {}
                };
                map
            });
    if count0 > count1 {
        0
    } else {
        1
    }
}

fn find_by_criteria(arr: &Vec<u16>, pos: u16, criteria: fn(u16, u16, &Vec<u16>) -> bool) -> u16 {
    let filtered: Vec<u16> = arr
        .iter()
        .filter(|&&bin| criteria(bin, pos, &arr))
        .map(|&x| x)
        .collect();

    if filtered.len() == 1 {
        return filtered[0];
    }

    find_by_criteria(&filtered, pos - 1, criteria)
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc3};

    #[test]
    fn aoc31_test() {
        assert_eq!(Aoc3::run(crate::Part::One, true), 198);
    }

    #[test]
    fn aoc31() {
        assert_eq!(Aoc3::run(crate::Part::One, false), 3009600);
    }

    #[test]
    fn aoc32_test() {
        assert_eq!(Aoc3::run(crate::Part::Two, true), 230);
    }

    #[test]
    fn aoc32() {
        assert_eq!(Aoc3::run(crate::Part::Two, false), 6940518);
    }
}
