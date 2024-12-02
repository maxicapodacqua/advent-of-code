defmodule Aoc2024.Day02Test do
  use ExUnit.Case
  doctest Aoc2024.Day02
  import Aoc2024.Day02

  test "part1" do
    input = "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"
    assert part1(input) == 2
  end

end
