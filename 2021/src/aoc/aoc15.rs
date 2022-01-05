use std::collections::BinaryHeap;

use crate::Aoc;

pub struct Aoc15 {}

impl Aoc<(Cavern, Limit), usize> for Aoc15 {
    fn input() -> String {
        String::from(include_str!("inputs/15.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/15_test.txt"))
    }

    fn parse_input(s: String) -> (Cavern, Limit) {
        let mut cavern: Cavern = Vec::new();
        for line in s.lines() {
            cavern.push(
                line.chars()
                    .map(|c| c.to_digit(10).unwrap() as usize)
                    .collect(),
            );
        }
        let limit = (cavern.len(), cavern[0].len());
        (cavern, limit)
    }

    fn part1((cavern, _): (Cavern, Limit), _: bool) -> usize {
        let limits = (cavern.len(), cavern[0].len());
        let mut dijsktra = new_dijkstra_cavern(&cavern);

        fill_dijkstra_cavern(&cavern, &mut dijsktra);
        print_cavern(&dijsktra);
        dijsktra[limits.0 - 1][limits.1 - 1]
    }

    fn part2((cavern, _): (Cavern, Limit), _: bool) -> usize {
        let cavern = calc_whole_cavern(&cavern);
        let limits = (cavern.len(), cavern[0].len());
        let mut dijsktra = new_dijkstra_cavern(&cavern);

        fill_dijkstra_cavern(&cavern, &mut dijsktra);
        dijsktra[limits.0 - 1][limits.1 - 1]
    }
}

type Cavern = Vec<Vec<usize>>;
type Limit = (usize, usize);
type Pos = (usize, usize);

#[derive(Copy, Clone, Eq, PartialEq)]
pub struct State {
    cost: usize,
    pos: Pos,
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        other.cost.cmp(&self.cost)
    }
}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}

fn new_dijkstra_cavern(cavern: &Cavern) -> Cavern {
    let mut dijkstra_cavern: Cavern = Vec::new();
    for row in cavern {
        dijkstra_cavern.push(vec![usize::MAX; row.len()]);
    }
    dijkstra_cavern[0][0] = 0;
    dijkstra_cavern
}

// from: https://doc.rust-lang.org/std/collections/binary_heap/index.html
fn fill_dijkstra_cavern(cavern: &Cavern, dijkstra: &mut Cavern) {
    let limits = (cavern.len(), cavern[0].len());
    let mut heap = BinaryHeap::new();
    heap.push(State {
        cost: 0,
        pos: (0, 0),
    });

    while let Some(State { cost, pos }) = heap.pop() {
        if cost > get_risk(dijkstra, pos) {
            continue;
        }

        for neighbor in get_neighbors(&pos, &limits) {
            let next = State {
                cost: cost + get_risk(cavern, neighbor),
                pos: neighbor,
            };

            if next.cost < get_risk(dijkstra, next.pos) {
                heap.push(next);
                dijkstra[next.pos.0][next.pos.1] = next.cost;
            }
        }
    }
}

fn calc_whole_cavern(cavern: &Cavern) -> Cavern {
    let mut whole_cavern = vec![vec![0; cavern[0].len() * 5]; cavern.len() * 5];
    for i in 0..whole_cavern.len() {
        for j in 0..whole_cavern[i].len() {
            let origin_i = i % cavern.len();
            let origin_j = j % cavern[0].len();
            whole_cavern[i][j] =
                match (cavern[origin_i][origin_j] + 1 * (i / cavern.len()) + (j / cavern[0].len()))
                    % 9
                {
                    0 => 9,
                    x => x,
                }
        }
    }
    whole_cavern
}

fn get_risk(cavern: &Cavern, pos: (usize, usize)) -> usize {
    cavern[pos.0][pos.1]
}

fn get_neighbors(pos: &(usize, usize), limits: &(usize, usize)) -> Vec<(usize, usize)> {
    let mut neighbors = Vec::new();

    if pos.0 > 0 {
        neighbors.push((pos.0 - 1, pos.1));
    }
    if pos.0 < limits.0 - 1 {
        neighbors.push((pos.0 + 1, pos.1));
    }
    if pos.1 > 0 {
        neighbors.push((pos.0, pos.1 - 1));
    }
    if pos.1 < limits.1 - 1 {
        neighbors.push((pos.0, pos.1 + 1));
    }

    neighbors
}

fn print_cavern(cavern: &Cavern) {
    for row in cavern {
        println!(
            "{}",
            row.iter()
                .map(|f| f.to_string())
                .collect::<Vec<String>>()
                .join("")
        )
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc15};

    #[test]
    fn aoc151_test() {
        assert_eq!(Aoc15::run(crate::Part::One, true), 40);
    }

    #[test]
    fn aoc151() {
        assert_eq!(Aoc15::run(crate::Part::One, false), 811);
    }

    #[test]
    fn aoc152_test() {
        assert_eq!(Aoc15::run(crate::Part::Two, true), 315);
    }

    #[test]
    fn aoc152() {
        assert_eq!(Aoc15::run(crate::Part::Two, false), 3012);
    }
}
