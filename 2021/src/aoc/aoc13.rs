use std::{collections::HashSet};

use crate::Aoc;

pub struct Aoc13 {}

impl Aoc<(Paper, Vec<Fold>), usize> for Aoc13 {
    fn input() -> String {
        String::from(include_str!("inputs/13.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/13_test.txt"))
    }

    fn parse_input(s: String) -> (Paper, Vec<Fold>) {
        let mut paper = HashSet::new();
        let mut folds = Vec::new();

        for line in s.lines() {
            if line.starts_with("fold") {
                let mut split = line.split("=");
                split.next();
                let num = split.next().unwrap().parse::<usize>().unwrap();
                match &line[11..12] {
                    "y" => folds.push(Fold::V(num)),
                    "x" => folds.push(Fold::H(num)),
                    _ => {}
                }
            } else if line != "" {
                let mut split = line.split(",");
                let x = split.next().unwrap().parse::<usize>().unwrap();
                let y = split.next().unwrap().parse::<usize>().unwrap();
                paper.insert((x, y));
            }
        }

        (paper, folds)
    }

    fn part1((mut paper, folds): (Paper, Vec<Fold>), _: bool) -> usize {
        fold_paper(&mut paper, &folds[0]);
        paper.len()
    }

    fn part2((mut paper, folds): (Paper, Vec<Fold>), _: bool) -> usize {
        for fold in folds {
            fold_paper(&mut paper, &fold);
        }
        let bounds = find_boundaries(&paper);
        print_fmt_paper(&paper, &bounds);
        0
    }
}

type Paper = HashSet<(usize, usize)>;

fn fold_paper(paper: &mut Paper, fold: &Fold) {
    let mut folded_paper: Paper = HashSet::new();
    for pos in paper.iter() {
        match fold {
            &Fold::H(fold_x) if pos.0 > fold_x => folded_paper.insert(new_pos(pos, fold)),
            &Fold::V(fold_y) if pos.1 > fold_y => folded_paper.insert(new_pos(pos, fold)),
            _ => false,
        };
    }
    paper.extend(&folded_paper);
    paper.retain(|(x, y)| match fold {
        Fold::H(fold_x) => x < fold_x,
        Fold::V(fold_y) => y < fold_y,
    });
}

fn new_pos(pos: &(usize, usize), fold: &Fold) -> (usize, usize) {
    let new_pos = match fold {
        Fold::H(x) => (pos.0 - (pos.0 - x) * 2, pos.1),
        Fold::V(y) => (pos.0, pos.1 - (pos.1 - y) * 2),
    };
    new_pos
}

fn find_boundaries(paper: &Paper) -> (usize, usize) {
    let mut max_x = 0;
    let mut max_y = 0;

    for pos in paper {
        if max_x < pos.0 {
            max_x = pos.0;
        }
        if max_y < pos.1 {
            max_y = pos.1;
        }
    }

    (max_x, max_y)
}

#[derive(Debug)]
pub enum Fold {
    H(usize),
    V(usize),
}

#[allow(dead_code)]
fn print_paper(paper: &Paper) {
    println!("paper({}): {:?}", paper.len(), paper);
}

#[allow(dead_code)]
fn print_fmt_paper(paper: &Paper, bounds: &(usize, usize)) {
    for j in 0..bounds.1+1 {
        let mut s = String::new();
        for i in 0..bounds.0+1 {
            match paper.contains(&(i, j)) {
                true => s.push('#'),
                false => s.push('.'),
            }
        }
        println!("{}", s);
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc13};

    #[test]
    fn aoc131_test() {
        assert_eq!(Aoc13::run(crate::Part::One, true), 17);
    }

    #[test]
    fn aoc131() {
        assert_eq!(Aoc13::run(crate::Part::One, false), 693);
    }

    #[test]
    fn aoc132_test() {
        assert_eq!(Aoc13::run(crate::Part::Two, true), 0);
    }

    #[test]
    fn aoc132() {
        assert_eq!(Aoc13::run(crate::Part::Two, false), 0);
    }
}
