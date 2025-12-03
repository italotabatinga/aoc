# typed: true

require_relative "../problem"
require "sorbet-runtime"

class Solution3 < Solutionable
  extend T::Sig

  def initialize
    @valids = T.let({}, T::Hash[Integer, T::Boolean])
  end

  sig { override.returns(Integer) }
  def problem = 3

  sig { override.params(input: String).returns(String) }
  def run_1(input)
    joltages = T.let([], T::Array[Integer])

    input.each_line do |bank|
      joltage = find_largest_joltage(bank, size: 2)
      joltages.push(joltage)
    end

    joltages.sum.to_s
  end

  sig { override.params(input: String).returns(String) }
  def run_2(input)
    joltages = T.let([], T::Array[Integer])

    input.each_line do |bank|
      joltage = find_largest_joltage(bank, size: 12)
      joltages.push(joltage)
    end

    joltages.sum.to_s
  end

  sig { params(bank: String, size: Integer).returns(Integer) }
  def find_largest_joltage(bank, size:)
    batteries = bank.strip.chars
    chosen_batteries = ""
    left_index = 0
    while size > 0
      max = T.cast(batteries[left_index..(-size)]&.max_by { |d| d.to_i }, String)
      left_index += T.cast(batteries[left_index..(-size)]&.find_index { |d| d == max }, Integer) + 1

      chosen_batteries += max
      size -= 1
    end

    chosen_batteries.to_i
  end
end
