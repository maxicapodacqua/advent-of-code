package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

func hasMarker(part []string) bool {
	aux := make([]string, len(part))
	copy(aux, part)
	slices.Sort(aux)
	compacted := slices.Compact(aux)

	return len(compacted) == 4
}

func findMarkerIndex(signal string) int {

	signalArr := strings.Split(signal, "")

	for i := 0; i <= len(signalArr)-4; i++ {
		tail := i + 4
		auxArr := signalArr[i:tail]
		isMarker := hasMarker(auxArr)
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

	res := findMarkerIndex(signal)

	fmt.Printf("Result: %v", res)

}
