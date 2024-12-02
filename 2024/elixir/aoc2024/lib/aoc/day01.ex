# create basic structure for day 1
defmodule Aoc2024.Day01 do
  def part1(input) do
    lines =
      input
      |> String.split("\n")
      |> Enum.map(&String.split(&1, "   "))

    left =
      Enum.map(lines, fn [a, _] -> a end)
      |> Enum.map(&String.to_integer(&1))
      |> Enum.sort()

    right =
      Enum.map(lines, fn [_, a] -> a end)
      |> Enum.map(&String.to_integer(&1))
      |> Enum.sort()

    distanceTotal =
      Enum.zip(left, right)
      |> Enum.map(fn {l, r} -> abs(l - r) end)
      |> Enum.reduce(fn x, acc -> x + acc end)

    # cols = lines
    # |> String.split("   ")
    # left |> IO.inspect()
    # right |> IO.inspect()
    distanceTotal
  end

  def part2(input) do
    lines =
      input
      |> String.split("\n")
      |> Enum.map(&String.split(&1, "   "))

    left =
      Enum.map(lines, fn [a, _] -> a end)
      |> Enum.map(&String.to_integer(&1))

    right =
      Enum.map(lines, fn [_, a] -> a end)
      |> Enum.map(&String.to_integer(&1))

    # Enum.frequencies does this exact thing :()
    mapFreq =
      Enum.reduce(right, %{}, fn element, acc ->
        case Map.get(acc, element) do
          nil ->
            Map.put(acc, element, 1)

          curr_value ->
            new_val = if is_nil(curr_value), do: 0, else: curr_value + 1
            Map.put(acc, element, new_val)
        end
      end)

    Enum.map(left, fn element ->
      freq =
        case Map.get(mapFreq, element) do
          nil -> 0
          x -> x
        end

      element * freq
    end)
    |> Enum.reduce(fn x, acc -> x + acc end)
  end

  def runPart1 do
    {:ok, content} = File.read("#{__DIR__}/input_day01.txt")
    # content |> IO.inspect()
    part1(content)
  end

  def runPart2 do
    {:ok, content} = File.read("#{__DIR__}/input_day01.txt")
    # content |> IO.inspect()
    part2(content)
  end
end
