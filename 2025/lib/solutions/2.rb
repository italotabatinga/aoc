# typed: true

require_relative "../problem"
require "sorbet-runtime"

class Solution2 < Solutionable
  extend T::Sig

  def initialize
    @valids = T.let({}, T::Hash[Integer, T::Boolean])
  end

  sig { override.returns(Integer) }
  def problem = 2

  sig { override.params(input: String).returns(String) }
  def run_1(input)
    count_invalid = 0
    sum_invalid = 0
    ranges = input.split(",")
    ranges.each do |range|
      left, right = range.split("-").map(&:to_i)

      (left..right).each do |val|
        is_invalid = check_invalid(val, false)
        if is_invalid
          count_invalid += 1
          sum_invalid += val
        end
      end
    end

    sum_invalid.to_s
  end

  sig { override.params(input: String).returns(String) }
  def run_2(input)
    count_invalid = 0
    sum_invalid = 0
    ranges = input.split(",")
    ranges.each do |range|
      left, right = range.split("-").map(&:to_i)

      (left..right).each do |val|
        is_invalid = check_invalid(val, true)
        if is_invalid
          count_invalid += 1
          sum_invalid += val
        end
      end
    end

    sum_invalid.to_s
  end

  sig { params(val: Integer, count_repeated: T::Boolean).returns(T::Boolean) }
  def check_invalid(val, count_repeated)
    cached = @valids[val]
    return cached if !cached.nil?

    res = if count_repeated
      calc_invalid_repeated(val)
    else
      calc_invalid_twice(val)
    end
    @valids[val] = res
    res
  end

  sig { params(val: Integer).returns(T::Boolean) }
  def calc_invalid_twice(val)
    val = val.to_s
    size = val.size
    if size % 2 != 0
      return false
    end
    seq_size = size / 2
    eq = T.let(true, T::Boolean)
    pos = seq_size
    if val[0...seq_size] != val[pos...(pos + seq_size)]
      eq = false
    end

    eq
  end

  sig { params(val: Integer).returns(T::Boolean) }
  def calc_invalid_repeated(val)
    val = val.to_s
    seq_size = 1
    size = val.size
    while seq_size * 2 <= size
      pos = seq_size
      if size % seq_size == 0
        eq = T.let(true, T::Boolean)
        while pos + seq_size <= size
          if val[0...seq_size] == val[pos...(pos + seq_size)]
          else
            eq = false
            break
          end
          pos += seq_size
        end

        if eq
          return true
        end
      end

      seq_size += 1
    end

    false
  end
end
