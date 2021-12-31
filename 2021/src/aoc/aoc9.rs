use std::collections::{HashMap, HashSet};

use crate::Aoc;

pub struct Aoc9 {}

impl Aoc<Floor, u32> for Aoc9 {
    fn input() -> String {
        String::from(include_str!("inputs/9.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/9_test.txt"))
    }

    fn parse_input(s: String) -> Floor {
        let mut hmap = HashMap::new();
        let rows_size = s.lines().count();
        let cols_size = s.lines().next().unwrap().chars().count();
        let size = (rows_size, cols_size);
        for (i, line) in s.lines().enumerate() {
            for (j, char) in line.chars().enumerate() {
                match char.to_digit(10) {
                    Some(v) => {
                        hmap.insert((i, j), v);
                    }
                    _ => {}
                }
            }
        }
        Floor::new(hmap, size)
    }

    fn part1(floor: Floor, _: bool) -> u32 {
        floor.risk()
    }

    fn part2(floor: Floor, _: bool) -> u32 {
        let mut basins: Vec<usize> = Vec::new();
        for point in floor.low_points() {
            basins.push(floor.basin(point).len())
        }

        basins.sort();
        basins
            .split_off(basins.len() - 3)
            .iter()
            .fold(1, |sum, &x| sum * x as u32)
    }
}

#[derive(Debug)]
pub struct Floor {
    hmap: HashMap<(usize, usize), u32>,
    size: (usize, usize),
}

impl Floor {
    pub fn new(hmap: HashMap<(usize, usize), u32>, size: (usize, usize)) -> Self {
        Self { hmap, size }
    }

    pub fn risk(&self) -> u32 {
        let mut risk = 0;
        for point in self.low_points() {
            risk += self.hmap[&point] + 1;
        }
        risk
    }

    pub fn low_points(&self) -> Vec<(usize, usize)> {
        let mut low_points: Vec<(usize, usize)> = Vec::new();

        for (pos, height) in &self.hmap {
            if self
                .neighboors(*pos)
                .iter()
                .all(|neighboor| self.hmap[neighboor] > *height)
            {
                low_points.push(*pos)
            }
        }
        low_points
    }

    pub fn basin(&self, pos: (usize, usize)) -> HashSet<(usize, usize)> {
        let mut basin = HashSet::new();
        basin.insert(pos);
        let pos_h = self.hmap[&pos];
        for point in self.neighboors(pos) {
            let point_h = self.hmap[&point];
            if point_h < 9 && point_h > pos_h {
                basin.extend(&self.basin(point));
            }
        }

        basin
    }

    pub fn neighboors(&self, pos: (usize, usize)) -> Vec<(usize, usize)> {
        let mut neighboors: Vec<(usize, usize)> = Vec::new();

        if pos.0 > 0 {
            neighboors.push((pos.0 - 1, pos.1))
        }
        if pos.0 < self.size.0 - 1 {
            neighboors.push((pos.0 + 1, pos.1))
        }
        if pos.1 > 0 {
            neighboors.push((pos.0, pos.1 - 1))
        }
        if pos.1 < self.size.1 - 1 {
            neighboors.push((pos.0, pos.1 + 1))
        }

        neighboors
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc9};

    #[test]
    fn aoc91_test() {
        assert_eq!(Aoc9::run(crate::Part::One, true), 15);
    }

    #[test]
    fn aoc91() {
        assert_eq!(Aoc9::run(crate::Part::One, false), 465);
    }

    #[test]
    fn aoc92_test() {
        assert_eq!(Aoc9::run(crate::Part::Two, true), 1134);
    }

    #[test]
    fn aoc92() {
        assert_eq!(Aoc9::run(crate::Part::Two, false), 1269555);
    }
}
