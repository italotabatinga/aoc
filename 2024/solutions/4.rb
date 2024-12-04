require_relative "problem_base"

class Problem4 < ProblemBase
  def problem = 4

  def build(part)
    @map = []
    input = (part == 1) ? input1 : input2
    input.each_line do |line|
      @map.push(line[0...-1].chars)
    end
    @max_i = @map.size
    @max_j = @map[0].size
  end

  def hv_directions = [[0, 1], [0, -1], [1, 0], [-1, 0]]

  def diag_directions = [[1, 1], [1, -1], [-1, 1], [-1, -1]]

  def directions = hv_directions + diag_directions

  def xmas = ["X", "M", "A", "S"]

  def mas = ["M", "A", "S"]

  def valid_pos?(pos_i, pos_j) = pos_i >= 0 && pos_i < @max_i && pos_j >= 0 && pos_j < @max_j

  def has_xmas?(pos_i, pos_j, dir_i, dir_j)
    xmas.each do |c|
      return false if !valid_pos?(pos_i, pos_j) || @map[pos_i][pos_j] != c

      pos_i += dir_i
      pos_j += dir_j
    end

    true
  end

  def has_mas?(pos_i, pos_j, dir_i, dir_j)
    pos_i -= dir_i
    pos_j -= dir_j

    mas.each do |c|
      return false if !valid_pos?(pos_i, pos_j) || @map[pos_i][pos_j] != c

      pos_i += dir_i
      pos_j += dir_j
    end

    true
  end

  def part1
    build(1)

    result = 0
    @map.each_with_index do |line, i|
      line.each_with_index do |_, j|
        directions.each do |dir|
          result += 1 if has_xmas?(i, j, *dir)
        end
      end
    end
    result
  end

  def part2
    build(2)

    result = 0
    @map.each_with_index do |line, i|
      line.each_with_index do |_, j|
        count_mas = diag_directions.sum do |dir|
          has_mas?(i, j, *dir) ? 1 : 0
        end
        result += 1 if count_mas == 2
      end
    end
    result
  end
end
