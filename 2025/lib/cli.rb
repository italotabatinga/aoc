# frozen_string_literal: true
# typed: true

require "optparse"
require "sorbet-runtime"

Dir[File.join(__dir__, "solutions", "*.rb")].each { |file| require file }
require_relative "problem"

class CliOptions < T::Struct
  prop :problem, T.nilable(Problem)
  prop :use_example_input, T::Boolean, default: false
end

class Cli
  extend T::Sig

  sig { returns(CliOptions) }
  attr_reader :options

  SOLUTIONS = T.let({
    1 => Solution1.new,
    2 => Solution2.new
  }, T::Hash[Integer, Solutionable])

  def self.run = new.run

  def initialize
    @failed_tests = []
    @options = CliOptions.new

    OptionParser.new do |opts|
      opts.banner = "Usage: aoc.rb [options]"

      opts.on("-r", "--run PROBLEM[.PART]", "Run the specified problem and part") do |v|
        options.problem = Problem.from_str(v)
      end

      opts.on("-e", "--example", "Run the example input for the specified problem and part") do
        options.use_example_input = true
      end
    end.parse!
  end

  def run
    problem = options.problem
    if !problem.nil?
      run_problem
    else
      puts "No valid options provided. Use --help for usage."
    end
  end

  private

  def run_problem
    problem = options.problem
    raise ArgumentError, "expected problem defined" if problem.nil?

    solution = SOLUTIONS[problem.num]
    raise ArgumentError, "solution not defined for problem #{problem}" if solution.nil?

    solution.run(problem.part, options.use_example_input)
  end
end
