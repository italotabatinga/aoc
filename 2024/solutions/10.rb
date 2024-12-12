require_relative "problem_base"
require_relative "collections"

class Problem10 < ProblemBase
  def problem = 10

  def build(part)
    @map = []
    @memoized = []
    @trailheads = []
    input = (part == 1) ? input1 : input2
    input.each_line.each_with_index do |line, i|
      vec = []
      vec_mem = []
      line[...-1].chars.each_with_index do |char, j|
        vec.push(char.to_i)
        vec_mem.push(nil)

        @trailheads.push(Vec2.new(i, j)) if char == "0"
      end
      @map.push(vec)
      @memoized.push(vec_mem)
    end
  end

  def pp(map)
    map.each do |row|
      row_str = row.map do |cell|
        if cell.is_a?(Set) || cell.is_a?(Array)
          cell.nil? ? " ." : cell.size.to_s.rjust(2)
        else
          cell.nil? ? " ." : cell.to_s.rjust(2)
        end
      end.join(" ")
      puts row_str
    end
  end

  def get(pos)
    return nil if pos.i < 0 || pos.i >= @map.size || pos.j < 0 || pos.j >= @map.first.size

    @map[pos.i][pos.j]
  end

  def mem(pos)
    return nil if pos.i < 0 || pos.i >= @map.size || pos.j < 0 || pos.j >= @map.first.size

    @memoized[pos.i][pos.j]
  end

  def add_mem(pos, other)
    @memoized[pos.i][pos.j] ||= []
    @memoized[pos.i][pos.j].push(*other)
  end

  def get_available_positions(pos)
    positions = []
    [Vec2.new(0, 1), Vec2.new(0, -1), Vec2.new(1, 0), Vec2.new(-1, 0)].each do |dpos|
      tentative = get(pos + dpos)
      next if tentative.nil?

      positions.push(pos + dpos) if tentative - get(pos) == 1
    end
    positions
  end

  def hike(pos)
    memoized = mem(pos)
    return memoized if !memoized.nil?

    if get(pos) == 9
      add_mem(pos, [pos])
      return mem(pos)
    end

    positions = get_available_positions(pos)
    return [] if positions.empty?

    positions.each do |next_pos|
      add_mem(pos, hike(next_pos))
    end
    mem(pos)
  end

  def part1
    build(1)

    @trailheads.each do |head|
      hike(head)
    end
    @trailheads.map { |head| mem(head).uniq.size }.sum
  end

  def part2
    build(2)

    @trailheads.each do |head|
      hike(head)
    end
    @trailheads.map { |head| mem(head).size }.sum
  end
end
