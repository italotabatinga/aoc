# typed: true

require_relative "../problem"
require "sorbet-runtime"

class Solution4 < Solutionable
  extend T::Sig

  EMPTY = "."
  ROLL = "@"

  DIRECTIONS = T.let([
    [-1, -1],
    [0, -1],
    [1, -1],
    [-1, 0],
    [1, 0],
    [-1, 1],
    [0, 1],
    [1, 1]
  ], T::Array[[Integer, Integer]])

  Grid = T.type_alias { T::Array[T::Array[String]] }

  def initialize
    @valids = T.let({}, T::Hash[Integer, T::Boolean])
  end

  sig { override.returns(Integer) }
  def problem = 4

  sig { override.params(input: String).returns(String) }
  def run_1(input)
    grid = T.let([], Grid)

    input.each_line do |line|
      rolls = line.strip.chars
      grid.push(rolls)
    end

    count = remove_rolls(grid)

    count.to_s
  end

  sig { override.params(input: String).returns(String) }
  def run_2(input)
    grid = T.let([], Grid)

    input.each_line do |line|
      rolls = line.strip.chars
      grid.push(rolls)
    end

    count = 0
    loop do
      tmp_count = remove_rolls(grid)
      break if tmp_count == 0

      count += tmp_count
    end

    count.to_s
  end

  sig { params(grid: Grid).returns(Integer) }
  def remove_rolls(grid)
    to_be_removed = T.let([], T::Array[[Integer, Integer]])
    grid.each_with_index do |row, j|
      row.each_with_index do |char, i|
        next if char != ROLL

        count_neighbors = 0
        DIRECTIONS.each do |x, y|
          new_i = i + x
          new_j = j + y

          if new_i >= 0 && new_i < row.size && new_j >= 0 && new_j < grid.size
            count_neighbors += 1 if T.must(grid[new_j])[new_i] == ROLL
          end
        end

        if count_neighbors < 4
          to_be_removed.push([i, j])
        end
      end
    end

    to_be_removed.each do |i, j|
      T.must(grid[j])[i] = EMPTY
    end

    to_be_removed.size
  end
end
