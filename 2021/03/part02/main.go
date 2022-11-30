package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCommonBitForPos(bitPos int, rows [][]string, defaultVal string) string {
	cOnes := 0
	for _, r := range rows {
		if r[bitPos] == "1" {
			cOnes++
		}
	}
	//fmt.Printf("rows/2: %v acc: %v\n", len(rows)/2, cOnes)

	cZeros := len(rows) - cOnes
	if cOnes > cZeros {
		return "1"
	}
	if cOnes < cZeros {
		return "0"
	}
	return defaultVal
}
func filterRows(rows [][]string, bit string, bitPos int) [][]string {
	var newRow [][]string
	for _, r := range rows {
		if r[bitPos] == bit {
			newRow = append(newRow, r)
		}
	}
	return newRow
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
	//fmt.Printf("rows %v\n", rows)

	copyRows := rows
	//consumedRows := rows
	// oxygen
	for i := 0; i < len(rows); i++ {
		//fmt.Printf("row %v \n ", i)
		for bitPos := 0; bitPos < len(rows[i]); bitPos++ {
			//fmt.Printf("bitPos %v bit %v \n", bitPos, bit)
			// get common bit for bitPos
			commonB := getCommonBitForPos(bitPos, rows, "1") //for oxygen default is 1
			// find all rows with commonB in bitPos
			rows = filterRows(rows, commonB, bitPos)
			//fmt.Printf("common bit: %v for pos: %v filtered row: %v len rows: %v \n", commonB, bitPos, rows, len(rows))
			if len(rows) == 1 {
				decimalOxygen, _ := strconv.ParseInt(strings.Join(rows[0][:], ""), 2, 64)
				fmt.Printf("Result decimalOxygen: %v %v", rows[0], decimalOxygen)
			}

		}
	}
	rows = copyRows
	//fmt.Printf("\n\n%v\n", rows)
	// CO
	for i := 0; i < len(rows); i++ {
		fmt.Printf("row %v \n ", i)
		for bitPos := 0; bitPos < len(rows[i]); bitPos++ {
			//fmt.Printf("bitPos %v bit %v \n", bitPos, bit)
			// get common bit for bitPos
			commonB := getCommonBitForPos(bitPos, rows, "1") //for CO default is 1
			// invert
			if commonB == "1" {
				commonB = "0"
			} else {
				commonB = "1"
			}
			// find all rows with commonB in bitPos
			rows = filterRows(rows, commonB, bitPos)
			//fmt.Printf("common bit: %v for pos: %v filtered row: %v len rows: %v \n", commonB, bitPos, rows, len(rows))
			if len(rows) == 1 {
				decimalCO, _ := strconv.ParseInt(strings.Join(rows[0][:], ""), 2, 64)
				fmt.Printf("Result decimalCO: %v %v", rows[0], decimalCO)
			}

		}
	}

	//sum(transposed)
	// if sum is higher than half len, most common is 1

}
