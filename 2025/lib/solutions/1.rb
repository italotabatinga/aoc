# typed: true

require_relative "../problem"
require "sorbet-runtime"

class Solution1 < Solutionable
  extend T::Sig

  sig { override.returns(Integer) }
  def problem = 1

  sig { override.params(use_example_input: T::Boolean).returns(String) }
  def run_1(use_example_input)
    "1"
  end

  sig { override.params(use_example_input: T::Boolean).returns(String) }
  def run_2(use_example_input)
    "2"
  end
end
