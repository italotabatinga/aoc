use std::collections::HashMap;

use crate::Aoc;

pub struct Aoc7 {}

impl Aoc<Vec<i64>, i64> for Aoc7 {
    fn input() -> String {
        String::from(include_str!("inputs/7.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/7_test.txt"))
    }

    fn parse_input(s: String) -> Vec<i64> {
        let mut v: Vec<i64> = s.split(",").map(|s| s.parse::<i64>().unwrap()).collect();
        v.sort();
        v
    }

    fn part1(crabs: Vec<i64>, _: bool) -> i64 {
        let mut pass: HashMap<i64, bool> = HashMap::new();
        let mut derivatives: HashMap<i64, i64> = HashMap::new();
        let len = crabs.len();

        // +1 pass
        let mut pos_der = 0;
        for i in 0..len {
            let crab = crabs[i];
            if !pass.contains_key(&crab) {
                pass.insert(crab, true);
                derivatives.insert(crab, pos_der);
            }
            pos_der += 1;
        }

        // -1 pass
        pass.clear();
        let mut pos_der = 0;
        for i in 0..len {
            let crab = crabs[len - 1 - i];
            if !pass.contains_key(&crab) {
                pass.insert(crab, true);
                match derivatives.get_mut(&crab) {
                    Some(v) => *v -= pos_der,
                    _ => {}
                }
            }
            pos_der += 1;
        }

        let mut min: i64 = i64::MAX;
        let mut pos: i64 = -1;
        for (k, v) in derivatives {
            if v.abs() < min {
                min = v.abs();
                pos = k;
            }
        }

        crabs.iter().fold(0, |acc, v| acc + (v - pos).abs())
    }

    fn part2(crabs: Vec<i64>, _: bool) -> i64 {
        fn cost(b: i64, e: i64) -> i64 {
            (b - e).abs() * ((b - e).abs() + 1) / 2
        }

        let grouped_crabs: HashMap<i64, i64> =
            crabs.iter().fold(HashMap::new(), |mut map, crab| {
                match map.get_mut(&crab) {
                    Some(v) => {
                        *v += 1;
                    }
                    None => {
                        map.insert(*crab, 1);
                    }
                };
                map
            });
        let pos_range = (crabs[0])..(crabs[crabs.len() - 1]);
        let mut min: i64 = i64::MAX;
        for end in pos_range {
            let mut sum = 0;
            for (begin, count) in grouped_crabs.iter() {
                sum += cost(*begin, end) * count
            }

            if sum < min {
                min = sum;
            }
        }

        min
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc7};

    #[test]
    fn aoc71_test() {
        assert_eq!(Aoc7::run(crate::Part::One, true), 37);
    }

    #[test]
    fn aoc71() {
        assert_eq!(Aoc7::run(crate::Part::One, false), 359648);
    }

    #[test]
    fn aoc72_test() {
        assert_eq!(Aoc7::run(crate::Part::Two, true), 168);
    }

    #[test]
    fn aoc72() {
        assert_eq!(Aoc7::run(crate::Part::Two, false), 100727924);
    }
}
