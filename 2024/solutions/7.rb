require_relative "problem_base"

class Problem7 < ProblemBase
  def problem = 7

  def build(part)
    @lines = []
    input = (part == 1) ? input1 : input2
    input.each_line do |line|
      result, numbers_str = line.split(":")
      numbers = numbers_str.split(" ").map(&:to_i)
      @lines.push([result.to_i, numbers])
    end
  end

  def try_line(result, numbers, partial, count, ops)
    return partial == result if count == 0

    b = numbers.first
    ops.any? do |op|
      new_partial = partial
      case op
      when "+"
        new_partial += b
      when "*"
        new_partial *= b
      when "||"
        new_partial = (new_partial.to_s + b.to_s).to_i
      end

      if new_partial > result
        false
      else
        try_line(result, numbers[1..], new_partial, count - 1, ops)
      end
    end
  end

  def part1
    build(1)

    calibration_result = 0
    @lines.each do |line|
      result, numbers = line
      calibration_result += result if try_line(result, numbers[1..], numbers.first, numbers.size - 1, ["+", "*"])
    end
    calibration_result
  end

  def part2
    build(2)

    calibration_result = 0
    @lines.each do |line|
      result, numbers = line

      calibration_result += result if try_line(result, numbers[1..], numbers.first, numbers.size - 1, ["+", "*", "||"])
    end
    calibration_result
  end
end
