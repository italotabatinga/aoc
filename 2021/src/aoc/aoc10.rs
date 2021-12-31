use crate::Aoc;

pub struct Aoc10 {}

impl Aoc<Vec<Line>, u64> for Aoc10 {
    fn input() -> String {
        String::from(include_str!("inputs/10.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/10_test.txt"))
    }

    fn parse_input(s: String) -> Vec<Line> {
        s.lines()
            .map(|line| Line::new(String::from(line)))
            .collect()
    }

    fn part1(input: Vec<Line>, _: bool) -> u64 {
        input.iter().fold(0, |acc, line| match line.illegal_char {
            Some(_) => line.syntax_error_score() + acc,
            None => acc,
        })
    }

    fn part2(input: Vec<Line>, _: bool) -> u64 {
        let mut scores: Vec<u64> = Vec::new();

        for line in input {
            if !line.completion_string.is_empty() {
                scores.push(line.completion_score());
            }
        }
        scores.sort();

        scores[scores.len() / 2]
    }
}

#[derive(Debug)]
pub struct Line {
    raw: String,
    pub illegal_char: Option<char>,
    pub completion_string: String,
}

impl Line {
    pub fn new(raw: String) -> Self {
        let processed = Self::process(raw.clone());
        Self {
            raw,
            illegal_char: match processed {
                Err(c) => Some(c),
                _ => None,
            },
            completion_string: match processed {
                Ok(s) => s,
                _ => String::new(),
            },
        }
    }

    fn process(s: String) -> Result<String, char> {
        let mut stack: Vec<char> = Vec::new();

        for c in s.chars() {
            match stack.last() {
                Some(last) => {
                    let op = Op::from(c);
                    let last_op = Op::from(*last);
                    if op.is_opening() {
                        stack.push(c);
                    } else if op.closes(last_op) {
                        stack.pop();
                    } else {
                        return Err(c);
                    }
                }
                None => stack.push(c),
            }
        }

        let mut completion_string = String::new();
        for c in stack.iter().rev() {
            let op = Op::from(*c);

            completion_string.push(op.get_closing());
        }

        Ok(completion_string)
    }

    pub fn syntax_error_score(&self) -> u64 {
        match self.illegal_char {
            Some(v) => match v {
                ')' => 3,
                ']' => 57,
                '}' => 1197,
                '>' => 25137,
                _ => 0,
            },
            None => 0,
        }
    }

    pub fn completion_score(&self) -> u64 {
        self.completion_string.chars().fold(0, |acc, c| {
            acc * 5
                + match c {
                    ')' => 1,
                    ']' => 2,
                    '}' => 3,
                    '>' => 4,
                    _ => 0,
                }
        })
    }
}

#[derive(PartialEq)]
pub enum Op {
    OpPar,
    ClPar,
    OpBr,
    ClBr,
    OpCBr,
    ClCBr,
    OpABr,
    ClABr,
}

impl Op {
    pub fn from(c: char) -> Self {
        match c {
            '(' => Self::OpPar,
            ')' => Self::ClPar,
            '[' => Self::OpBr,
            ']' => Self::ClBr,
            '{' => Self::OpCBr,
            '}' => Self::ClCBr,
            '<' => Self::OpABr,
            '>' => Self::ClABr,
            _ => {
                panic!("Unexpected char \'{}\'", c);
            }
        }
    }

    pub fn is_opening(&self) -> bool {
        match *self {
            Self::OpPar => true,
            Self::OpBr | Self::OpCBr | Self::OpABr => true,
            _ => false,
        }
    }

    pub fn closes(&self, op: Self) -> bool {
        match *self {
            Self::ClPar => op == Self::OpPar,
            Self::ClBr => op == Self::OpBr,
            Self::ClCBr => op == Self::OpCBr,
            Self::ClABr => op == Self::OpABr,
            _ => false,
        }
    }

    pub fn get_closing(&self) -> char {
        match *self {
            Self::OpPar => ')',
            Self::OpBr => ']',
            Self::OpCBr => '}',
            Self::OpABr => '>',
            _ => ' ',
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc10};

    #[test]
    fn aoc101_test() {
        assert_eq!(Aoc10::run(crate::Part::One, true), 26397);
    }

    #[test]
    fn aoc101() {
        assert_eq!(Aoc10::run(crate::Part::One, false), 362271);
    }

    #[test]
    fn aoc102_test() {
        assert_eq!(Aoc10::run(crate::Part::Two, true), 288957);
    }

    #[test]
    fn aoc102() {
        assert_eq!(Aoc10::run(crate::Part::Two, false), 1698395182);
    }
}
