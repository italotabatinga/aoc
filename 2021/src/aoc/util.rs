#[derive(Debug)]
pub enum Part {
    One,
    Two,
}

pub trait Aoc<I, R>
where
    I: std::fmt::Debug,
    R: std::fmt::Debug,
{
    fn input() -> I;
    fn part1_test_input() -> I;
    fn part2_test_input() -> I;

    fn part1(input: I, is_test: bool) -> R;
    fn part2(input: I, is_test: bool) -> R;

    fn run(part: Part, is_test: bool) {
        let input = if is_test {
            match part {
                Part::One => Self::part1_test_input(),
                Part::Two => Self::part1_test_input(),
            }
        } else {
            Self::input()
        };
        let answer = match part {
            Part::One => Self::part1(input, is_test),
            Part::Two => Self::part2(input, is_test),
        };

        println!("Answer: {:?}", answer);
    }
}
