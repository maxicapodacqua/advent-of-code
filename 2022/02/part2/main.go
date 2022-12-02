package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var shapeScore = map[string]int{
	"A": 1, // rock
	"B": 2, // paper
	"C": 3, //scissors
}

var outcomeMap = map[string]int{
	"X": 0, // lose
	"Y": 3, // draw
	"Z": 6, //win
}

var winningMap = map[string]string{
	"A": "B", // rock < paper
	"B": "C", // paper < scissors
	"C": "A", // scissors < rock
}

// invert of losing map
var losingMap = map[string]string{
	"B": "A",
	"C": "B",
	"A": "C",
}

// rock > scissors

// paper > rock

// scissors > paper

func guessShape(theirs, ours string) int {
	// draw returns same shape as theirs
	if ours == "Y" {
		return shapeScore[theirs]
	}

	// loss
	if ours == "X" {
		return shapeScore[losingMap[theirs]]
	}

	return shapeScore[winningMap[theirs]]
}

func calculateScore(game []string) int {
	theirs, ours := game[0], game[1]
	outcomeRound := outcomeMap[ours]
	ourScore := guessShape(theirs, ours)

	fmt.Printf("calculate score, game: %v ourScore: %v, outcome: %v\n", game, ourScore, outcomeRound)

	return ourScore + outcomeRound
}
func main() {
	f, err := os.Open("./2022/02/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	csvR := csv.NewReader(f)
	csvR.Comma = ' '
	games, err := csvR.ReadAll()

	if err != nil {
		panic(err)
	}

	fmt.Printf("games: %v \n", games)
	var scores []int

	acc := 0
	for _, game := range games {
		score := calculateScore(game)
		scores = append(scores, score)
		acc += score
	}

	fmt.Printf("scores: %v\n", scores)
	fmt.Printf("acc total score: %v\n", acc)

}
