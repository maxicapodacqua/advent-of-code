package part2

import (
	"bufio"
	"strconv"
)

const MAX = 100

func strToInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return res
}


func Part2(start int, scanner *bufio.Scanner) int {
	acc := 0
	for scanner.Scan() {
		row := scanner.Text()
		rotation,distance := row[0:1], strToInt(row[1:])
		// A more brute force approach
		var step int
		if rotation == "L" {
			step = -1
		} else {
			step =1
		}

		// Literally crank the dial by the distance
		for range distance {
			start = (start + step) % MAX
			// And if we step on a 0, count it
			if start == 0 {
				acc++
			}
		}

		// fmt.Printf("\n%v\nStart: %v", row, start)
		// fmt.Printf(" Acc= %v, Move to= %v\n", acc, start)
	}

	return acc
}