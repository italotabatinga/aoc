# typed: true

require_relative "../problem"
require "sorbet-runtime"

class Solution1 < Solutionable
  extend T::Sig

  sig { override.returns(Integer) }
  def problem = 1

  sig { override.params(input: String).returns(String) }
  def run_1(input)
    zeros_count = 0
    position = 50
    input.each_line do |line|
      direction = line[0]
      count = line[1..].to_i

      if direction == "R"
        position = (position += count) % 100
      elsif direction == "L"
        position = (position -= count) % 100
      end

      zeros_count += 1 if position.zero?
    end

    zeros_count.to_s
  end

  sig { override.params(input: String).returns(String) }
  def run_2(input)
    zeros_count = 0
    position = 50
    total_positions = 100

    input.each_line do |line|
      direction = line[0]
      count = line[1..].to_i

      zeros_count += count / total_positions
      count %= total_positions

      if direction == "R"
        position += count

        if position >= total_positions
          zeros_count += 1
        end
      elsif direction == "L"
        prev_pos = position
        position -= count
        if prev_pos > 0 && position <= 0
          zeros_count += 1
        end
      end
      position %= total_positions
    end

    zeros_count.to_s
  end
end
