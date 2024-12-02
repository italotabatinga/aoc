#!/usr/bin/env ruby
require "optparse"

class AocTest
  attr_reader :problem, :part, :expected, :exception

  def initialize(problem, part, expected)
    @problem = problem
    @part = part
    @expected = expected
    @actual = nil
    @success = nil
    @exception = nil
  end

  def run(&block)
    @actual = block.call
    @success = @actual == @expected
  rescue => e
    @exception = e
    @success = false
  end

  def pass? = @success

  def failed? = !@success

  def to_s
    if exception
      "Test #{problem}.#{part}: ERROR\n#{exception}"
    else
      "Test #{problem}.#{part}: Expected #{@expected}, but got #{@actual}"
    end
  end
end

class AdventOfCodeCLI
  def initialize
    @options = {}
    @failed_tests = []
    OptionParser.new do |opts|
      opts.banner = "Usage: aoc.rb [options]"

      opts.on("-r", "--run PROBLEM[.PART]", "Run the specified problem and part") do |v|
        @options[:run] = v
      end

      opts.on("-t", "--test [PROBLEM.PART]", "Test the specified problem and part or all if not specified") do |v|
        @options[:test] = v || "all"
      end

      opts.on("-e", "--example", "Run the example input for the specified problem and part") do
        @options[:example] = true
      end
    end.parse!
  end

  def run
    if @options[:run]
      problem, part = @options[:run].split(".")
      run_problem(problem, part, @options[:example])
    elsif @options[:test]
      if @options[:test] == "all"
        test_problems
      else
        test_problem(@options[:test])
      end
    else
      puts "No valid options provided. Use --help for usage."
    end
  end

  private

  def run_problem(problem, part = nil, is_example = false)
    require_relative "solutions/#{problem}"
    solution = Object.const_get("Problem#{problem.capitalize}").new(is_example: is_example)
    solution.run(part)
    solution.print
  end

  def test_problems(problem_string = nil)
    expected_answers = File.read("outputs.txt").split("\n").map { |line| line.split(" ") }.to_h
    if problem_string
      problem, part = problem_string.split(".")
      expected = expected_answers["#{problem}.#{part}"]
      raise "Output not defined for #{problem_string}" if expected.nil?

      test_problem(problem, part, expected)
    else
      expected_answers.each do |problem_string, expected|
        problem, part = problem_string.split(".")
        test_problem(problem, part, expected)
      end
    end

    puts
    @failed_tests.each { |t| puts t }
  end

  def test_problem(problem, part, expected)
    require_relative "solutions/#{problem}"
    solution = Object.const_get("Problem#{problem.capitalize}").new
    test = AocTest.new(problem, part, expected)
    test.run { solution.run(part) }
    @failed_tests.push(test) if test.failed?

    if test.failed?
      print "\e[31m\u2717\e[0m"  # Red cross
    else
      print "\e[32m\u25CF\e[0m"  # Green dot
    end
  end
end

AdventOfCodeCLI.new.run
