require_relative "problem_base"
require_relative "collections"

class Problem8 < ProblemBase
  def problem = 8

  def build(part)
    @map = []
    @antennas = {}
    @antinodes = Set.new
    @i_size = 0
    @j_size = 0
    input = (part == 1) ? input1 : input2
    input.each_line.each_with_index do |line, i|
      line[...-1].chars.each_with_index do |char, j|
        @j_size = j if j > @j_size
        next if char == "."
        @antennas[char] ||= []
        @antennas[char].push(Vec2.new(i, j))
      end
      @i_size = i if i > @i_size
    end

    @i_size += 1
    @j_size += 1
  end

  def part1
    build(1)

    @antennas.each do |_, antennas|
      antennas[...-1].each_with_index do |ant_a, i|
        antennas[i + 1..].each_with_index do |ant_b, j|
          diff = ant_b - ant_a
          potential_antinodes = [ant_b + diff, ant_a - diff]
          potential_antinodes.each do |potential_antinode|
            if potential_antinode.x >= 0 && potential_antinode.x < @i_size && potential_antinode.y >= 0 && potential_antinode.y < @j_size
              @antinodes.add(potential_antinode)
            end
          end
        end
      end
    end

    @antinodes.size
  end

  def part2
    build(2)

    @antennas.each do |_, antennas|
      antennas[...-1].each_with_index do |ant_a, i|
        antennas[i + 1..].each_with_index do |ant_b, j|
          diff = ant_b - ant_a
          init_pos = ant_a - diff
          @antinodes.add(ant_a)
          @antinodes.add(ant_b)
          while init_pos.x >= 0 && init_pos.x < @i_size && init_pos.y >= 0 && init_pos.y < @j_size
            @antinodes.add(init_pos)
            init_pos -= diff
          end
          init_pos = ant_b + diff
          while init_pos.x >= 0 && init_pos.x < @i_size && init_pos.y >= 0 && init_pos.y < @j_size
            @antinodes.add(init_pos)
            init_pos += diff
          end
        end
      end
    end

    @antinodes.size
  end
end
