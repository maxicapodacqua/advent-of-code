package part1

import (
	"bufio"
	"io"
	"regexp"
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

func Part1(reader io.ReadSeeker, scanner *bufio.Scanner) int {
	acc := 0
	re := regexp.MustCompile(`((\*|\+)|[1-9])+`)
	var operations []string
	// Got lazy trying to figure out how to make golang read the last line first, so I made this AMAZING solution..

	// Get last line
	for scanner.Scan() {
		operations = re.FindAllString(scanner.Text(), -1)
	}
	// Reset cursor
	_, err := reader.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	// Creating array of results
	colResults := make([]int, len(operations))
	// And setting defaults so multiplication doesn't stays on 0
	for i, op := range operations {
		if op == "*" {
			colResults[i] = 1
		} else {
			// Golang default, but just to make it readable
			colResults[i] = 0
		}
	}

	// Reset scanner with reader reference
	scanner = bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '*' || line[0] == '+' {
			// last line
			break
		}
		elementStr := re.FindAllString(line, -1)
		for i, elem := range elementStr {
			val := strToInt(elem)
			op := operations[i]
			if op == "*" {
				colResults[i] *= val
			}
			if op == "+" {
				colResults[i] += val
			}
		}
	}

	for _, col := range colResults {
		acc += col
	}

	return acc
}
