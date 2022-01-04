use std::{
    collections::{HashMap},
    usize::{MAX, MIN},
};

use crate::Aoc;

pub struct Aoc14 {}

impl Aoc<(Polymer, char, Rule), usize> for Aoc14 {
    fn input() -> String {
        String::from(include_str!("inputs/14.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/14_test.txt"))
    }

    fn parse_input(s: String) -> (Polymer, char, Rule) {
        let mut lines = s.lines();
        let first_line = lines.next().unwrap();
        let mut polymer: Polymer = HashMap::new();
        for i in 0..first_line.len() - 1 {
            let s = &first_line[i..i + 2];

            match polymer.get_mut(s) {
                Some(c) => {
                    *c += 1;
                }
                None => {
                    polymer.insert(s.to_string(), 1);
                }
            }
        }

        let mut rules = HashMap::new();
        lines.next(); // skip empty line
        for line in lines {
            let mut split = line.split(" -> ");
            let pair = split.next().unwrap();
            let new_char = split.next().unwrap();

            let mut pair_chars = pair.chars();
            let c1 = pair_chars.next().unwrap();
            let c2 = pair_chars.next().unwrap();
            let n1 = new_char.chars().next().unwrap();

            rules.insert(
                pair.to_string(),
                [[c1, n1].iter().collect(), [n1, c2].iter().collect()],
            );
        }
        (polymer, first_line.chars().next().unwrap(), rules)
    }

    fn part1((mut polymer, first, rules): (Polymer, char, Rule), _: bool) -> usize {
        const STEPS: usize = 10;
        for _ in 0..STEPS {
            polymer = apply_rules(polymer, &rules);
        }
        let occurrences = count_occurrences(polymer, first);
        let (min, max) = occurrences_min_max(&occurrences);
        max - min
    }

    fn part2((mut polymer, first, rules): (Polymer, char, Rule), _: bool) -> usize {
        const STEPS: usize = 40;
        for _ in 0..STEPS {
            polymer = apply_rules(polymer, &rules);
        }
        let occurrences = count_occurrences(polymer, first);
        let (min, max) = occurrences_min_max(&occurrences);
        max - min
    }
}

type Polymer = HashMap<String, usize>;
type Rule = HashMap<String, [String; 2]>;

fn apply_rules(polymer: Polymer, rules: &Rule) -> Polymer {
    let mut new_polymer: Polymer = HashMap::new();
    for (k, count) in polymer {
        match rules.get(&k) {
            Some([pair1, pair2]) => {
                match new_polymer.get_mut(pair1) {
                    Some(x) => {
                        *x += count;
                    }
                    None => {
                        new_polymer.insert(pair1.to_string(), count);
                    }
                };
                match new_polymer.get_mut(pair2) {
                    Some(x) => {
                        *x += count;
                    }
                    None => {
                        new_polymer.insert(pair2.to_string(), count);
                    }
                };
            }
            None => {
                new_polymer.insert(k, count);
            }
        }
    }
    new_polymer
}

fn count_occurrences(polymer: Polymer, first: char) -> HashMap<char, usize> {
    let mut occurrences = HashMap::new();
    occurrences.insert(first, 1);
    for (k,v) in polymer {
        let c = k.chars().nth(1).unwrap();
        match occurrences.get(&c) {
            Some(&count) => {
                occurrences.insert(c, count + v);
            }
            None => {
                occurrences.insert(c, v);
            }
        }
    }
    occurrences
}

fn occurrences_min_max(occurrences: &HashMap<char, usize>) -> (usize, usize) {
    let (mut min, mut max) = (MAX, MIN);
    for (_, &v) in occurrences {
        if min > v {
            min = v;
        }
        if max < v {
            max = v;
        }
    }
    (min, max)
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc14};

    #[test]
    fn aoc141_test() {
        assert_eq!(Aoc14::run(crate::Part::One, true), 1588);
    }

    #[test]
    fn aoc141() {
        assert_eq!(Aoc14::run(crate::Part::One, false), 2587);
    }

    #[test]
    fn aoc142_test() {
        assert_eq!(Aoc14::run(crate::Part::Two, true), 2188189693529);
    }

    #[test]
    fn aoc142() {
        assert_eq!(Aoc14::run(crate::Part::Two, false), 3318837563123);
    }
}
