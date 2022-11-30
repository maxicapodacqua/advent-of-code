package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func countIncreases(measurements [][]string) int {
	c := 0
	var prev int
	for i, m := range measurements {
		val, err := strconv.Atoi(m[0])
		if err != nil {
			fmt.Errorf(err.Error())
		}
		if i == 0 {
			prev = val
			continue
		}

		if val > prev {
			//fmt.Printf("val %v higher than %v \n", val, prev)
			c++
		}
		prev = val

		//fmt.Println(val)
	}
	return c
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer f.Close()
	csvR := csv.NewReader(f)
	measurements, err := csvR.ReadAll()

	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Printf("%v", countIncreases(measurements)) // 7
}
