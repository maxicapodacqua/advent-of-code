package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/maxicapodacqua/advent-of-code/2024/01/part1"
)

func main() {

	f,err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	res := part1.Part1(scanner)
	fmt.Printf("RESULT!! -> %v\n", res)
}