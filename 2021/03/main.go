package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func sum(slice [][]string) []int {

	acc := make([]int, len(slice))

	for i, item := range slice {
		//fmt.Printf("row %v\n", item)
		acc[i] = 0
		for _, bit := range item {
			//fmt.Printf("bit %v\n", bit)
			bitInt, _ := strconv.Atoi(bit)
			acc[i] += bitInt
		}
	}
	//fmt.Printf("acc %v\n", acc)

	return acc
}

func calculateRates(readings []int, bitsCount int) (gamma, epsilon []string) {
	gamma, epsilon = make([]string, len(readings)), make([]string, len(readings))
	//var gamma []int
	for i, r := range readings {
		//fmt.Printf("r: %v bCount: %v, r > bitsCount/2, %v\n", r, bitsCount, r > bitsCount/2)
		if r > bitsCount/2 {
			gamma[i] = "1"
			epsilon[i] = "0"
		} else {
			gamma[i] = "0"
			epsilon[i] = "1"
		}
	}

	return gamma, epsilon
}

func main() {
	f, err := os.Open("./03/input_csv.txt")
	//f, err := os.Open("input_test.txt")
	if err != nil {
		panic(err)
		//fmt.Errorf(err.Error())
	}
	defer f.Close()
	csvR := csv.NewReader(f)
	//csvR.Comma = ''
	rows, err := csvR.ReadAll()
	if err != nil {
		panic(err)
	}
	transposed := transpose(rows)
	sumTransposed := sum(transposed)
	bitsCount := len(rows)
	gamma, epsilon := calculateRates(sumTransposed, bitsCount)

	//strings.Join(gamma[:], "")
	strRateGamma, strRateEpsilon := strings.Join(gamma[:], ""), strings.Join(epsilon[:], "")

	decimalGamma, _ := strconv.ParseInt(strRateGamma, 2, 64)
	decimalEpsilon, _ := strconv.ParseInt(strRateEpsilon, 2, 64)
	//fmt.Printf("rows %v\n", rows)
	fmt.Printf("bit count %v\n", bitsCount)
	//fmt.Printf("transpose %v\n", transposed)
	//fmt.Printf("sum transposed %v\n", sumTransposed)
	fmt.Printf("calculate rates %v %v\n", gamma, epsilon)
	fmt.Printf("string bytes  gamma: %v epsilon: %v\n", strRateGamma, strRateEpsilon)
	fmt.Printf("decimals  gamma: %v epsilon: %v\n", decimalGamma, decimalEpsilon)
	fmt.Printf("gamma * epsilon %v \n", decimalGamma*decimalEpsilon)
	//fmt.Printf("string bytes %v\n", strBytes)

	//sum(transposed)
	// if sum is higher than half len, most common is 1

}
