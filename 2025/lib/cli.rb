# typed: true
# frozen_string_literal: true

require "optparse"
require "sorbet-runtime"

Dir[File.join(__dir__, "solutions", "*.rb")].each { |file| require file }
require_relative "problem"

SOLUTIONS = T.let({
  1 => Solution1,
  2 => Solution2,
  3 => Solution3
}, T::Hash[Integer, T.class_of(Solutionable)])

ANSWERS = T.let({
  1 => ["1102", "6175"],
  2 => ["12599655151", "20942028255"],
  3 => ["17193", "171297349921310"]
}, T::Hash[Integer, [String, String]])

#: (Problem, T::Boolean) -> String
def run_problem(problem, use_example_input)
  solution_class = SOLUTIONS[problem.num]
  raise ArgumentError, "solution not defined for problem #{problem}" if solution_class.nil?

  solution = solution_class.new
  solution.run(problem.part, use_example_input)
end

class Commandable
  extend T::Sig
  extend T::Helpers

  abstract!

  sig { abstract.params(args: T::Array[String]).void }
  def parse(args)
  end

  sig { abstract.void }
  def execute
  end

  sig { abstract.returns(String) }
  def help
  end

  sig { abstract.returns(String) }
  def description
  end
end

class RunCommand < Commandable
  extend T::Sig

  class Options < T::Struct
    prop :problem, T.nilable(Problem)
    prop :use_example_input, T::Boolean, default: false
  end

  def initialize
    @options = T.let(Options.new, Options)
  end

  sig { override.returns(String) }
  def description = "Runs a given problem"

  sig { override.returns(String) }
  def help
    "Usage: cli.rb run PROBLEM.PART [options]\n\n" \
    "Options:\n" \
    "  -e, --example    Use example input\n" \
    "  -h, --help       Show this help message"
  end

  sig { override.params(args: T::Array[String]).void }
  def parse(args)
    OptionParser.new do |opts|
      opts.banner = "Usage: cli.rb run PROBLEM.PART [options]"

      opts.on("-e", "--example", "Use example input") do
        @options.use_example_input = true
      end

      opts.on("-h", "--help", "Show this help message") do
        puts help
        exit
      end
    end.parse!(args)

    if args.empty?
      puts "Error: PROBLEM.PART is required"
      puts help
      exit 1
    end

    problem_str = args[0]
    begin
      @options.problem = Problem.from_str(problem_str)
    rescue ArgumentError => e
      puts "Error: #{e.message}"
      exit 1
    end
  end

  sig { override.void }
  def execute
    problem = @options.problem
    if problem.nil?
      puts "Error: No problem specified"
      exit 1
    end

    result = run_problem(problem, @options.use_example_input)
    puts result
  end
end

class TestCommand < Commandable
  extend T::Sig

  sig { override.returns(String) }
  def description = "Tests every problem defined"

  sig { override.void }
  def execute
    puts "Running all tests..."
    ANSWERS.each do |num, (part_1_expected, part_2_expected)|
      solution_class = SOLUTIONS[num]
      next if solution_class.nil?

      solution = solution_class.new

      print "Problem #{num}.1..."
      part_1_got = solution.run(ProblemPart::First, false)
      if part_1_got == part_1_expected
        puts "\e[32mok\e[0m"
      else
        puts "\e[31mfailed\e[0m (expected: #{part_1_expected}, got: #{part_1_got})"
      end

      solution = solution_class.new
      print "Problem #{num}.2..."
      part_2_got = solution.run(ProblemPart::Second, false)
      if part_2_got == part_2_expected
        puts "\e[32mok\e[0m"
      else
        puts "\e[31mfailed\e[0m (expected: #{part_2_expected}, got: #{part_2_got})"
      end
    end
  end

  sig { override.returns(String) }
  def help
    "Usage: cli.rb test [options]\n\n" \
    "Options:\n" \
    "  -h, --help       Show this help message"
  end

  sig { override.params(args: T::Array[String]).void }
  def parse(args)
    OptionParser.new do |opts|
      opts.banner = "Usage: cli.rb test [options]"

      opts.on("-h", "--help", "Show this help message") do
        puts help
        exit
      end
    end.parse!(args)
  end
end

class Cli
  extend T::Sig

  sig { returns(T.nilable(Commandable)) }
  attr_reader :command

  Commands = T.let({
    "run" => RunCommand.new,
    "test" => TestCommand.new
  }, T::Hash[String, Commandable])

  def self.execute = new.execute

  def initialize
    @command = T.let(nil, T.nilable(Commandable))

    first_command_index = ARGV.find_index { |arg| !arg.start_with?("-", "--") }

    # Parse global options before the command
    global_args = first_command_index.nil? ? ARGV : ARGV[...first_command_index]

    OptionParser.new do |opts|
      opts.banner = "Usage: cli.rb [options] [command]"

      opts.on("-h", "--help", "Prints this help") do
        help
        exit
      end
    end.parse(global_args)

    # Get command name
    command_str = first_command_index.nil? ? nil : ARGV[first_command_index]

    if command_str.nil?
      help
      exit
    end

    if Commands[command_str].nil?
      puts "Error: Unknown command '#{command_str}'"
      help
      exit 1
    end

    @command = Commands[command_str]

    # Parse command-specific arguments
    command_args = ARGV[(T.must(first_command_index) + 1)..]
    T.must(@command).parse(T.must(command_args))
  end

  def execute
    if @command.nil?
      help
      exit 1
    end

    @command.execute
  end

  def help
    string = "\e[1mUsage:\e[0m cli.rb [options] [command]\n\n" \
             "\e[1mCommands:\e[0m\n"
    commands_section = Commands.map do |name, command|
      "  #{name.ljust(8)} #{command.description}"
    end.join("\n")
    string += commands_section

    puts string
  end
end
