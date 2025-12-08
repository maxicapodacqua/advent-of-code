package part1

import (
	"bufio"
	"fmt"
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

func Part1(scanner *bufio.Scanner) int {
	acc := 0
	var manifoldFull [][]string
	scanner.Scan()
	line := scanner.Text()
	prev := strings.Split(line, "")
	manifoldFull = append(manifoldFull, prev)
	for scanner.Scan() {
		line := scanner.Text()
		current := strings.Split(line, "")
		fmt.Printf("prev    =%v\nmanifold=%v\n\n", prev, current)

		for i, elem := range current {
			if elem == "." && (prev[i] == "|" || prev[i] == "S") {
				current[i] = "|"
			} else if elem == "^" && prev[i]=="|" {
				// replace left and right elements
				current[i-1] = "|"
				current[i+1] = "|"
				// A split happened, must be counted
				acc++
			} else if elem == "|" {
				// Nothing to do
			}
		}


		manifoldFull = append(manifoldFull, current)
		prev = current
	}
	// for _, r := range manifoldFull {
	// 	fmt.Printf("%s\n", r)
	// }

	fmt.Printf("manifoldFull=%v\n", manifoldFull)
	return acc
}
