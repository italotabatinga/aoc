# typed: true

require_relative "../problem"
require "sorbet-runtime"

class Solution6 < Solutionable
  extend T::Sig

  sig { override.returns(Integer) }
  def problem = 6

  Ops = T.type_alias { T::Array[T::Array[String]] }

  sig { override.params(input: String).returns(String) }
  def run_1(input)
    operations = T.let([], Ops)
    input.each_line.with_index do |line, i|
      elements = line.strip.split(" ")
      elements.each_with_index do |el, j|
        operations.push([]) if i == 0
        T.must(operations[j]).push(el)
      end
    end

    sum = 0
    operations.each do |op|
      math_op = op[-1]
      val = case math_op
      when "*"
        T.must(op[...-1]).reduce(1) { |acc, curr| acc * curr.to_i }
      when "+"
        T.must(op[...-1]).reduce(0) { |acc, curr| acc + curr.to_i }
      end

      sum += T.must(val)
    end

    sum.to_s
  end

  sig { override.params(input: String).returns(String) }
  def run_2(input)
    lines = input.lines
    math_ops_line = T.must(lines[-1])
    math_ops = math_ops_line.scan(/([*+]\s*)/).flatten
    math_ops = math_ops.map.with_index do |ops_str, i|
      next ops_str if i == math_ops.size - 1
      ops_str[...-1]
    end

    cum_col = 0
    sum = 0
    math_ops.each do |ops|
      math = ops.strip
      size = ops.size - 1
      cum = (math == "*") ? 1 : 0
      while size >= 0
        str = ""
        T.must(lines[...-1]).each do |line|
          digit = T.must(line[cum_col + size])
          next if digit == " "
          str += digit
        end
        case math
        when "*"
          cum *= str.to_i
        when "+"
          cum += str.to_i
        end
        size -= 1
      end
      sum += cum
      cum_col += ops.size + 1
    end

    sum.to_s
  end
end
