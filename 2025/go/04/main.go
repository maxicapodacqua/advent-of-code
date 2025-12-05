package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/maxicapodacqua/advent-of-code/2025/04/part1"
)

func main() {

	f,err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	res1 := part1.Part1(scanner)
	fmt.Printf("Part 1 -> %v\n", res1)

	// // Reset file pointer to read again in new scanner
	// _, err = f.Seek(0, io.SeekStart)
	// if err != nil {
	// 	panic(err)
	// }

	// res2 := part2.Part2(bufio.NewScanner(f))
	// fmt.Printf("Part 2 -> %v\n", res2)
}