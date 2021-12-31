use std::env;

mod aoc;

pub use aoc::*;

fn main() {
    let args: Vec<String> = env::args().collect();
    let module = match args.get(1) {
        Some(module) => module,
        None => {
            println!("Excepted a module on the first argument.");
            return;
        }
    };

    let is_test = args.contains(&String::from("test"));
    match module.as_str() {
        "1.1" => {
            Aoc1::run(Part::One, is_test);
        }
        "1.2" => {
            Aoc1::run(Part::Two, is_test);
        }
        "2.1" => {
            Aoc2::run(Part::One, is_test);
        }
        "2.2" => {
            Aoc2::run(Part::Two, is_test);
        }
        "3.1" => {
            Aoc3::run(Part::One, is_test);
        }
        "3.2" => {
            Aoc3::run(Part::Two, is_test);
        }
        "4.1" => {
            Aoc4::run(Part::One, is_test);
        }
        "4.2" => {
            Aoc4::run(Part::Two, is_test);
        }
        "5.1" => {
            Aoc5::run(Part::One, is_test);
        }
        "5.2" => {
            Aoc5::run(Part::Two, is_test);
        }
        "6.1" => {
            Aoc6::run(Part::One, is_test);
        }
        "6.2" => {
            Aoc6::run(Part::Two, is_test);
        }
        "7.1" => {
            Aoc7::run(Part::One, is_test);
        }
        "7.2" => {
            Aoc7::run(Part::Two, is_test);
        }
        "8.1" => {
            Aoc8::run(Part::One, is_test);
        }
        "8.2" => {
            Aoc8::run(Part::Two, is_test);
        }
        _ => {
            println!("Module \"{}\" not found!", module);
            return;
        }
    };
}
