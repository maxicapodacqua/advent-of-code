package part1

import (
	"bufio"
	"strconv"
	"strings"
)

func IsInvalidProjectID(projectID string) bool {
	strLen := len(projectID)
	isEvenLen := (strLen % 2) == 0
	if !isEvenLen {
		return false
	}
	left, right := projectID[0:(strLen/2)], projectID[(strLen/2):]
	return left == right
}

func strToInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return res
}


func Part1(scanner *bufio.Scanner) int {
	acc := 0
	scanner.Scan()
	input := scanner.Text()
	projectRanges := strings.SplitSeq(input, ",")
	for projectRange := range projectRanges {
		splits := strings.Split(projectRange, "-")
		low, high := strToInt(splits[0]), strToInt(splits[1])
		// THE FASTEST ALGORRTHMNMM EVERRR
		for projectId := low; projectId <= high; projectId++ {
			if IsInvalidProjectID(strconv.Itoa(projectId)) {
				acc += projectId
			}
		}
	}

	return acc

}
