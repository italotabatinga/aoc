class Vec2
  attr_reader :x, :y

  alias_method :i, :x
  alias_method :j, :y
  alias_method :a, :x
  alias_method :b, :y

  def initialize(x, y)
    @x = x
    @y = y
  end

  def initialize_dup(other) = new(other.x, other.y)

  def to_s = "(#{x}, #{y})"

  alias_method :inspect, :to_s

  def +(other) = Vec2.new(@x + other.x, @y + other.y)

  def -(other) = Vec2.new(@x - other.x, @y - other.y)

  def ==(other) = @x == other.x && @y == other.y

  alias_method :eql?, :==

  def hash = [x, y].hash
end
