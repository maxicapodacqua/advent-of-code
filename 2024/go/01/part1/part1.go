package part1

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

// Pair up the smallest number in the left list with the smallest number in the right list,
// then the second-smallest left number with the second-smallest right number, and so on.
func Part1(scanner *bufio.Scanner) int {
	var leftList []int
	var rightList []int
	for scanner.Scan() {
		row := scanner.Text()
		splitRes := strings.Split(row, "   ")
		l, r := splitRes[0], splitRes[1]
		// fmt.Printf("row=%v splitRes=%v l=%v r=%v\n", row, splitRes, l, r)
		lVal,_ := strconv.Atoi(l)
		rVal,_ := strconv.Atoi(r)
		leftList = append(leftList, lVal)
		rightList = append(rightList, rVal)
	}
	// fmt.Printf("LeftList=%v \nRightList=%v\n", leftList, rightList)
	sort.Ints(leftList)
	sort.Ints(rightList)

	// fmt.Printf("SORTED: \nLeftList=%v \nRightList=%v\n", leftList, rightList)
	acc := 0
	// Now calculate the value
	for i:=0; i<len(leftList); i++ {
		l,r := leftList[i], rightList[i]
		diff := l - r
		// Nasty way to get abs in golang
		absDiff := max(diff, -diff)
		acc += absDiff
	}

	return acc
}