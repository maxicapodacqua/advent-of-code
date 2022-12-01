package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("./2022/01/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var calories []int
	max := 0

	accum := 0
	for scanner.Scan() {
		row := scanner.Text()

		if row == "" {

			if accum > max {
				max = accum
			}

			calories = append(calories, accum)
			accum = 0
			continue
		}

		cal, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}
		accum += cal

		fmt.Printf("Line: %v\n", cal)
	}

	// the last line is excluded by this library
	calories = append(calories, accum)
	if accum > max {
		max = accum
	}

	fmt.Printf("calories: %v\n", calories)
	fmt.Printf("max calories: %v\n", max)

	//max := 0
	//for _, c := range calories {
	//	if c > max {
	//		max = c
	//	}
	//}

	//fmt.Printf("inventory: %v", inventory)
}
