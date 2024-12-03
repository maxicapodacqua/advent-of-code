defmodule Aoc2024.Day03 do
  def part1(input) do
    Regex.scan(~r/mul\((?<a>\d{1,3})\,(?<b>\d{1,3})\)/, input, [])
    |> Enum.map(fn matchData ->
      matchTuple = List.to_tuple(matchData)
      a = String.to_integer(elem(matchTuple,1))
      b = String.to_integer(elem(matchTuple,2))
      a * b
     end)
    |> Enum.sum()
  end

  def part2(_input) do
  end

  def runPart1() do
    {:ok, content} = File.read("#{__DIR__}/input_day03.txt")
    part1(content)
  end
end
