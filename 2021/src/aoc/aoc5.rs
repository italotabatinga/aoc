use std::collections::HashMap;

use crate::Aoc;

pub struct Aoc5 {}

impl Aoc<Vec<Line>, i32> for Aoc5 {
    fn input() -> String {
        String::from(include_str!("inputs/5.txt"))
    }

    fn parse_input(s: String) -> Vec<Line> {
        fn parse_str_point(s: &str) -> Point {
            // s like "1,2"
            let str_coords: Vec<&str> = s.split(",").collect();
            Point::new(
                str_coords[0].parse::<i32>().unwrap(),
                str_coords[1].parse::<i32>().unwrap(),
            )
        }

        s.split("\n")
            .map(|str_line| {
                let str_points: Vec<&str> = str_line.split(" -> ").collect();
                Line::new(
                    parse_str_point(str_points[0]),
                    parse_str_point(str_points[1]),
                )
            })
            .collect()
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/5_test.txt"))
    }

    fn part1(input: Vec<Line>, _: bool) -> i32 {
        let lines: Vec<Line> = input
            .into_iter()
            .filter(|line| line.is_horizontal() || line.is_vertical())
            .collect();
        let mut map: HashMap<(i32, i32), i32> = HashMap::new();

        for line in lines {
            for point in line.path() {
                let key = (point.x, point.y);
                match map.get_mut(&key) {
                    Some(v) => *v += 1,
                    None => {
                        map.insert(key, 1);
                    }
                }
            }
        }

        map.values()
            .fold(0, |count, &v| if v >= 2 { count + 1 } else { count })
    }

    fn part2(input: Vec<Line>, _: bool) -> i32 {
        let lines: Vec<Line> = input;
        let mut map: HashMap<(i32, i32), i32> = HashMap::new();

        for line in lines {
            for point in line.path() {
                let key = (point.x, point.y);
                match map.get_mut(&key) {
                    Some(v) => *v += 1,
                    None => {
                        map.insert(key, 1);
                    }
                }
            }
        }

        map.values()
            .fold(0, |count, &v| if v >= 2 { count + 1 } else { count })
    }
}

#[derive(Clone, Copy, Debug)]
pub struct Point {
    x: i32,
    y: i32,
}

impl Point {
    pub fn new(x: i32, y: i32) -> Self {
        Self { x, y }
    }

    pub fn dir(&self, rhs: &Point) -> Self {
        let dx = rhs.x - self.x;
        let dy = rhs.y - self.y;
        Point {
            x: dx.checked_div(dx.abs()).unwrap_or(dx),
            y: dy.checked_div(dy.abs()).unwrap_or(dy),
        }
    }

    pub fn add(&self, rhs: &Point) -> Self {
        Point {
            x: self.x + rhs.x,
            y: self.y + rhs.y,
        }
    }

    pub fn eq(&self, rhs: &Point) -> bool {
        self.x == rhs.x && self.y == rhs.y
    }
}

#[derive(Debug)]
pub struct Line {
    a: Point,
    b: Point,
    dir: Point,
}

impl Line {
    pub fn new(a: Point, b: Point) -> Self {
        let dir = a.dir(&b);
        Self { a, b, dir }
    }

    pub fn is_vertical(&self) -> bool {
        self.a.x == self.b.x
    }

    pub fn is_horizontal(&self) -> bool {
        self.a.y == self.b.y
    }

    pub fn path(&self) -> Vec<Point> {
        let mut point = self.a;
        let mut path: Vec<Point> = Vec::new();
        while !point.eq(&self.b) {
            path.push(point.clone());
            point = point.add(&self.dir);
        }
        path.push(self.b);
        path
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc5};

    #[test]
    fn aoc51_test() {
        assert_eq!(Aoc5::run(crate::Part::One, true), 5);
    }

    #[test]
    fn aoc51() {
        assert_eq!(Aoc5::run(crate::Part::One, false), 5608);
    }

    #[test]
    fn aoc52_test() {
        assert_eq!(Aoc5::run(crate::Part::Two, true), 12);
    }

    #[test]
    fn aoc52() {
        assert_eq!(Aoc5::run(crate::Part::Two, false), 20299);
    }
}