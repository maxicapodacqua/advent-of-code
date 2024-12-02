defmodule Aoc2024.Day01Test do
  use ExUnit.Case
  doctest Aoc2024.Day01
  import Aoc2024.Day01

  test "example test" do
    input = "3   4
4   3
2   5
1   3
3   9
3   3"
    assert part1(input) === 11
  end
end
