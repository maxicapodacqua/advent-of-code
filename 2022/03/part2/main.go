package main

import (
	"bufio"
	"fmt"
	"os"
)

func getMatchChar(firstPart, secondPart, thirdPart string) string {
	for i := 0; i < len(firstPart); i++ {
		for j := 0; j < len(secondPart); j++ {
			for k := 0; k < len(thirdPart); k++ {
				if firstPart[i] == secondPart[j] && secondPart[j] == thirdPart[k] {
					match := string(secondPart[j])
					fmt.Printf("Match found: %v %v\n", match, secondPart[j])
					return match
				}
			}
		}
	}
	panic("Expected match")
	return ""
}

var charToScore = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func main() {
	f, err := os.Open("./2022/03/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	totalScore := 0

	for {
		scanner.Scan()
		row1 := scanner.Text()
		scanner.Scan()
		row2 := scanner.Text()
		res := scanner.Scan()
		row3 := scanner.Text()

		if res == false {
			break
		}

		fmt.Printf("row1: %v row2: %v row3: %v \n", row1, row2, row3)

		match := getMatchChar(row1, row2, row3)

		score := charToScore[match]
		fmt.Printf("score: %v\n", score)
		totalScore += score
	}
	fmt.Printf("--->Total score: %v", totalScore)
}
