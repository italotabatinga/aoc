class ProblemBase
  attr_reader :answers

  def initialize(is_example: false)
    @answers = ["", ""]
    @runned_parts = [false, false]
    @inputs = [nil, nil]
    @is_example = is_example
  end

  def run(part = nil)
    if part.nil?
      internal_part1
      internal_part2
    elsif part == "1"
      internal_part1
    elsif part == "2"
      internal_part2
    else
      raise ArgumentError, "Unexpected input for part: #{part}"
    end
  end

  def problem
    raise NotImplementedError, "This method must be implemented in a subclass"
  end

  def part1
    raise NotImplementedError, "This method must be implemented in a subclass"
  end

  def part2
    raise NotImplementedError, "This method must be implemented in a subclass"
  end

  def input1 = @inputs[0]

  def input2 = @inputs[1]

  def print
    @runned_parts.each_with_index do |has_run, part|
      puts "Solution #{problem}.#{part + 1}: #{@answers[part]}" if has_run
    end
  end

  private

  def internal_part1
    if input1.nil?
      load_input(0)
    end
    @answers[0] = part1.to_s
    @runned_parts[0] = true
    @answers[0]
  end

  def internal_part2
    if input2.nil?
      load_input(1)
    end

    @answers[1] = part2.to_s
    @runned_parts[1] = true
    @answers[1]
  end

  def load_input(part)
    is_example_suffix = @is_example ? "_example" : ""

    begin
      file_path = "inputs/#{problem}_#{part + 1}#{is_example_suffix}.txt"
      if !File.exist?(file_path)
        file_path = "inputs/#{problem}#{is_example_suffix}.txt"
        raise "Input file not found for #{problem}.#{part + 1}#{is_example_suffix}" unless File.exist?(file_path)
      end
      @inputs[part] = File.read(file_path)
    rescue => e
      raise "Failed to load input: #{e.message}"
    end
  end
end
