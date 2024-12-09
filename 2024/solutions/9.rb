require_relative "problem_base"

class File
  attr_reader :index, :size
  attr_accessor :empty

  def initialize(index, size, empty = 0)
    @index = index
    @size = size
    @empty = empty
  end

  def allocate_at_right(file)
    file.empty = @empty - file.size
    @empty = 0
  end

  def to_s = "File(#{@index}, #{size}, #{@empty})"
  alias :inspect :to_s
end

class Problem9 < ProblemBase
  def problem = 9

  def build(part)
    @hd = []
    @files = []
    @last_file_pos = 0
    input = (part == 1) ? input1 : input2
    is_reading_file = true
    files_read = 0
    input.strip.chars.each do |digit_str|
      if is_reading_file
        digit_str.to_i.times { @hd.push(files_read) }
        @last_file_pos = @hd.size - 1
        @files.push(File.new(files_read, digit_str.to_i))

        files_read += 1
      else
        digit_str.to_i.times { @hd.push(nil) }
        @files.last.empty = digit_str.to_i
      end
      is_reading_file = !is_reading_file
    end
  end

  def part1
    build(1)

    i = 0
    j = @last_file_pos
    while i < j
      while @hd[i] != nil
        i += 1
      end

      tmp = @hd[j]
      @hd[j] = @hd[i]
      @hd[i] = tmp
      i += 1
      j -= 1

      while @hd[j] == nil
        j -= 1
      end
    end

    result = 0
    @hd.each_with_index do |block, index|
      break if block.nil?
      result += index * block
    end
    result
  end

  def part2
    build(2)

    i = 0
    while @files[i].empty == 0
      i += 1
    end

    j = @files.size - 1
    while i < j
      k = @files[i...j].find_index do |file|
        file.empty >= @files[j].size
      end

      if k != nil
        k += i
        file = @files[j]
        @files[j-1].empty += file.size + file.empty
        @files.delete_at(j)
        @files.insert(k+1, file)
        @files[k].allocate_at_right(file)
      else
        j -= 1
      end

      while @files[i].empty == 0
        i += 1
      end
    end

    result = 0
    index = 0
    @files.each do |file|
      file.size.times do
        result += file.index * index
        index += 1
      end

      index += file.empty
    end
    result
  end
end
