package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func countIncreases(measurements []int) int {
	c := 0
	var prev int
	for i, val := range measurements {
		//val, err := strconv.Atoi(m[0])
		//if err != nil {
		//	fmt.Errorf(err.Error())
		//}
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

func reduceSum(arr [][]string) int {
	acc := 0
	for _, item := range arr {
		val, _ := strconv.Atoi(item[0])
		acc += val
	}
	return acc
}

func reduceMeasurements(measurements [][]string) []int {
	//tail := 0
	//head := tail + 3
	var res []int
	for tail := 0; tail <= len(measurements); tail++ {

		head := tail + 3
		if head > len(measurements) {
			fmt.Println("Reached end of measurements")
			break
		}
		res = append(res, reduceSum(measurements[tail:head]))

		//fmt.Printf("measurements[%v:%v] %v sum: %v\n", tail, head, measurements[tail:head], reduceSum(measurements[tail:head]))

	}
	return res
}

func main() {

	f, err := os.Open("input.txt")
	//f, err := os.Open("input_test.txt")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer f.Close()
	csvR := csv.NewReader(f)
	measurements, err := csvR.ReadAll()

	if err != nil {
		fmt.Errorf(err.Error())
	}

	//fmt.Printf("%v", countIncreases(measurements))

	reduced := reduceMeasurements(measurements)

	fmt.Printf("reduced: %v\n", reduced)
	fmt.Printf("result: %v\n", countIncreases(reduced))

	//fmt.Printf("%v\n", measurements[:8])
	//fmt.Printf("%v\n", measurements[0:3])
	//fmt.Printf("%v\n", measurements[1:4])
	//fmt.Printf("%v\n", measurements[2:5])
	//fmt.Printf("%v\n", measurements[3:6])
	//fmt.Printf("%v\n", measurements[4:7])
	//a, b := 5, 8
	//fmt.Printf("%v\n", measurements[a:b])
}
