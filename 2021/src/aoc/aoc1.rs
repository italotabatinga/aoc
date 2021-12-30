use super::util::Aoc;

pub struct Aoc1 {}

impl Aoc<Vec<i32>, i32> for Aoc1 {
    fn input() -> String {
        String::from(include_str!("inputs/1.txt"))
    }
    
    fn part1_test_input() -> String {
        String::from(include_str!("inputs/1_test1.txt"))
    }
    
    fn part2_test_input() -> String {
        String::from(include_str!("inputs/1_test2.txt"))
    }

    fn part1(input: Vec<i32>, _: bool) -> i32 {
        let measurements = input;

        let mut prev = measurements.get(0).unwrap();
        let mut count = 0;

        for measurement in measurements.iter() {
            if measurement > prev {
                count += 1;
            }
            prev = measurement;
        }

        count
    }

    fn part2(input: Vec<i32>, _: bool) -> i32 {
        let measurements = input;
        let mut prev = measurements[0] + measurements[1] + measurements[2];
        let mut count = 0;

        for w in measurements.windows(3) {
            let measurement = w[0] + w[1] + w[2];
            if measurement > prev {
                count += 1;
            }
            prev = measurement;
        }

        count
    }

    fn parse_input(s: String) -> Vec<i32> {
        s.split("\n")
            .map(|s| s.trim().parse::<i32>().unwrap())
            .collect()
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc1};

    #[test]
    fn aoc11_test() {
        assert_eq!(Aoc1::run(crate::Part::One, true), 7);
    }

    #[test]
    fn aoc11() {
        assert_eq!(Aoc1::run(crate::Part::One, false), 1655);
    }

    #[test]
    fn aoc12_test() {
        assert_eq!(Aoc1::run(crate::Part::Two, true), 5);
    }

    #[test]
    fn aoc12() {
        assert_eq!(Aoc1::run(crate::Part::Two, false), 1683);
    }
}
