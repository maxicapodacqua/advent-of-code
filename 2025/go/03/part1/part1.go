package part1

import (
	"bufio"
	"slices"
	"strconv"
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

func getLargestJoltage(batteries []int) int {
	// Find max digit in batteries[0:len(batteries)-1]
	// excluding end as we need two digits

	maxVal, index := getMaxAndIndex(batteries[0 : len(batteries)-1])

	// Then cut the slice at this max digit and find the next max digit
	cutBatteries := batteries[index+1 : ]
	nextMaxVal, _ := getMaxAndIndex(cutBatteries)

	return strToInt(intToStr(maxVal) + intToStr(nextMaxVal))
}

func Part1(scanner *bufio.Scanner) int {
	acc := 0
	for scanner.Scan() {
		row := scanner.Text()
		var batteries []int
		for _, batteryStr := range row {
			batteries = append(batteries, strToInt(string(batteryStr)))
		}
		acc += getLargestJoltage(batteries)
	}

	return acc

}
