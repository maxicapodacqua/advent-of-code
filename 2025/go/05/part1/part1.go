package part1

import (
	"bufio"
	"strconv"
	"strings"
)

func strToInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return res
}

func intToStr(input int) string {
	return strconv.Itoa(input)
}

func IsIngredientFresh(ingredientId int, ranges [][]int) bool {
	for _, r := range ranges {
		if ingredientId >= r[0] && ingredientId <= r[1] {
			return true
		}
	}
	return false
}

func Part1(scanner *bufio.Scanner) int {
	acc := 0
	var ranges [][]int
	for scanner.Scan() {
		row := scanner.Text()
		// is a range?
		if strings.Contains(row, "-") {
			rangesRaw := strings.Split(row, "-")
			left, right := strToInt(rangesRaw[0]), strToInt(rangesRaw[1])
			ranges = append(ranges, []int{left, right})
		} else if row == "" {
			// the empty row separator
			continue
		} else {
			// is an ingredient id
			ingredientId := strToInt(row)
			if IsIngredientFresh(ingredientId, ranges) {
				acc++
			}

		}

	}
	return acc
}
