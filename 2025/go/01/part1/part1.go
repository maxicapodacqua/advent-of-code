package part1

import (
	"bufio"
	"strconv"
)

const MAX = 100

func nextStep(init, motion int) int {
	return (MAX + init + motion) % MAX
}

func strToInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return res
}

func calcMotion(rotation string, distance int) int {
	if rotation == "L" {
		return -distance
	}
	return distance
}

func Part1(scanner *bufio.Scanner) int {
	acc := 0
	start := 50
	for scanner.Scan() {
		row := scanner.Text()
		rotation,distance := row[0:1], row[1:]
		motion := calcMotion(rotation, strToInt(distance))

		// fmt.Printf("rotation=%v, distance=%v, motion=%v\n", rotation, distance, motion)

		start = nextStep(start, motion)

		// fmt.Printf("Move to= %v\n", start)
		if start == 0 {
			acc++
		}
	}

	return acc
}