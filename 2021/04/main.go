package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Board struct {
	Values     [5][5]string
	Multiplier [5][5]int // reduced array with 0 Values and 1 where there were matches
}

func (b *Board) AppendRow(row [5]string) {
	// finds first empty spot and inserts there
	for i, val := range b.Values {
		if val[0] == "" {
			b.Values[i] = row
			//b.Multiplier[i] = [5]int{0, 0, 0, 0, 0}
			break
		}
	}
}

func (b *Board) CheckNumberDraw(number string) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if strings.Trim(b.Values[i][j], " ") == strings.Trim(number, " ") {
				r := &b.Multiplier[i]
				r[j] = 1
				return true
			}
		}
	}
	return false
}

//func (b Board) Multiplier()  {
//	fmt.Printf("multiplier: %v", b.multiplier)
//}

func main() {

	f, err := os.Open("./04/input_test.txt")
	//f, err := os.Open("input_test.txt")
	if err != nil {
		panic(err)
		//fmt.Errorf(err.Error())
	}
	defer f.Close()
	csvR := csv.NewReader(f)
	csvR.FieldsPerRecord = -1
	game, err := csvR.ReadAll()
	if err != nil {
		panic(err)
	}
	var numbersDraw []string
	//var boards [][][]string
	//var board [][]string
	boardId := 0

	var boardsSlice []Board

	b := Board{}
	for i, g := range game {

		if i == 0 {
			numbersDraw = g
			continue
		}
		if g[0] == "=" {
			boardsSlice = append(boardsSlice, b)
			b = Board{}
			boardId++
			continue
		}
		//boardRowId := (i - 1) % 5
		var aux [5]string
		copy(aux[:], g[0:5])
		b.AppendRow(aux)

		//fmt.Printf("g %v boardId: %v boardRowId: %v i: %v aux: %v b: %v \n", g, boardId, boardRowId, i, aux, b)
	}
	fmt.Printf("numbers draw: %v\n", numbersDraw)
	fmt.Printf("boards: %v\n", boardsSlice)

	for _, number := range numbersDraw {
		fmt.Printf("...Drawing %v...\n", number)
		for _, board := range boardsSlice {
			if board.CheckNumberDraw(number) {
				fmt.Printf("      number %v found in board: %v\n", number, board.Values)
				fmt.Printf("      multiplier: %v\n", board.Multiplier)
			}
		}
	}

}
