use std::ops::Range;

use crate::Aoc;

pub struct Aoc16 {}

impl Aoc<PacketChar, usize> for Aoc16 {
    fn input() -> String {
        String::from(include_str!("inputs/16.txt"))
    }

    fn part1_test_input() -> String {
        String::from(include_str!("inputs/16_test.txt"))
    }

    fn part2_test_input() -> String {
        String::from(include_str!("inputs/16_test_2.txt"))
    }

    fn parse_input(s: String) -> PacketChar {
        let mut vec = Vec::new();
        for c in s.chars() {
            vec.extend(hex_char_into_bit_vec(&c));
        }

        vec
    }

    fn part1(packetchar: PacketChar, _: bool) -> usize {
        let (packets, _) = iterate_on_packet(&packetchar, 0..packetchar.len(), 1);
        sum_versions(&packets)
    }

    fn part2(packetchar: PacketChar, _: bool) -> usize {
        let (packets, _) = iterate_on_packet(&packetchar, 0..packetchar.len(), 1);
        evaluate_packet(&packets[0])
    }
}

type PacketChar = Vec<char>;

#[derive(Debug)]
enum Packet {
    Literal {
        version: usize,
        type_id: usize,
        value: usize,
    },
    Op {
        version: usize,
        type_id: usize,
        length_type: usize,
        sub_packets: Vec<Packet>,
    },
}

#[derive(Debug)]
enum BitOp {
    Version,
    Type,
    Literal,
    LengthType,
    TotalLength,
    NumSubPackets,
}

fn sum_versions(packets: &Vec<Packet>) -> usize {
    packets.iter().fold(0, |sum, packet| match packet {
        Packet::Literal { version, .. } => sum + version,
        Packet::Op {
            version,
            sub_packets,
            ..
        } => sum + version + sum_versions(sub_packets),
    })
}

fn evaluate_packet(packet: &Packet) -> usize {
    match packet {
        Packet::Literal { value, .. } => *value,
        Packet::Op {
            type_id,
            sub_packets,
            ..
        } => match type_id {
            0 => sub_packets
                .iter()
                .fold(0, |sum, spkt| sum + evaluate_packet(spkt)),
            1 => sub_packets
                .iter()
                .fold(1, |sum, spkt| sum * evaluate_packet(spkt)),
            2 => sub_packets
                .iter()
                .fold(usize::MAX, |min, spkt| min.min(evaluate_packet(spkt))),
            3 => sub_packets
                .iter()
                .fold(usize::MIN, |max, spkt| max.max(evaluate_packet(spkt))),
            5 => {
                if evaluate_packet(&sub_packets[0]) > evaluate_packet(&sub_packets[1]) {
                    1
                } else {
                    0
                }
            }
            6 => {
                if evaluate_packet(&sub_packets[0]) < evaluate_packet(&sub_packets[1]) {
                    1
                } else {
                    0
                }
            }
            7 => {
                if evaluate_packet(&sub_packets[0]) == evaluate_packet(&sub_packets[1]) {
                    1
                } else {
                    0
                }
            }
            _ => 0,
        },
    }
}

fn iterate_on_packet(
    packet: &PacketChar,
    range: Range<usize>,
    mut remaining_packets: i32,
) -> (Vec<Packet>, usize) {
    let mut next_bits_op = BitOp::Version;
    let mut i = range.start;
    let mut version = 0;
    let mut type_id = 0;
    let mut length_type = 0;
    let mut literal = 0;
    let mut packets = Vec::new();
    while i < range.end && remaining_packets > 0 {
        match next_bits_op {
            BitOp::Version => {
                version =
                    usize::from_str_radix(&packet[i..i + 3].iter().collect::<String>(), 2).unwrap();
                next_bits_op = BitOp::Type;
                i += 3;
            }
            BitOp::Type => {
                type_id =
                    usize::from_str_radix(&packet[i..i + 3].iter().collect::<String>(), 2).unwrap();
                match type_id {
                    4 => {
                        next_bits_op = BitOp::Literal;
                    }
                    _ => next_bits_op = BitOp::LengthType,
                }
                i += 3;
            }
            BitOp::Literal => {
                let num =
                    usize::from_str_radix(&packet[i + 1..i + 5].iter().collect::<String>(), 2)
                        .unwrap();
                literal <<= 4;
                literal += num;
                if packet[i] == '0' {
                    next_bits_op = BitOp::Version;
                    packets.push(Packet::Literal {
                        version: version,
                        type_id,
                        value: literal,
                    });
                    remaining_packets -= 1;
                    literal = 0;
                }
                i += 5;
            }
            BitOp::LengthType => {
                length_type =
                    usize::from_str_radix(&packet[i..i + 1].iter().collect::<String>(), 2).unwrap();
                match length_type {
                    0 => next_bits_op = BitOp::TotalLength,
                    1 => next_bits_op = BitOp::NumSubPackets,
                    _ => {}
                }
                i += 1;
            }
            BitOp::TotalLength => {
                let total_length =
                    usize::from_str_radix(&packet[i..i + 15].iter().collect::<String>(), 2)
                        .unwrap();

                let (sub_packets, next_i) =
                    iterate_on_packet(packet, (i + 15)..(i + 15 + total_length), i32::MAX);
                packets.push(Packet::Op {
                    version,
                    type_id,
                    length_type,
                    sub_packets,
                });
                remaining_packets -= 1;
                next_bits_op = BitOp::Version;

                i = next_i
            }
            BitOp::NumSubPackets => {
                let total_length =
                    i32::from_str_radix(&packet[i..i + 11].iter().collect::<String>(), 2).unwrap();
                let (sub_packets, next_i) =
                    iterate_on_packet(packet, (i + 11)..range.end, total_length);

                packets.push(Packet::Op {
                    version,
                    type_id,
                    length_type,
                    sub_packets,
                });
                remaining_packets -= 1;
                next_bits_op = BitOp::Version;

                i = next_i
            }
        }
    }
    (packets, i)
}

fn hex_char_into_bit_vec(c: &char) -> Vec<char> {
    match c {
        '0' => vec!['0', '0', '0', '0'],
        '1' => vec!['0', '0', '0', '1'],
        '2' => vec!['0', '0', '1', '0'],
        '3' => vec!['0', '0', '1', '1'],
        '4' => vec!['0', '1', '0', '0'],
        '5' => vec!['0', '1', '0', '1'],
        '6' => vec!['0', '1', '1', '0'],
        '7' => vec!['0', '1', '1', '1'],
        '8' => vec!['1', '0', '0', '0'],
        '9' => vec!['1', '0', '0', '1'],
        'A' => vec!['1', '0', '1', '0'],
        'B' => vec!['1', '0', '1', '1'],
        'C' => vec!['1', '1', '0', '0'],
        'D' => vec!['1', '1', '0', '1'],
        'E' => vec!['1', '1', '1', '0'],
        'F' => vec!['1', '1', '1', '1'],
        _ => vec![],
    }
}

#[cfg(test)]
mod tests {
    use crate::{Aoc, Aoc16};

    #[test]
    fn aoc161_test() {
        assert_eq!(Aoc16::run(crate::Part::One, true), 31);
    }

    #[test]
    fn aoc161() {
        assert_eq!(Aoc16::run(crate::Part::One, false), 953);
    }

    #[test]
    fn aoc162_test() {
        assert_eq!(Aoc16::run(crate::Part::Two, true), 1);
    }

    #[test]
    fn aoc162() {
        assert_eq!(Aoc16::run(crate::Part::Two, false), 246225449979);
    }
}
