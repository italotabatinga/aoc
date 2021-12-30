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
    fn input() -> Vec<String> {
        Vec::from([
            "forward 2",
            "forward 2",
            "down 7",
            "forward 6",
            "down 7",
            "forward 4",
            "down 7",
            "up 2",
            "forward 4",
            "down 2",
            "down 4",
            "down 5",
            "forward 4",
            "up 1",
            "up 2",
            "forward 9",
            "forward 7",
            "down 1",
            "forward 4",
            "forward 7",
            "forward 9",
            "down 2",
            "down 3",
            "forward 1",
            "forward 4",
            "forward 6",
            "up 2",
            "forward 6",
            "down 2",
            "down 5",
            "up 4",
            "down 5",
            "forward 6",
            "down 7",
            "down 4",
            "up 5",
            "forward 9",
            "forward 5",
            "down 8",
            "forward 3",
            "forward 1",
            "down 7",
            "down 5",
            "forward 2",
            "down 3",
            "forward 8",
            "down 1",
            "forward 4",
            "up 8",
            "up 8",
            "up 8",
            "forward 2",
            "down 3",
            "down 5",
            "up 8",
            "forward 1",
            "forward 3",
            "up 1",
            "up 2",
            "up 4",
            "down 8",
            "forward 5",
            "down 9",
            "down 9",
            "forward 9",
            "up 3",
            "down 2",
            "down 9",
            "up 8",
            "forward 3",
            "up 2",
            "down 7",
            "forward 9",
            "forward 1",
            "up 1",
            "forward 9",
            "forward 1",
            "down 1",
            "forward 7",
            "up 3",
            "up 5",
            "forward 3",
            "forward 8",
            "forward 4",
            "down 3",
            "forward 5",
            "forward 4",
            "forward 2",
            "down 2",
            "down 3",
            "forward 1",
            "up 2",
            "up 2",
            "down 9",
            "down 6",
            "down 6",
            "up 2",
            "down 6",
            "forward 6",
            "forward 5",
            "down 8",
            "up 3",
            "up 2",
            "up 1",
            "forward 8",
            "down 3",
            "down 5",
            "forward 8",
            "forward 7",
            "down 9",
            "down 7",
            "down 9",
            "down 2",
            "forward 8",
            "down 9",
            "forward 2",
            "forward 2",
            "down 5",
            "up 8",
            "forward 9",
            "down 7",
            "down 9",
            "forward 2",
            "forward 9",
            "forward 1",
            "forward 3",
            "forward 1",
            "forward 3",
            "down 4",
            "forward 6",
            "forward 4",
            "down 7",
            "up 4",
            "up 5",
            "forward 1",
            "down 4",
            "down 8",
            "down 3",
            "down 8",
            "up 3",
            "forward 2",
            "up 7",
            "down 1",
            "down 7",
            "down 3",
            "forward 3",
            "forward 4",
            "up 8",
            "down 2",
            "forward 8",
            "down 8",
            "up 8",
            "up 5",
            "forward 5",
            "forward 1",
            "forward 5",
            "up 4",
            "down 5",
            "down 1",
            "forward 3",
            "forward 4",
            "down 8",
            "down 9",
            "forward 4",
            "down 5",
            "down 5",
            "forward 9",
            "forward 7",
            "down 3",
            "down 4",
            "forward 2",
            "down 7",
            "down 3",
            "down 9",
            "down 8",
            "forward 6",
            "forward 3",
            "forward 1",
            "forward 5",
            "down 4",
            "down 4",
            "forward 9",
            "down 9",
            "down 5",
            "up 8",
            "forward 1",
            "forward 3",
            "down 1",
            "down 4",
            "forward 1",
            "up 5",
            "down 9",
            "forward 8",
            "down 8",
            "down 2",
            "down 3",
            "forward 2",
            "forward 9",
            "forward 7",
            "up 3",
            "down 3",
            "forward 4",
            "down 8",
            "forward 7",
            "down 3",
            "forward 3",
            "down 5",
            "forward 8",
            "forward 7",
            "forward 4",
            "down 2",
            "up 5",
            "forward 9",
            "forward 8",
            "down 7",
            "up 5",
            "up 8",
            "down 4",
            "down 9",
            "forward 5",
            "up 6",
            "forward 7",
            "down 4",
            "forward 8",
            "up 4",
            "up 2",
            "down 8",
            "forward 4",
            "down 8",
            "down 3",
            "down 3",
            "forward 1",
            "up 7",
            "forward 5",
            "forward 1",
            "down 5",
            "forward 4",
            "down 3",
            "down 4",
            "forward 4",
            "down 5",
            "up 9",
            "forward 6",
            "down 6",
            "up 6",
            "down 8",
            "forward 2",
            "down 9",
            "down 8",
            "forward 7",
            "down 2",
            "forward 1",
            "forward 4",
            "down 7",
            "forward 1",
            "up 8",
            "forward 6",
            "forward 2",
            "forward 5",
            "up 1",
            "forward 6",
            "up 7",
            "down 8",
            "down 5",
            "down 2",
            "forward 4",
            "forward 3",
            "down 7",
            "up 5",
            "forward 1",
            "forward 2",
            "forward 3",
            "forward 6",
            "down 9",
            "down 2",
            "forward 7",
            "up 3",
            "down 4",
            "forward 3",
            "forward 3",
            "forward 6",
            "up 2",
            "down 4",
            "forward 8",
            "down 3",
            "down 4",
            "forward 1",
            "forward 3",
            "forward 8",
            "forward 1",
            "forward 6",
            "forward 1",
            "forward 7",
            "down 8",
            "up 3",
            "up 5",
            "forward 4",
            "forward 6",
            "down 3",
            "forward 5",
            "forward 5",
            "down 9",
            "forward 4",
            "down 5",
            "forward 6",
            "down 1",
            "forward 2",
            "forward 3",
            "forward 2",
            "up 4",
            "forward 9",
            "forward 7",
            "forward 6",
            "up 1",
            "down 6",
            "down 1",
            "forward 5",
            "up 4",
            "up 9",
            "forward 5",
            "forward 3",
            "up 9",
            "down 4",
            "forward 7",
            "up 9",
            "down 7",
            "forward 5",
            "down 4",
            "down 9",
            "up 4",
            "forward 2",
            "forward 2",
            "forward 1",
            "forward 3",
            "forward 4",
            "down 3",
            "down 5",
            "up 9",
            "forward 8",
            "up 5",
            "down 3",
            "forward 5",
            "down 4",
            "forward 1",
            "forward 1",
            "up 5",
            "down 7",
            "up 7",
            "down 3",
            "forward 7",
            "down 3",
            "down 3",
            "down 5",
            "down 2",
            "down 7",
            "down 5",
            "down 1",
            "up 3",
            "forward 5",
            "forward 9",
            "up 1",
            "down 8",
            "up 5",
            "forward 4",
            "down 5",
            "forward 5",
            "down 3",
            "up 8",
            "forward 2",
            "up 7",
            "down 4",
            "forward 1",
            "up 6",
            "down 7",
            "down 4",
            "forward 9",
            "forward 3",
            "forward 7",
            "up 5",
            "down 1",
            "down 3",
            "up 2",
            "forward 3",
            "down 8",
            "up 7",
            "forward 9",
            "forward 9",
            "forward 9",
            "down 3",
            "forward 4",
            "forward 9",
            "down 6",
            "up 1",
            "up 7",
            "down 3",
            "forward 4",
            "forward 9",
            "down 9",
            "up 4",
            "down 2",
            "forward 6",
            "forward 4",
            "down 5",
            "up 3",
            "down 6",
            "up 4",
            "down 5",
            "down 7",
            "down 3",
            "up 1",
            "down 2",
            "forward 1",
            "forward 8",
            "up 7",
            "forward 4",
            "down 5",
            "down 3",
            "up 7",
            "forward 1",
            "forward 7",
            "up 6",
            "forward 4",
            "forward 5",
            "forward 3",
            "forward 1",
            "down 5",
            "up 3",
            "up 3",
            "down 8",
            "forward 1",
            "forward 2",
            "up 6",
            "down 3",
            "down 6",
            "down 5",
            "down 4",
            "up 1",
            "down 9",
            "forward 3",
            "down 8",
            "up 9",
            "down 3",
            "forward 6",
            "down 8",
            "forward 2",
            "forward 4",
            "forward 6",
            "forward 3",
            "forward 6",
            "up 6",
            "up 3",
            "down 6",
            "up 1",
            "forward 9",
            "forward 2",
            "up 6",
            "forward 7",
            "down 5",
            "forward 5",
            "forward 7",
            "down 6",
            "down 9",
            "forward 9",
            "forward 1",
            "up 4",
            "forward 3",
            "forward 2",
            "forward 2",
            "up 4",
            "down 7",
            "down 2",
            "up 2",
            "down 2",
            "up 9",
            "forward 4",
            "forward 3",
            "forward 3",
            "up 4",
            "down 7",
            "forward 5",
            "down 5",
            "forward 1",
            "down 4",
            "forward 4",
            "up 1",
            "down 5",
            "forward 2",
            "up 3",
            "up 1",
            "down 9",
            "up 9",
            "down 3",
            "forward 1",
            "down 6",
            "forward 8",
            "down 5",
            "forward 7",
            "down 4",
            "forward 2",
            "down 3",
            "down 1",
            "forward 6",
            "forward 4",
            "down 9",
            "forward 3",
            "up 4",
            "up 8",
            "down 2",
            "up 5",
            "up 8",
            "up 7",
            "down 2",
            "down 9",
            "forward 8",
            "down 4",
            "up 7",
            "forward 2",
            "up 4",
            "forward 7",
            "down 3",
            "down 3",
            "down 6",
            "up 3",
            "down 6",
            "down 6",
            "up 7",
            "down 6",
            "up 9",
            "down 6",
            "forward 3",
            "forward 8",
            "forward 3",
            "down 9",
            "forward 6",
            "forward 8",
            "down 7",
            "forward 3",
            "forward 7",
            "forward 3",
            "forward 8",
            "forward 6",
            "down 9",
            "up 4",
            "forward 1",
            "forward 6",
            "forward 2",
            "forward 4",
            "down 4",
            "down 9",
            "forward 7",
            "forward 4",
            "forward 4",
            "up 3",
            "up 6",
            "forward 4",
            "forward 7",
            "forward 4",
            "down 2",
            "forward 8",
            "up 4",
            "forward 8",
            "up 2",
            "up 4",
            "down 5",
            "forward 5",
            "up 2",
            "up 8",
            "up 9",
            "forward 6",
            "forward 2",
            "up 9",
            "forward 5",
            "up 3",
            "forward 8",
            "forward 8",
            "down 5",
            "forward 6",
            "up 6",
            "down 6",
            "down 1",
            "forward 2",
            "down 1",
            "down 5",
            "forward 2",
            "down 3",
            "up 7",
            "forward 6",
            "up 9",
            "down 3",
            "up 3",
            "forward 3",
            "down 9",
            "down 4",
            "down 2",
            "forward 7",
            "forward 2",
            "down 4",
            "forward 5",
            "up 1",
            "down 6",
            "down 6",
            "up 2",
            "down 2",
            "down 4",
            "down 4",
            "forward 8",
            "down 3",
            "down 5",
            "down 3",
            "up 2",
            "down 6",
            "down 9",
            "up 8",
            "forward 3",
            "down 9",
            "down 6",
            "down 7",
            "down 6",
            "forward 9",
            "up 1",
            "forward 9",
            "down 8",
            "forward 8",
            "up 5",
            "down 8",
            "forward 5",
            "down 6",
            "down 9",
            "forward 1",
            "up 5",
            "up 5",
            "forward 4",
            "forward 1",
            "forward 6",
            "down 1",
            "down 4",
            "down 3",
            "down 1",
            "down 1",
            "down 9",
            "down 3",
            "forward 6",
            "up 9",
            "up 4",
            "down 1",
            "up 6",
            "forward 3",
            "down 2",
            "down 8",
            "forward 1",
            "forward 8",
            "forward 4",
            "down 4",
            "down 5",
            "up 5",
            "up 5",
            "forward 7",
            "forward 6",
            "forward 4",
            "down 5",
            "forward 7",
            "down 4",
            "forward 2",
            "forward 5",
            "forward 3",
            "forward 5",
            "up 6",
            "up 9",
            "down 6",
            "forward 8",
            "down 3",
            "up 2",
            "forward 8",
            "down 9",
            "forward 9",
            "down 3",
            "down 4",
            "down 5",
            "up 3",
            "forward 6",
            "forward 6",
            "forward 8",
            "up 6",
            "up 4",
            "forward 2",
            "down 1",
            "down 5",
            "forward 1",
            "forward 8",
            "up 5",
            "down 2",
            "forward 4",
            "down 9",
            "up 5",
            "down 8",
            "forward 6",
            "forward 3",
            "down 8",
            "up 1",
            "down 5",
            "down 6",
            "forward 2",
            "forward 7",
            "down 8",
            "forward 2",
            "forward 8",
            "forward 1",
            "down 1",
            "down 3",
            "down 8",
            "down 1",
            "forward 1",
            "forward 8",
            "up 4",
            "up 5",
            "forward 2",
            "forward 5",
            "forward 5",
            "forward 9",
            "forward 2",
            "forward 4",
            "forward 4",
            "down 1",
            "up 8",
            "forward 3",
            "down 9",
            "forward 9",
            "up 5",
            "forward 7",
            "forward 3",
            "up 5",
            "up 5",
            "down 9",
            "forward 9",
            "up 2",
            "forward 1",
            "forward 1",
            "down 5",
            "down 7",
            "forward 2",
            "up 8",
            "up 7",
            "down 4",
            "down 6",
            "down 6",
            "down 6",
            "forward 4",
            "forward 7",
            "down 6",
            "forward 3",
            "forward 5",
            "down 2",
            "up 3",
            "forward 3",
            "forward 6",
            "forward 9",
            "forward 7",
            "forward 6",
            "down 4",
            "down 9",
            "forward 4",
            "forward 4",
            "down 8",
            "down 7",
            "down 1",
            "down 6",
            "up 7",
            "forward 3",
            "forward 9",
            "forward 9",
            "up 7",
            "forward 2",
            "up 5",
            "up 6",
            "down 7",
            "down 7",
            "down 5",
            "forward 8",
            "down 1",
            "up 4",
            "down 4",
            "forward 4",
            "forward 3",
            "forward 7",
            "forward 4",
            "down 6",
            "down 3",
            "down 3",
            "forward 6",
            "forward 7",
            "forward 6",
            "up 5",
            "down 9",
            "forward 9",
            "forward 5",
            "up 5",
            "down 9",
            "up 8",
            "down 6",
            "down 8",
            "down 8",
            "down 8",
            "forward 2",
            "down 8",
            "down 4",
            "down 6",
            "forward 8",
            "forward 3",
            "forward 2",
            "down 5",
            "down 3",
            "forward 1",
            "forward 2",
            "up 9",
            "forward 6",
            "down 8",
            "up 2",
            "down 2",
            "forward 7",
            "up 6",
            "up 2",
            "forward 1",
            "down 2",
            "forward 7",
            "forward 4",
            "down 6",
            "forward 4",
            "up 4",
            "forward 3",
            "down 2",
            "down 5",
            "up 5",
            "down 2",
            "down 1",
            "down 9",
            "forward 3",
            "down 5",
            "up 8",
            "down 6",
            "forward 3",
            "forward 9",
            "up 6",
            "down 2",
            "forward 6",
            "forward 7",
            "down 1",
            "up 9",
            "up 9",
            "up 9",
            "up 4",
            "forward 2",
            "down 7",
            "down 1",
            "forward 8",
            "down 2",
            "down 1",
            "forward 9",
            "forward 8",
            "up 8",
            "down 5",
            "forward 8",
            "down 1",
            "down 4",
            "down 3",
            "down 1",
            "forward 2",
            "down 2",
            "down 6",
            "forward 5",
            "down 9",
            "forward 8",
            "down 9",
            "up 6",
            "up 5",
            "forward 8",
            "forward 3",
            "down 9",
            "up 6",
            "down 8",
            "down 3",
            "down 2",
            "up 1",
            "up 9",
            "up 7",
            "up 1",
            "up 4",
            "forward 7",
            "down 8",
            "up 3",
            "forward 1",
            "up 3",
            "up 7",
            "forward 4",
            "up 3",
            "down 1",
            "forward 3",
            "forward 4",
            "forward 7",
            "down 4",
            "down 4",
            "down 2",
            "up 1",
            "up 5",
            "up 3",
            "forward 6",
            "forward 6",
            "forward 9",
            "down 9",
            "forward 2",
            "up 3",
            "forward 3",
            "down 2",
            "down 4",
            "forward 6",
            "forward 7",
            "down 5",
            "forward 9",
            "up 7",
            "down 9",
            "down 1",
            "down 1",
            "forward 1",
            "up 1",
            "up 5",
            "down 5",
            "forward 4",
            "up 3",
            "down 1",
            "down 3",
            "forward 3",
            "forward 6",
            "up 4",
            "down 2",
            "forward 1",
            "forward 5",
            "forward 3",
            "forward 2",
            "forward 4",
            "down 3",
            "forward 2",
            "down 4",
            "down 7",
            "up 4",
            "down 7",
            "forward 8",
            "forward 2",
            "down 4",
            "down 2",
            "forward 8",
            "up 7",
            "forward 1",
            "up 4",
            "forward 4",
            "up 7",
            "forward 2",
            "forward 7",
            "forward 9",
            "down 8",
            "down 2",
            "forward 8",
            "down 3",
            "up 5",
            "forward 3",
            "up 9",
            "forward 8",
            "forward 5",
            "up 2",
            "down 5",
            "down 9",
            "up 9",
            "forward 1",
            "forward 3",
            "down 3",
            "down 6",
            "forward 9",
            "forward 5",
            "down 4",
            "up 7",
            "forward 4",
            "forward 3",
            "down 2",
            "up 1",
            "forward 3",
            "down 4",
            "down 8",
            "up 4",
            "down 9",
            "forward 6",
            "down 7",
            "down 3",
            "down 9",
            "forward 3",
            "down 8",
            "forward 7",
            "forward 7",
            "down 1",
            "down 1",
            "forward 1",
            "down 4",
            "forward 8",
        ])
        .iter()
        .map(|&s| String::from(s))
        .collect()
    }

    fn part1_test_input() -> Vec<String> {
        Vec::from([
            "forward 5",
            "down 5",
            "forward 8",
            "up 3",
            "down 8",
            "forward 2",
        ])
        .iter()
        .map(|&s| String::from(s))
        .collect()
    }

    fn part2_test_input() -> Vec<String> {
        Self::part1_test_input()
    }

    fn part1(input: Vec<String>) -> i32 {
        let commands = input;

        let mut submarine = Submarine::new();

        for command in commands {
            submarine.process_command(command);
        }

        submarine.mult()
    }

    fn part2(input: Vec<String>) -> i32 {
        let commands = input;

        let mut submarine = Submarine::new();

        for command in commands {
            submarine.process_command_with_aim(command);
        }

        submarine.mult()
    }
}
