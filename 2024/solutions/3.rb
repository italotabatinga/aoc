require_relative "problem_base"

class Problem3 < ProblemBase
  def problem = 3

  def build(part)
    @line = []
    input = (part == 1) ? input1 : input2
    @line = input
  end

  def part1
    build(1)

    result = 0
    mul_occurrences = @line.scan(/mul\(\d+,\d+\)/)
    mul_occurrences.each do |mul|
      a, b = mul[4..-2].split(",").map(&:to_i)
      result += a * b
    end
    result
  end

  def part2
    build(2)

    result = 0
    should_do = true
    operations = @line.scan(/(mul\(\d+,\d+\))|(do\(\))|(don't\(\))/).map { |x| x.compact.first }
    operations.each do |op|
      if op == "do()"
        should_do = true
      elsif op == "don't()"
        should_do = false
      else
        next if !should_do

        a, b = op[4..-2].split(",").map(&:to_i)
        result += a * b
      end
    end
    result
  end
end
