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

func Part2(scanner *bufio.Scanner) int {
	acc := 0
	// get ranges from input and keep them in memory
	var ranges [][]int
	for scanner.Scan() {
		row := scanner.Text()
		// is a range?
		if strings.Contains(row, "-") {
			rangesRaw := strings.Split(row, "-")
			left, right := strToInt(rangesRaw[0]), strToInt(rangesRaw[1])
			ranges = append(ranges, []int{left, right})
		} else if row == "" {
			// the empty row separator
			break
		}
	}

	// sort them by their left end
	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})
	// After sorting, we can start merging overlapping ranges
	// two ranges a and b overlap if:
	// a[0] <= b[0] <= a[1]
	// If an overlapping is happening, then create a new range
	// whose values will be:
	// a[0], max(a[1], b[1])
	queue := make([][]int, 0)
	// Using a queue approach
	// Adding first element to start checking on ranges
	queue = append(queue, ranges[0])
	for i, b := range ranges {
		// Skip first one as it is in the queue already
		if i == 0 {
			continue
		}
		// This is a queue.pop()
		a := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if IsOverlapping(a, b) {
			newRange := []int{min(a[0], b[0]), max(a[1], b[1])}
			queue = append(queue, newRange)
		} else {
			queue = append(queue, a)
			queue = append(queue, b)
		}
	}

	for _, r := range queue {
		diff := r[1] - r[0] + 1
		acc += diff
	}

	return acc
}

func IsOverlapping(a, b []int) bool {
	return (a[0] <= b[0] && b[0] <= a[1]) 
}
