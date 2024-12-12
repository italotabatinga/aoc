require_relative "problem_base"

class Problem11 < ProblemBase
  def problem = 11

  def build(part)
    @stones_hash = {}
    input = (part == 1) ? input1 : input2
    input.strip.split(" ").map(&:to_i).each do |stone|
      @stones_hash[stone] ||= 0
      @stones_hash[stone] += 1
    end
  end

  def blink
    new_hash = {}

    @stones_hash.each do |stone, quantity|
      if stone == 0
        new_hash[1] ||= 0
        new_hash[1] += quantity
        next
      end

      stone_str = stone.to_s
      length = stone_str.size

      if length % 2 == 0
        a = stone_str[...length / 2].to_i
        b = stone_str[length / 2..].to_i
        new_hash[a] ||= 0
        new_hash[a] += quantity
        new_hash[b] ||= 0
        new_hash[b] += quantity
      else
        a = stone * 2024
        new_hash[a] ||= 0
        new_hash[a] += quantity
      end
    end

    @stones_hash = new_hash
  end

  def stones_count = @stones_hash.values.sum

  def part1
    build(1)

    25.times.each { blink }
    stones_count
  end

  def part2
    build(2)

    75.times.each { blink }
    stones_count
  end
end
