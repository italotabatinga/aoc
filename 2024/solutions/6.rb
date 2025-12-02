require_relative "problem_base"
require_relative "collections"

class Problem6 < ProblemBase
  def problem = 6

  def build(part)
    @blockers = Set.new([])
    @visited = Set.new([])
    @rotations = []
    @loop_blockers = Set.new([])
    @guard_initial = nil
    @guard = nil
    @curr_direction_index = 0
    @max_x = 0
    @max_y = 0
    input = (part == 1) ? input1 : input2
    input.each_line.each_with_index do |line, y|
      line.chars.each_with_index do |char, x|
        case char
        when "^"
          @guard = Vec2.new(x, y)
          @guard_initial = Vec2.new(x, y)
        when "#"
          @blockers.add(Vec2.new(x, y))
        end

        @max_x = x if x > @max_x
      end

      @max_y = y if y > @max_y
    end
    @visited.add(@guard)
  end

  def rotate
    @rotations.push(@guard)
    @curr_direction_index = (@curr_direction_index + 1) % 4
  end

  def directions = [Vec2.new(0, -1), Vec2.new(1, 0), Vec2.new(0, 1), Vec2.new(-1, 0)]

  def direction = directions[@curr_direction_index]

  def next_direction = directions[(@curr_direction_index + 1) % 4]

  def walk
    next_guard = @guard + direction

    n = 3
    while @rotations.size >= n
      last_rotations = @rotations[-n..-1]

      # puts "\t last rotations #{last_rotations}"
      if (last_rotations.first.x == @guard.x && last_rotations.last.y == @guard.y) || (last_rotations.first.y == @guard.y && last_rotations.last.x == @guard.x)
        # puts "\t\t found loop blocker #{next_guard}"
        @loop_blockers.add(next_guard)
      end

      n += 4
    end

    if next_guard.x < 0 || next_guard.x > @max_x || next_guard.y < 0 || next_guard.y > @max_y
      # puts "\toff board #{next_guard}"
      return false
    elsif @blockers.include?(next_guard)
      # puts "\tblocked by #{next_guard}"
      rotate
      walk
    else
      # puts "\twalked to #{next_guard}"

      @visited.add(next_guard)
      @guard = next_guard
    end

    true
  end

  def part1
    build(1)

    has_walked = walk
    while has_walked
      has_walked = walk
    end
    @visited.size
  end

  def part2
    build(2)

    puts "blockers #{@blockers}"
    puts "guard #{@guard}"
    puts

    has_walked = walk
    while has_walked
      has_walked = walk
    end

    puts "visited #{@visited}"
    puts "rotations #{@rotations}"
    puts

    @loop_blockers.delete(@guard_initial)
    @loop_blockers.size
  end
end
