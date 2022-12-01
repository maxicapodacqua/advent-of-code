package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	accum := 0
	for scanner.Scan() {
		row := scanner.Text()

		if row == "" {
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

	fmt.Printf("calories: %v\n", calories)
	sort.Ints(calories)
	topThree := calories[len(calories)-3:]
	fmt.Printf("top 3 calories: %v\n", topThree)

	total := 0
	for _, val := range topThree {
		total += val
	}

	fmt.Printf("Total of top three: %v\n", total)
}
