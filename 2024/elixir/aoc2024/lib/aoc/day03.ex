defmodule Aoc2024.Day03 do
  def part1(input) do
    Regex.scan(~r/mul\((?<a>\d{1,3})\,(?<b>\d{1,3})\)/, input, [])
    |> Enum.map(fn matchData ->
      matchTuple = List.to_tuple(matchData)
      a = String.to_integer(elem(matchTuple, 1))
      b = String.to_integer(elem(matchTuple, 2))
      a * b
    end)
    |> Enum.sum()
  end

  def part2(input) do
    # Capture all operations, mul, do and don't
    Regex.scan(~r/mul\((\d{1,3})\,(\d{1,3})\)|do\(\)|don\'t\(\)/, input)
    # very interesting way of using reduce
    |> Enum.reduce({:continue, 0}, fn
      # pattern mathing the parameters used in the reduce callback
      ["don't()"], {_, acc} ->
        {:skip, acc}

      ["do()"], {_, acc} ->
        {:continue, acc}

      [_, _, _], {:skip, acc} ->
        {:skip, acc}

      [_, a, b], {:continue, acc} ->
        {:continue, acc + String.to_integer(a) * String.to_integer(b)}
    end)
    |> elem(1)
  end

  def runPart1() do
    {:ok, content} = File.read("#{__DIR__}/input_day03.txt")
    part1(content)
  end
  def runPart2() do
    {:ok, content} = File.read("#{__DIR__}/input_day03.txt")
    part2(content)
  end
end
