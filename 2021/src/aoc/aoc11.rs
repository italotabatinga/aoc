use crate::Aoc;

pub struct Aoc11 {}

impl Aoc<Vec<Vec<u32>>, u32> for Aoc11 {
    fn input() -> String {
        String::from(include_str!("inputs/11.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/11_test.txt"))
    }

    fn parse_input(s: String) -> Vec<Vec<u32>> {
        s.lines()
            .map(|line| line.chars().map(|c| c.to_digit(10).unwrap()).collect())
            .collect()
    }

    fn part1(mut input: Vec<Vec<u32>>, _: bool) -> u32 {
        const STEPS: u32 = 100;
        let mut count: u32 = 0;

        for _ in 0..STEPS {
            count += run_step(&mut input);
        }

        count
    }

    fn part2(mut input: Vec<Vec<u32>>, _: bool) -> u32 {
        const MAX_STEPS: u32 = 10000;

        for i in 0..MAX_STEPS {
            let count = run_step(&mut input);

            if count == 100 {
                return i + 1;
            }
        }

        MAX_STEPS
    }
}

const SIZE: (usize, usize) = (10, 10);

fn run_step(map: &mut Vec<Vec<u32>>) -> u32 {
    pump_energy(map);

    for i in 0..SIZE.0 {
        for j in 0..SIZE.1 {
            try_flash((i, j), map);
        }
    }

    count_flashes(map)
}

fn pump_energy(map: &mut Vec<Vec<u32>>) {
    for i in 0..SIZE.0 {
        for j in 0..SIZE.1 {
            map[i][j] += 1;
        }
    }
}

fn try_flash((i, j): (usize, usize), map: &mut Vec<Vec<u32>>) {
    if map[i][j] <= 9 {
        return;
    }

    map[i][j] = 0;

    for (m, n) in get_neigbors((i, j)) {
        if map[m][n] == 0 {
            continue;
        }
        map[m][n] += 1;

        try_flash((m, n), map);
    }
}

fn count_flashes(map: &Vec<Vec<u32>>) -> u32 {
    let mut count = 0;
    for i in 0..SIZE.0 {
        for j in 0..SIZE.1 {
            if map[i][j] == 0 {
                count += 1;
            }
        }
    }
    count
}

fn get_neigbors(pos: (usize, usize)) -> Vec<(usize, usize)> {
    let mut neighbors: Vec<(usize, usize)> = Vec::new();

    if pos.0 > 0 {
        neighbors.push((pos.0 - 1, pos.1));
        if pos.1 > 0 {
            neighbors.push((pos.0 - 1, pos.1 - 1))
        }
        if pos.1 < SIZE.1 - 1 {
            neighbors.push((pos.0 - 1, pos.1 + 1))
        }
    }
    if pos.0 < SIZE.0 - 1 {
        neighbors.push((pos.0 + 1, pos.1));
        if pos.1 > 0 {
            neighbors.push((pos.0 + 1, pos.1 - 1))
        }
        if pos.1 < SIZE.1 - 1 {
            neighbors.push((pos.0 + 1, pos.1 + 1))
        }
    }
    if pos.1 > 0 {
        neighbors.push((pos.0, pos.1 - 1));
    }
    if pos.1 < SIZE.1 - 1 {
        neighbors.push((pos.0, pos.1 + 1));
    }

    neighbors
}

#[allow(dead_code)]
fn print_map(map: Vec<Vec<u32>>) {
    for row in map {
        println!(
            "{}",
            row.iter()
                .fold(String::new(), |acc, x| acc + &x.to_string())
        )
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc11};

    #[test]
    fn aoc111_test() {
        assert_eq!(Aoc11::run(crate::Part::One, true), 1656);
    }

    #[test]
    fn aoc111() {
        assert_eq!(Aoc11::run(crate::Part::One, false), 1594);
    }

    #[test]
    fn aoc112_test() {
        assert_eq!(Aoc11::run(crate::Part::Two, true), 195);
    }

    #[test]
    fn aoc112() {
        assert_eq!(Aoc11::run(crate::Part::Two, false), 437);
    }
}
