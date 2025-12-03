package part2

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/maxicapodacqua/advent-of-code/2025/02/part1"
)


func canConsumeProjectID(part, projectID string) bool {
	if part == "" {
		return false
	}
	// Empty projectID means that it fully consumed the string
	if projectID == "" {
		return true
	}
	// Cut the projectId by the part we are testing, if the match is found, keep doing it until we reach empty projectId
	after, found := strings.CutPrefix(projectID, part)
	if found {
		// YOOO if this recursion works I'm getting a drink
		return canConsumeProjectID(part, after)
	}
	return false
}

func IsInvalidProjectID(projectID string) bool {
	// This just to save some computation and make cases where the string has only two chars work
	if part1.IsInvalidProjectID(projectID) {
		return true
	}
	// Try for every part of the projectID if the whole projectID can be split into equal parts
	for i := 1; i < len(projectID)-1; i++ {
		part := projectID[0:i]
		if canConsumeProjectID(part, projectID) {
			return true
		}
	}
	return false
}

func strToInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return res
}

func Part2(scanner *bufio.Scanner) int {
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
