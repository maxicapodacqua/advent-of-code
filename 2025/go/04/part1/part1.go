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

func canAccessRollOfPaper(i,j int,matrix [][]string) bool {
	motions := [8][2]int {
		// x,y
		{0, -1}, //north
		{+1, -1}, //ne
		{+1, 0}, //east
		{+1, +1}, //se
		{0, +1}, //south
		{-1, +1}, //sw
		{-1, 0}, //west
		{-1, -1}, //nw
	}
	heigh := len(matrix)
	width := len(matrix[0])

	rollCount := 0

	for _, motion := range motions {
		motionX, motionY := motion[0], motion[1]
		newI := i + motionX
		newJ := j + motionY
		// Verify is not out of bounds
		if newJ >= width || newJ < 0 {
			continue;
		}
		if newI >= heigh || newI < 0 {
			continue;
		}

		if matrix[newI][newJ] == "@" {
			rollCount++
		}

		// If the count reached 4, stop
		if rollCount == 4 {
			return false
		}

	}
	
	return true
}
func Part1(scanner *bufio.Scanner) int {
	acc := 0
	var matrix [][]string
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}
	for i, row := range matrix {
		for j, value := range row {
			if value == "@" && canAccessRollOfPaper(i,j,matrix) {
				acc++
			}
		}
	}
	return acc

}
