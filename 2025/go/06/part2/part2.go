package part2

import (
	"bufio"
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

func Part2(scanner *bufio.Scanner) int {
	acc := 0
	return acc
}
