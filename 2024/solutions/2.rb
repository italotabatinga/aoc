require_relative 'problem_base'

class Problem2 < ProblemBase
  def problem = 2

  def build(part)
    @reports = []
    input = part == 1 ? input1 : input2
    input.each_line do |line|
      @reports.push(line.split(' ').map(&:to_i))
    end
  end

  def is_safe(report, test_bad_level = false)
    safe = true
    last = nil
    report.each_cons(2) do |a, b|
      diff = (a - b).abs
      if (diff < 1 || diff > 3) || (last != nil && ((b >= a && a <= last) || (b <= a && a >= last)))
        safe = false
        tentative_safe = false
        if test_bad_level
          report.each_with_index do |_, i|
            tentative_report = report.dup.tap{ |x| x.delete_at(i) }
            tentative_safe = is_safe(tentative_report)

            break if tentative_safe
          end

          safe ||= tentative_safe
        end

        break if !safe
      end

      last = a
    end

    safe
  end

  def part1
    build(1)

    count_safe = 0
    @reports.each do |r|
      safe = is_safe(r)
      count_safe += 1 if safe
    end
    count_safe
  end

  def part2
    build(2)

    count_safe = 0
    @reports.each do |r|
      safe = is_safe(r, true)
      count_safe += 1 if safe
    end
    count_safe
  end
end
