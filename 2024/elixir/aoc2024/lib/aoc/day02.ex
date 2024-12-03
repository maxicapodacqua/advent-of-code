defmodule Aoc2024.Day02 do
  def part1(input) do
    # input |> IO.inspect()

    lines =
      input
      |> String.split("\n")

    # IO.puts("Lines #{Enum.join(lines, ",")}")

    lines
    |> Enum.map(fn x ->
      levels = x |> String.split(" ") |> Enum.map(&String.to_integer/1)
      levels
    end)
    |> Enum.filter(&is_safe?/1)
    |> length()
  end

  def part2(input) do
    input
    |> String.split("\n")
    |> Enum.map(fn x ->
      # this can be simplyfied
      x |> String.split(" ") |> Enum.map(&String.to_integer/1)
    end)
    |> Enum.filter(&is_almost_safe?/1)
    |> length()
  end

  @spec is_almost_safe?(list()) :: boolean()
  def is_almost_safe?(list) do
    # generate all permutations of list, including the original list
    permutations = [list | Enum.map(0..(length(list) -1), fn index -> List.delete_at(list, index) end)]
    # if any of the permutations are safe, then it's all safe
    Enum.any?(permutations, &is_safe?/1)
  end

  # If next element is higher than head, then expect whole list to be in asc order
  def is_safe?([h, n | _] = list) when h < n, do: is_safe?(:asc, list)
  # if head is higher than next, then whole list should be in desc order
  def is_safe?([h, n | _] = list) when h > n, do: is_safe?(:desc, list)
  # if head and next are equal, then it's not valid
  def is_safe?([h, n | _]) when h == n, do: false

  # check list continues being asc order and distance between values is "at least one and at most three"
  def is_safe?(:asc, [h, n | rest]) when abs(h - n) in 1..3 and h < n,
    do: is_safe?(:asc, [n | rest])

  # check list continues being desc order and distance between values is "at least one and at most three"
  def is_safe?(:desc, [h, n | rest]) when abs(h - n) in 1..3 and h > n,
    do: is_safe?(:desc, [n | rest])

  # If finished iterating the list (exhausted all elements until empty list) , then the list is valid
  def is_safe?(_, [_last]), do: true
  # otherwise is not valid
  def is_safe?(_, _), do: false

  def runPart1 do
    {:ok, content} = File.read("#{__DIR__}/input_day02.txt")
    part1(content)
  end
  def runPart2 do
    {:ok, content} = File.read("#{__DIR__}/input_day02.txt")
    part2(content)
  end
end
