# typed: true

require_relative "../problem"
require "sorbet-runtime"

class Solution5 < Solutionable
  extend T::Sig

  sig { override.returns(Integer) }
  def problem = 5

  class Range < T::Struct
    extend T::Sig

    const :left, Integer
    const :right, Integer

    sig { returns(String) }
    def to_s = "R(#{left}, #{right})"

    sig { returns(String) }
    def inspect = to_s
  end

  def initialize
    @valids = T.let({}, T::Hash[Integer, T::Boolean])
  end

  sig { override.params(input: String).returns(String) }
  def run_1(input)
    ranges_str, ingredients_str = input.split("\n\n")
    ranges = T.let([], T::Array[[Integer, Integer]])
    T.cast(ranges_str, String).each_line do |line|
      left, right = line.strip.split("-")
      ranges.push([left.to_i, right.to_i])
    end

    count = 0
    T.cast(ingredients_str, String).each_line do |line|
      ingredient = line.strip.to_i
      count += 1 if is_fresh?(ingredient, ranges)
    end

    count.to_s
  end

  sig { override.params(input: String).returns(String) }
  def run_2(input)
    ranges_str, _ingredients_str = input.split("\n\n")
    ranges = T.let([], T::Array[Range])
    T.cast(ranges_str, String).each_line do |line|
      left, right = line.strip.split("-")
      ranges.push(Range.new(left: left.to_i, right: right.to_i))
    end

    ranges = ranges.uniq { |r| r.to_s }.sort_by { |r| [r.left, r.right] }

    cur_l = T.must(ranges.first).left
    cur_r = T.must(ranges.first).right
    count = 0
    T.must(ranges[1..]).each do |r|
      if r.left > cur_r + 1
        count += cur_r - cur_l + 1
        cur_l = r.left
        cur_r = r.right
      else
        cur_r = [cur_r, r.right].max
      end
    end
    count += cur_r - cur_l + 1

    count.to_s
  end

  sig { params(ingredient: Integer, ranges: T::Array[[Integer, Integer]]).returns(T::Boolean) }
  def is_fresh?(ingredient, ranges)
    ranges.each do |left, right|
      return true if ingredient.between?(left, right)
    end

    false
  end
end
