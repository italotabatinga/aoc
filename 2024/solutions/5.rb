require_relative "problem_base"

class Problem5 < ProblemBase
  def problem = 5

  def build(part)
    @order_map = {}
    @lines = []
    input = (part == 1) ? input1 : input2
    reading_map = true
    input.each_line do |line|
      if reading_map
        if line.size == 1 && line[0] == "\n"
          reading_map = false
          next
        end

        a, b = line[0...-1].split("|").map(&:to_i)
        @order_map[a] ||= []
        @order_map[a].push(b)
      else
        @lines.push(line[0...-1].split(",").map(&:to_i))
      end
    end
  end

  def valid_update?(update)
    update.each_with_index do |page, i|
      update[i + 1..].each do |other_page|
        if @order_map[other_page]&.include?(page)
          return false
        end
      end
    end

    true
  end

  def part1
    build(1)

    result = 0
    @lines.each do |update|
      if valid_update?(update)
        result += update[update.size / 2]
      end
    end
    result
  end

  def part2
    build(2)

    result = 0
    @lines.each do |update|
      if !valid_update?(update)
        sorted_update = update.sort do |a, b|
          if @order_map[a]&.include?(b)
            1
          elsif @order_map[b]&.include?(a)
            -1
          else
            0
          end
        end
        result += sorted_update[sorted_update.size / 2]
      end
    end
    result
  end
end
