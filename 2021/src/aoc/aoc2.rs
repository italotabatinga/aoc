use super::util::Aoc;

struct Submarine {
    position: i32,
    depth: i32,
    aim: i32,
}

impl Submarine {
    pub fn new() -> Self {
        Self {
            position: 0,
            depth: 0,
            aim: 0,
        }
    }

    pub fn mult(&self) -> i32 {
        self.position * self.depth
    }

    fn parse_command(command: String) -> (String, i32) {
        let parsed_command: Vec<&str> = command.split(' ').collect();
        let (cmd, str_value) = (parsed_command[0], parsed_command[1]);
        let value = match str_value.parse::<i32>() {
            Ok(value) => value,
            _ => {
                println!("Unparsable value: \"{}\"", str_value);
                0
            }
        };

        (String::from(cmd), value)
    }

    pub fn process_command(&mut self, command: String) {
        let (cmd, value) = Self::parse_command(command);

        match cmd.as_str() {
            "forward" => self.position += value,
            "up" => self.depth -= value,
            "down" => self.depth += value,
            _ => println!("Unknown command: \"{}\"", cmd),
        }
    }

    pub fn process_command_with_aim(&mut self, command: String) {
        let (cmd, value) = Self::parse_command(command);

        match cmd.as_str() {
            "forward" => {
                self.position += value;
                self.depth += self.aim * value;
            }
            "up" => self.aim -= value,
            "down" => self.aim += value,
            _ => println!("Unknown command: \"{}\"", cmd),
        }
    }
}

pub struct Aoc2 {}

impl Aoc<Vec<String>, i32> for Aoc2 {
    fn input() -> String {
        String::from(include_str!("inputs/2.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/2_test.txt"))
    }

    fn part1(input: Vec<String>, _: bool) -> i32 {
        let commands = input;

        let mut submarine = Submarine::new();

        for command in commands {
            submarine.process_command(command);
        }

        submarine.mult()
    }

    fn part2(input: Vec<String>, _: bool) -> i32 {
        let commands = input;

        let mut submarine = Submarine::new();

        for command in commands {
            submarine.process_command_with_aim(command);
        }

        submarine.mult()
    }

    fn parse_input(s: String) -> Vec<String> {
        s.split("\n").map(|s| String::from(s.trim())).collect()
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc2};

    #[test]
    fn aoc21_test() {
        assert_eq!(Aoc2::run(crate::Part::One, true), 150);
    }

    #[test]
    fn aoc21() {
        assert_eq!(Aoc2::run(crate::Part::One, false), 1635930);
    }

    #[test]
    fn aoc22_test() {
        assert_eq!(Aoc2::run(crate::Part::Two, true), 900);
    }

    #[test]
    fn aoc22() {
        assert_eq!(Aoc2::run(crate::Part::Two, false), 1781819478);
    }
}