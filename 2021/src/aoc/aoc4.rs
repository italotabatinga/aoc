use std::collections::HashSet;

use crate::Aoc;

pub struct Aoc4 {}

impl Aoc<(Vec<u8>, Vec<Board>), u32> for Aoc4 {
    fn input() -> String {
        String::from(include_str!("inputs/4.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/4_test.txt"))
    }

    fn part1(input: (Vec<u8>, Vec<Board>), _: bool) -> u32 {
        let (chosen_numbers, mut boards) = input;

        for chosen in chosen_numbers {
            for board in &mut boards {
                board.try_check(chosen);
            }

            for board in &mut boards {
                if board.has_won {
                    return board.sum_of_unchecked() * chosen as u32;
                }
            }
        }
        0
    }

    fn part2(input: (Vec<u8>, Vec<Board>), _: bool) -> u32 {
        let (chosen_numbers, mut boards) = input;
        let mut winning_boards: HashSet<usize> = HashSet::new();
        for chosen in chosen_numbers {
            for board in &mut boards {
                board.try_check(chosen);
            }

            for (i, board) in &mut boards.iter().enumerate() {
                if board.has_won {
                    if !winning_boards.contains(&i) {
                        if winning_boards.len() == boards.len() - 1 {
                          return board.sum_of_unchecked() * chosen as u32;
                        }
                        winning_boards.insert(i);
                    }
                }
            }
        }
        0
    }

    fn parse_input(s: String) -> (Vec<u8>, Vec<Board>) {
        let parsed: Vec<String> = s.split("\n\n").map(|s| String::from(s)).collect();
        let chosen_numbers: Vec<u8> = parse_chosen_numbers(&parsed[0]);
        let boards: Vec<Board> = (1..parsed.len())
            .into_iter()
            .map(|i| parse_board(&parsed[i]))
            .collect();
        (chosen_numbers, boards)
    }
}

#[derive(Debug)]
pub struct Board {
    mat: Vec<Vec<u8>>,
    checks: [[bool; 5]; 5],
    pub has_won: bool,
}

impl Board {
    pub fn new(mat: Vec<Vec<u8>>) -> Self {
        Self {
            mat,
            checks: [[false; 5]; 5],
            has_won: false,
        }
    }

    pub fn try_check(&mut self, num: u8) {
        for (i, arr) in self.mat.iter().enumerate() {
            for (j, &value) in arr.iter().enumerate() {
                if value == num {
                    self.checks[i][j] = true;

                    if self.has_completed(i, j) {
                        self.has_won = true;
                    }
                }
            }
        }
    }

    fn has_completed(&self, row: usize, col: usize) -> bool {
        // check row
        let mut has_completed = true;
        for check in self.checks[row] {
            has_completed = has_completed && check;
        }
        if has_completed {
            return has_completed;
        }

        // check col
        has_completed = true;
        for i in 0..5 {
            let check = self.checks[i][col];
            has_completed = has_completed && check;
        }
        has_completed
    }

    pub fn sum_of_unchecked(&self) -> u32 {
        let mut sum: u32 = 0;
        for (i, row) in self.mat.iter().enumerate() {
            for (j, num) in row.iter().enumerate() {
                if !self.checks[i][j] {
                    sum += *num as u32;
                }
            }
        }
        sum
    }
}

fn parse_chosen_numbers(s: &String) -> Vec<u8> {
    s.split(",")
        .map(|s| match s.parse::<u8>() {
            Ok(v) => v,
            _ => {
                println!("Failed to parse \"{}\"", s);
                0
            }
        })
        .collect()
}

fn parse_board(s: &String) -> Board {
    let vec_board: Vec<Vec<u8>> = s
        .split("\n")
        .map(|s| {
            s.trim()
                .split_whitespace()
                .map(|s| match s.parse::<u8>() {
                    Ok(v) => v,
                    _ => {
                        println!("Failed to parse \"{}\"", s);
                        0
                    }
                })
                .collect()
        })
        .collect();
    Board::new(vec_board)
}


#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc4};

    #[test]
    fn aoc41_test() {
        assert_eq!(Aoc4::run(crate::Part::One, true), 4512);
    }

    #[test]
    fn aoc41() {
        assert_eq!(Aoc4::run(crate::Part::One, false), 51034);
    }

    #[test]
    fn aoc42_test() {
        assert_eq!(Aoc4::run(crate::Part::Two, true), 1924);
    }

    #[test]
    fn aoc42() {
        assert_eq!(Aoc4::run(crate::Part::Two, false), 5434);
    }
}