package part2

import (
	"bufio"
	"slices"
	"strconv"
	"strings"
)

func strToInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return res
}

func intToStr(input int) string {
	return strconv.Itoa(input)
}

func getMaxAndIndex(slice []int) (int, int) {
	maxVal := slices.Max(slice)
	index := slices.Index(slice, maxVal)
	return maxVal, index
}

func GetLargestJoltage(batteries []int) int {
	var res []string

	// Now this..probably a good use case to practice recursion
	// this solution is based on the solution for part 1
	// The idea is to always check the head of the array
	// creating a sub array and finding the max,
	// the length of the sub array depends on how many digits are needed
	// Start with 12, and keep decreasing until all digits are found
	input := batteries
	digitsNeeded := 12
	for i := digitsNeeded; i > 0; i-- {
		l := len(input)
		regionEnd := l - (i - 1)
		portionToCheck := input[0:regionEnd]
		maxVal, indexWhereMaxFound := getMaxAndIndex(portionToCheck)
		// Add max to result list
		res = append(res,intToStr(maxVal))

		// Cut input and keep running
		input = input[indexWhereMaxFound+1:]
	}

	return strToInt(strings.Join(res,""))
}

func Part2(scanner *bufio.Scanner) int {
	acc := 0
	for scanner.Scan() {
		row := scanner.Text()
		var batteries []int
		for _, batteryStr := range row {
			batteries = append(batteries, strToInt(string(batteryStr)))
		}
		acc += GetLargestJoltage(batteries)
	}

	return acc

}
