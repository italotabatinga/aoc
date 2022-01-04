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
    fn input() -> String;
    fn part1_test_input() -> String;
    fn part2_test_input() -> String {
        Self::part1_test_input()
    }
    fn parse_input(s: String) -> I;

    fn part1(input: I, is_test: bool) -> R;
    fn part2(input: I, is_test: bool) -> R;

    fn run(part: Part, is_test: bool) -> R {
        let input = Self::parse_input(if is_test {
            match part {
                Part::One => Self::part1_test_input(),
                Part::Two => Self::part2_test_input(),
            }
        } else {
            Self::input()
        });
        let answer = match part {
            Part::One => Self::part1(input, is_test),
            Part::Two => Self::part2(input, is_test),
        };

        if !cfg!(test) {
            println!("Answer: {:?}", answer)
        }
        answer
    }
}
