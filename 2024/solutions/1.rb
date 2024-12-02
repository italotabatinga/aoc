require_relative "problem_base"

class Problem1 < ProblemBase
  def problem = 1

  def build_lists(part)
    @left = []
    @right = []
    input = (part == 1) ? input1 : input2
    input.each_line do |line|
      a, b = line.split(" ")
      @left.push(a.to_i)
      @right.push(b.to_i)
    end
  end

  def sort_lists
    @left.sort!
    @right.sort!
  end

  def part1
    build_lists(1)
    sort_lists

    @left.sort!
    @right.sort!

    sum = 0
    @left.zip(@right).each do |l, r|
      sum += (l - r).abs
    end
    sum
  end

  def part2
    build_lists(2)
    r_occurrences = {}
    similarities = {}

    @right.each do |r|
      r_occurrences[r] ||= 0
      r_occurrences[r] += 1
    end

    similarity_score = 0
    @left.each do |l|
      known_similarity = similarities[l]
      if known_similarity.nil?
        r_occurrence = r_occurrences[l] || 0
        known_similarity = l * r_occurrence
        similarities[l] = known_similarity
      end

      similarity_score += known_similarity
    end
    similarity_score
  end
end
