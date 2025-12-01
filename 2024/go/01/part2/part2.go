package part2

import (
	"bufio"
	"strconv"
	"strings"
)

// Calculate a total similarity score by adding up each number in the left list
// after multiplying it by the number of times that number appears in the right list.
func Part2(scanner *bufio.Scanner) int {
	// As I generate the a map of ocurrences as I generate the right list,
	// then iterate on the left list to calculate the similarities

	// Stores occurences of numbers in right list
	rMap := map[int]int{}
	var lList []int
	// var rList []int
	for scanner.Scan() {
		row := scanner.Text()
		splitRes := strings.Split(row, "   ")
		l, r := splitRes[0], splitRes[1]
		// fmt.Printf("row=%v splitRes=%v l=%v r=%v\n", row, splitRes, l, r)
		// I should really use a helper func here because this is going to happen often
		lVal,_ := strconv.Atoi(l)
		rVal,_ := strconv.Atoi(r)

		lList = append(lList, lVal)

		// Alright yo, map is ready
		rMap[rVal]++
	}

	// fo sure
	// fmt.Printf("rMap=%v\n", rMap)

	acc := 0
	// Now let's do some math
	for _,k := range lList {
		multiplier := rMap[k]
		acc += k * multiplier
		
	}

	return acc
}