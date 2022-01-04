use std::collections::HashMap;

use crate::Aoc;

pub struct Aoc12 {}

impl Aoc<Cave, usize> for Aoc12 {
    fn input() -> String {
        String::from(include_str!("inputs/12.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/12_test.txt"))
    }

    fn parse_input(s: String) -> Cave {
        let mut cave = HashMap::new();
        s.lines().for_each(|conn| {
            let split: Vec<&str> = conn.split('-').collect();
            let a = split[0];
            let b = split[1];
            insert_connection(&mut cave, a, b);
            insert_connection(&mut cave, b, a);
        });
        cave
    }

    fn part1(cave: Cave, _: bool) -> usize {
        let paths = explore(&cave, String::from("start"), |start, path| {
            !path.contains(start) || !is_small_cave(start)
        });
        paths.len()
    }

    fn part2(cave: Cave, _: bool) -> usize {
        let paths = explore(&cave, String::from("start"), |start, path| {
            if start == "start" || start == "end" {
                !path.contains(start)    
            } else {
                !path.contains(start) || !is_small_cave(start) || !has_two_small_caves(path)
            }
        });
        paths.len()
    }
}

type Cave = HashMap<String, Vec<String>>;
type Path = Vec<String>;

fn insert_connection(cave: &mut Cave, a: &str, b: &str) {
    match cave.get_mut(a) {
        Some(vec) => vec.push(b.to_string()),
        None => {
            cave.insert(a.to_string(), Vec::from([b.to_string()]));
        }
    }
}

fn is_small_cave(s: &String) -> bool {
    *s == s.to_lowercase()
}

fn explore(cave: &Cave, start: String, can_visit: fn(&String, &Path) -> bool) -> Vec<Path> {
    let mut paths = Vec::new();

    navigate(cave, &start, Vec::new(), &mut paths, can_visit);

    paths
}

fn navigate(
    cave: &Cave,
    start: &String,
    path: Path,
    paths: &mut Vec<Path>,
    can_visit: fn(&String, &Path) -> bool,
) {
    let mut path = path.clone();
    if can_visit(start, &path) {
        path.push(start.clone())
    } else {
        return;
    }

    if start == "end" {
        paths.push(path.clone());
    } else {
        match cave.get(start) {
            Some(caves) => {
                for cv in caves {
                    navigate(cave, cv, path.clone(), paths, can_visit);
                }
            }
            None => {}
        }
    }
}

fn has_two_small_caves(path: &Path) -> bool {
    let mut smalls: Path = path.iter().filter(|&x| is_small_cave(&x)).map(|x| x.clone()).collect();
    smalls.sort();

    for vec in smalls.windows(2) {
        let curr = &vec[0];
        let next = &vec[1];
        if *curr == *next {
            return true
        }
    }

    false
}

#[allow(dead_code)]
fn print_cave(cave: Cave) {
    println!("{:?}", cave);
}

#[allow(dead_code)]
fn print_exploration(paths: &Vec<Path>) {
    for path in paths {
        println!("{}", path.join(","))
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc12};

    #[test]
    fn aoc121_test() {
        assert_eq!(Aoc12::run(crate::Part::One, true), 226);
    }

    #[test]
    fn aoc121() {
        assert_eq!(Aoc12::run(crate::Part::One, false), 3708);
    }

    #[test]
    fn aoc122_test() {
        assert_eq!(Aoc12::run(crate::Part::Two, true), 3509);
    }

    // skipping due to its duration
    // #[test] 
    // fn aoc122() {
    //     assert_eq!(Aoc12::run(crate::Part::Two, false), 93858);
    // }
}
