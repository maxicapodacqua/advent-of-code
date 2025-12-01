package part2

import (
	"bufio"
	"fmt"
	"math"
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

func calcPassedZero(init, motion int) int {
	spin := init + motion
	fmt.Printf(" init=%v, motion=%v, Spin=%v\n", init, motion, spin)
	// if spin == 0 {
	// 	return 1
	// }
	div := float64(spin)/float64(MAX)
	floor := int(math.Floor(div))
	abs := max(floor, -floor)
	fmt.Printf(" div=%v, floor=%v, abs=%v\n", div, floor, abs)
	return abs
}

func Part2(start int, scanner *bufio.Scanner) int {
	acc := 0
	for scanner.Scan() {
		row := scanner.Text()
		rotation,distance := row[0:1], row[1:]
		motion := calcMotion(rotation, strToInt(distance))

		fmt.Printf("START: %v \n rotation=%v, distance=%v, motion=%v\n",start, rotation, distance, motion)

		passedZero := calcPassedZero(start,motion)

		fmt.Printf(" Passed Zero=%v\n", passedZero)

		acc += passedZero


		start = nextStep(start, motion)

		fmt.Printf(" Acc= %v, Move to= %v\n", acc, start)
		// if start == 0 {
		// 	acc++
		// }
	}

	return acc
}