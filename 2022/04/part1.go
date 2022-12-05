package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	From int
	To   int
}

func strToInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}
func NewAssignment(row string) Assignment {
	split := strings.Split(row, "-")
	f, t := strToInt(split[0]), strToInt(split[1])

	return Assignment{
		From: f,
		To:   t,
	}
}

func (a Assignment) contains(b Assignment) bool {
	return (a.From >= b.From && a.To <= b.To) || (b.From >= a.From && b.To <= a.To)
}

func (a Assignment) overlap(b Assignment) bool {
	for i := a.From; i <= a.To; i++ {
		// using the contains of a single element to don't repeat code
		res := b.contains(Assignment{
			From: i,
			To:   i,
		})
		if res {
			return true
		}
	}
	return false
}

func main() {

	f, err := os.Open("./2022/04/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	csvR := csv.NewReader(f)
	rows, err := csvR.ReadAll()
	if err != nil {
		panic(err)
	}

	containAcc, overlapAcc := 0, 0

	for _, row := range rows {
		partA, partB := NewAssignment(row[0]), NewAssignment(row[1])
		contain := partA.contains(partB)
		overlap := partA.overlap(partB)
		fmt.Printf("row: %v, partA: %#v, partB: %#v, contain? %v, overlap? %v\n", row, partA, partB, contain, overlap)

		if contain {
			containAcc++
		}
		if overlap {
			overlapAcc++
		}
	}
	fmt.Printf("-----> Score: containAcc: %v, overlapAcc: %v", containAcc, overlapAcc)
}
