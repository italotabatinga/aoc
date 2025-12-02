require_relative "problem_base"
require_relative "collections"

class Problem12 < ProblemBase
  def problem = 12

  def build(part)
    @garden = []
    @visited= []
    input = (part == 1) ? input1 : input2
    input.each_line.each_with_index do |line, i|
      vec = line[...-1].chars
      @visited.push([false] * vec.size)
      @garden.push(vec)
    end
  end

  def visited?(i, j) = @visited[i][j]

  def visit(p) = @visited[p.i][p.j] = true

  def offgarden?(p) = p.i < 0 || p.j < 0 || p.i >= @garden.size || p.j >= @garden.first.size

  def get(p) = @garden[p.i][p.j]

  def calculate_region(i, j)
    to_visit = Set.new([Vec2.new(i, j)])

    area = 0
    perimeter = 0
    sides = Set.new([])
    while !to_visit.empty?
      v = to_visit.first
      to_visit.delete(v)
      plant = get(v)
      puts "Visiting #{v} -> #{plant}"
      visit(v)
      area += 1

      [Vec2.new(1, 0), Vec2.new(-1, 0), Vec2.new(0, 1), Vec2.new(0, -1)].each do |dp|
        next_v = v + dp

        if offgarden?(next_v)
          perimeter += 1
          sides.add(Vec2.new(v, next_v))
        else
          next_plant = get(next_v)
          if next_plant == plant
            to_visit.add(next_v) if !visited?(next_v.i, next_v.j)
          else
            perimeter += 1
            sides.add(Vec2.new(v, next_v))
          end
        end
      end
    end

    puts "Plant #{plant}"
    puts "  - #{sides}"

    count_sides = 1
    side = sides.first
    while !sides.empty
      sides.delete(side)

      a = side.a
      b = side.b

      horizontal_walk = [Vec2.new(1, 0), Vec2.new(-1, 0)]
      vertical_walk = [Vec2.new(0, 1), Vec2.new(0, -1)]
      if a.i == b.i
        horizontal_walk.each do |d|
          tentative = Vec2.new(a + d, b + d)
          if sides.include?(tenatative)
            side
          end
        end
      else
      end
    end

    [area, perimeter, count_sides]
  end

  def part1
    build(1)

    puts "garden #{@garden}"
    puts "visited #{@visited}"

    total_price = 0
    @garden.each_with_index do |row, i|
      row.each_with_index do |plant, j|
        next if visited?(i, j)

        area, perimeter, _ = calculate_region(i, j)
        price = area * perimeter
        puts "#{plant} region"
        puts "  - Area:      #{area}"
        puts "  - Perimiter: #{perimeter}"
        puts "  - Price:     #{price}"
        puts
        total_price += price
      end
    end
    total_price
  end

  def part2
    build(2)

    75.times.each { blink }
    stones_count
  end
end
