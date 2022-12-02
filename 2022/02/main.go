package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var shapeScore = map[string]int{
	"X": 1, // rock
	"Y": 2, // paper
	"Z": 3, //scissors
}

var stdShapes = map[string]string{
	"A": "R", // rock
	"B": "P", // paper
	"C": "S", // scissors
	"X": "R",
	"Y": "P",
	"Z": "S",
}

// rock > scissors

// paper > rock

// scissors > paper
func calculateOutcome(theirs, ours string) int {
	theirShape, ourShape := stdShapes[theirs], stdShapes[ours]
	// draw
	if theirShape == ourShape {
		return 3
	}
	// win
	if theirShape == "R" && ourShape == "P" {
		return 6
	}
	// win
	if theirShape == "P" && ourShape == "S" {
		return 6
	}
	// win
	if theirShape == "S" && ourShape == "R" {
		return 6
	}

	// loss
	return 0

}
func calculateScore(game []string) int {
	theirs, ours := game[0], game[1]
	ourScore := shapeScore[ours]
	outcomeRound := calculateOutcome(theirs, ours)

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
