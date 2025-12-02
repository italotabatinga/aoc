# typed: true
# frozen_string_literal: true

require "sorbet-runtime"

class ProblemPart < T::Enum
  enums do
    First = new
    Second = new
  end

  def to_s
    case self
    when First
      "1"
    when Second
      "2"
    else T.absurd(self)
    end
  end

  #: (String?) -> ProblemPart
  def self.from_str(str)
    case str
    when "1"
      First
    when "2"
      Second
    else
      raise ArgumentError, "expected 'part' to be 1 or 2"
    end
  end
end

class Problem < T::Struct
  const :num, Integer
  const :part, ProblemPart

  #: (String?) -> Problem?
  def self.from_str(problem_string)
    return if problem_string.nil?

    num_str, part_str = problem_string.split(".")
    num = num_str.to_i
    part = ProblemPart.from_str(part_str)

    new(num:, part:)
  end
end

class Solutionable
  extend T::Sig
  extend T::Helpers

  abstract!

  sig { abstract.returns(Integer) }
  def problem
  end

  sig { abstract.params(input: String).returns(String) }
  def run_1(input)
  end

  sig { abstract.params(input: String).returns(String) }
  def run_2(input)
  end

  sig { params(part: ProblemPart, use_example_input: T::Boolean).returns(String) }
  def run(part, use_example_input)
    input = read_input(part, use_example_input)

    case part
    when ProblemPart::First
      run_1(input)
    when ProblemPart::Second
      run_2(input)
    else T.absurd(part)
    end

    # puts "Problem #{problem}.#{part} - #{result}"
  end

  sig { params(part: ProblemPart, use_example_input: T::Boolean).returns(String) }
  def read_input(part, use_example_input)
    file_name_base = File.join("inputs", problem.to_s + (use_example_input ? "_example.txt" : ".txt"))
    file_name_with_part = File.join("inputs", "#{problem}_#{part}" + (use_example_input ? "_example.txt" : ".txt"))

    if File.exist?(file_name_with_part)
      File.read(file_name_with_part)
    elsif File.exist?(file_name_base)
      File.read(file_name_base)
    else
      raise "file not found, expected #{file_name_with_part} or #{file_name_base} to be defined"
    end
  end
end
