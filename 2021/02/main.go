package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// directionToMultiplier
// forward X increases the horizontal position by X units.
// down X increases the depth by X units.
// up X decreases the depth by X units.
func directionToMultiplier(direction string) int {
	if direction == "up" {
		return -1
	}
	return 1
}
func calculateFinalPosition(steps [][]string) (horizontal, vertical, aim int) {
	horizontal, vertical, aim = 0, 0, 0

	for _, r := range steps {
		if len(r) == 0 {
			break
		}
		direction := r[0]
		val, _ := strconv.Atoi(r[1])
		multitiplier := directionToMultiplier(direction)
		//fmt.Printf("dir: %v  val: %v multi: %v acc: h: %v v: %v a: %v\n",
		//	direction, val, multitiplier, horizontal, vertical, aim)
		if direction == "forward" {
			horizontal += val * multitiplier
			vertical += val * aim
		}
		if direction == "down" || direction == "up" {
			//vertical += val * multitiplier
			aim += val * multitiplier
		}
		//if direction == "down" {
		//}

	}

	return horizontal, vertical, aim
}

func main() {
	f, err := os.Open("./02/input.txt")
	//f, err := os.Open("input_test.txt")
	if err != nil {
		panic(err)
		//fmt.Errorf(err.Error())
	}
	defer f.Close()
	csvR := csv.NewReader(f)
	csvR.Comma = ' '
	steps, err := csvR.ReadAll()

	if err != nil {
		panic(err)
		//fmt.Errorf(err.Error())
	}
	//fmt.Printf("%v", steps)
	horizontal, vertical, aim := calculateFinalPosition(steps)
	//
	fmt.Printf("h: %v  v: %v aim: %v final result: %v \n ", horizontal, vertical, aim, horizontal*vertical)
}
