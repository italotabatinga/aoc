use std::collections::{HashMap, HashSet};

use crate::Aoc;

pub struct Aoc8 {}

impl Aoc<Vec<Entry>, i32> for Aoc8 {
    fn input() -> String {
        String::from(include_str!("inputs/8.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/8_test.txt"))
    }

    fn parse_input(s: String) -> Vec<Entry> {
        s.split("\n")
            .map(|str_entry| {
                let parts: Vec<&str> = str_entry.split(" | ").collect();
                Entry::new(
                    parts[0]
                        .split_whitespace()
                        .map(|signals| Digit::new(String::from(signals)))
                        .collect(),
                    parts[1]
                        .split_whitespace()
                        .map(|signals| Digit::new(String::from(signals)))
                        .collect(),
                )
            })
            .collect()
    }

    fn part1(input: Vec<Entry>, _: bool) -> i32 {
        input.iter().fold(0, |count, entry| {
            count
                + entry
                    .output
                    .iter()
                    .fold(0, |sub_count, output| match output.signals.len() {
                        2 | 3 | 4 | 7 => sub_count + 1,
                        _ => sub_count,
                    })
        })
    }

    fn part2(input: Vec<Entry>, _: bool) -> i32 {
        let mut sum = 0;
        for mut entry in input {
            entry.build_mapper();
            sum += entry.num()
        }
        sum
    }
}

#[derive(Debug)]
pub struct Entry {
    patterns: Vec<Digit>,
    output: Vec<Digit>,
    pub mapper: HashMap<String, i32>,
}

impl Entry {
    pub fn new(patterns: Vec<Digit>, output: Vec<Digit>) -> Self {
        Self {
            patterns,
            output,
            mapper: HashMap::new(),
        }
    }

    pub fn build_mapper(&mut self) {
        let one = self.patterns.iter().find(|&digit| digit.len == 2).unwrap();
        let seven = self.patterns.iter().find(|&digit| digit.len == 3).unwrap();
        let four = self.patterns.iter().find(|&digit| digit.len == 4).unwrap();
        let eight = self.patterns.iter().find(|&digit| digit.len == 7).unwrap();
        let a = one.sym_diff(seven);

        let adg = self
            .patterns
            .iter()
            .filter(|&digit| digit.len == 5)
            .fold(None::<Digit>, |option, digit| match option {
                Some(intersec) => Some(intersec.intersection(digit)),
                None => Some(digit.clone()),
            })
            .unwrap();
        let d = adg.intersection(&four);
        let g = adg.sym_diff(&d).sym_diff(&a);
        let b = four.sym_diff(&seven).sym_diff(&a).sym_diff(&d);
        let e = seven
            .sym_diff(&eight)
            .sym_diff(&b)
            .sym_diff(&d)
            .sym_diff(&g);
        let abgf = self
            .patterns
            .iter()
            .filter(|&digit| digit.len == 6)
            .fold(None::<Digit>, |option, digit| match option {
                Some(intersec) => Some(intersec.intersection(digit)),
                None => Some(digit.clone()),
            })
            .unwrap();
        let f = abgf.sym_diff(&a).sym_diff(&b).sym_diff(&g);
        let c = one.sym_diff(&f);

        // println!(
        //     "a: {:?}, b: {:?}, c: {:?}, d: {:?}, e: {:?}, f: {:?}, g: {:?}",
        //     a.first(),
        //     b.first(),
        //     c.first(),
        //     d.first(),
        //     e.first(),
        //     f.first(),
        //     g.first()
        // );
        self.mapper = HashMap::from([
            (eight.sym_diff(&d).code(), 0),
            (one.code(), 1),
            (eight.sym_diff(&b).sym_diff(&f).code(), 2),
            (eight.sym_diff(&b).sym_diff(&e).code(), 3),
            (four.code(), 4),
            (eight.sym_diff(&c).sym_diff(&e).code(), 5),
            (eight.sym_diff(&c).code(), 6),
            (seven.code(), 7),
            (eight.code(), 8),
            (eight.sym_diff(&e).code(), 9),
        ]);
    }

    pub fn num(&self) -> i32 {
        self.output.iter().enumerate().fold(0, |acc, (i, digit)| {
            let mult = 10_i32.pow((self.output.len() - i - 1) as u32);
            acc + self.mapper[&digit.code()] * mult
        })
    }
}

#[derive(Clone, Debug)]
pub struct Digit {
    pub signals: HashSet<String>,
    pub len: usize,
}

impl Digit {
    pub fn new(signals: String) -> Self {
        Self {
            len: signals.len(),
            signals: HashSet::from_iter(signals.chars().map(|s| String::from(s))),
        }
    }

    pub fn first(&self) -> String {
        self.signals.iter().next().unwrap().clone()
    }

    pub fn sym_diff(&self, rhs: &Digit) -> Self {
        Digit::new(
            self.signals
                .symmetric_difference(&rhs.signals)
                .into_iter()
                .fold(String::new(), |signals, signal| signals + signal),
        )
    }

    pub fn intersection(&self, rhs: &Digit) -> Self {
        Digit::new(
            self.signals
                .intersection(&rhs.signals)
                .into_iter()
                .fold(String::new(), |signals, signal| signals + signal),
        )
    }

    pub fn union(&self, rhs: &Digit) -> Self {
        Digit::new(
            self.signals
                .union(&rhs.signals)
                .into_iter()
                .fold(String::new(), |signals, signal| signals + signal),
        )
    }

    pub fn code(&self) -> String {
        let mut chars: Vec<&String> = self.signals.iter().collect();
        chars.sort();
        chars.iter().fold(String::new(), |code, s| code + s)
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc8};

    #[test]
    fn aoc81_test() {
        assert_eq!(Aoc8::run(crate::Part::One, true), 26);
    }

    #[test]
    fn aoc81() {
        assert_eq!(Aoc8::run(crate::Part::One, false), 449);
    }

    #[test]
    fn aoc82_test() {
        assert_eq!(Aoc8::run(crate::Part::Two, true), 61229);
    }

    #[test]
    fn aoc82() {
        assert_eq!(Aoc8::run(crate::Part::Two, false), 968175);
    }
}
