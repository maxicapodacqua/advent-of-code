package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

func hasMarker(part []string, charCount int) bool {
	aux := make([]string, len(part))
	copy(aux, part)
	slices.Sort(aux)
	compacted := slices.Compact(aux)

	return len(compacted) == charCount
}

func findMarkerIndex(signal string, charCount int) int {

	signalArr := strings.Split(signal, "")

	for i := 0; i <= len(signalArr)-charCount; i++ {
		tail := i + charCount
		auxArr := signalArr[i:tail]
		isMarker := hasMarker(auxArr, charCount)
		fmt.Printf("part: %v hasMarker: %v tail: %v\n", auxArr, isMarker, tail)
		if isMarker {
			return tail
		}
	}

	panic("NOOO")

	return 0
}

func main() {

	f, err := os.Open("./2022/06/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	signal := scanner.Text()

	res := findMarkerIndex(signal, 4) // part 1
	fmt.Printf("Result part 1: %v", res)

	res = findMarkerIndex(signal, 14) // part 2
	fmt.Printf("Result part 2: %v", res)

}
