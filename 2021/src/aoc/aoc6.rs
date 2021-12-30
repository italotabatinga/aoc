use crate::Aoc;

pub struct Aoc6 {}

impl Aoc<Vec<i64>, i64> for Aoc6 {
    fn input() -> String {
        String::from(include_str!("inputs/6.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/6_test.txt"))
    }

    fn parse_input(s: String) -> Vec<i64> {
        let mut groups = [0; 9];
        s.split(",")
            .map(|s| s.parse::<usize>().unwrap())
            .for_each(|timer| {
                groups[timer] += 1;
            });
        Vec::from(groups)
    }

    fn part1(mut groups: Vec<i64>, _: bool) -> i64 {
        let sim_days = 80;

        for _ in 0..sim_days {
            let new_growth = groups.remove(0);
            groups[6] += new_growth;
            groups.push(new_growth)
        }
        groups.iter().fold(0, |sum, v| sum + v)
    }

    fn part2(mut groups: Vec<i64>, _: bool) -> i64 {
        let sim_days = 256;
        for _ in 0..sim_days {
            let new_growth = groups.remove(0);
            groups[6] += new_growth;
            groups.push(new_growth)
        }
        groups.iter().fold(0, |sum, v| sum + v)
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc6};

    #[test]
    fn aoc61_test() {
        assert_eq!(Aoc6::run(crate::Part::One, true), 5934);
    }

    #[test]
    fn aoc61() {
        assert_eq!(Aoc6::run(crate::Part::One, false), 359344);
    }

    #[test]
    fn aoc62_test() {
        assert_eq!(Aoc6::run(crate::Part::Two, true), 26984457539);
    }

    #[test]
    fn aoc62() {
        assert_eq!(Aoc6::run(crate::Part::Two, false), 1629570219571);
    }
}
