package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Configuration []*list.List

func getInitialConfigurationTest() Configuration {
	/*
		    [D]
		[N] [C]
		[Z] [M] [P]
		 1   2   3
	*/

	var config Configuration

	aux := list.New()
	aux.PushFront("Z")
	aux.PushFront("N")
	config = append(config, aux)

	aux = list.New()
	aux.PushFront("M")
	aux.PushFront("C")
	aux.PushFront("D")
	config = append(config, aux)

	aux = list.New()
	aux.PushFront("P")
	config = append(config, aux)

	return config
}
func getInitialConfiguration() Configuration {
	/*
		    [P]                 [C] [C]
		    [W]         [B]     [G] [V] [V]
		    [V]         [T] [Z] [J] [T] [S]
		    [D] [L]     [Q] [F] [Z] [W] [R]
		    [C] [N] [R] [H] [L] [Q] [F] [G]
		[F] [M] [Z] [H] [G] [W] [L] [R] [H]
		[R] [H] [M] [C] [P] [C] [V] [N] [W]
		[W] [T] [P] [J] [C] [G] [W] [P] [J]
		 1   2   3   4   5   6   7   8   9
	*/

	var config Configuration

	aux := list.New()
	aux.PushFront("W")
	aux.PushFront("R")
	aux.PushFront("F")
	config = append(config, aux)

	aux = list.New()
	aux.PushFront("T")
	aux.PushFront("H")
	aux.PushFront("M")
	aux.PushFront("C")
	aux.PushFront("D")
	aux.PushFront("V")
	aux.PushFront("W")
	aux.PushFront("P")
	config = append(config, aux)

	aux = list.New()
	aux.PushFront("P")
	aux.PushFront("M")
	aux.PushFront("Z")
	aux.PushFront("N")
	aux.PushFront("L")
	config = append(config, aux)

	aux = list.New()
	aux.PushFront("J")
	aux.PushFront("C")
	aux.PushFront("H")
	aux.PushFront("R")
	config = append(config, aux)

	/*
		    [P]                 [C] [C]
		    [W]         [B]     [G] [V] [V]
		    [V]         [T] [Z] [J] [T] [S]
		    [D] [L]     [Q] [F] [Z] [W] [R]
		    [C] [N] [R] [H] [L] [Q] [F] [G]
		[F] [M] [Z] [H] [G] [W] [L] [R] [H]
		[R] [H] [M] [C] [P] [C] [V] [N] [W]
		[W] [T] [P] [J] [C] [G] [W] [P] [J]
		 1   2   3   4   5   6   7   8   9
	*/

	aux = list.New()
	aux.PushFront("C")
	aux.PushFront("P")
	aux.PushFront("G")
	aux.PushFront("H")
	aux.PushFront("Q")
	aux.PushFront("T")
	aux.PushFront("B")
	config = append(config, aux)

	aux = list.New()
	aux.PushFront("G")
	aux.PushFront("C")
	aux.PushFront("W")
	aux.PushFront("L")
	aux.PushFront("F")
	aux.PushFront("Z")
	config = append(config, aux)

	/*
		    [P]                 [C] [C]
		    [W]         [B]     [G] [V] [V]
		    [V]         [T] [Z] [J] [T] [S]
		    [D] [L]     [Q] [F] [Z] [W] [R]
		    [C] [N] [R] [H] [L] [Q] [F] [G]
		[F] [M] [Z] [H] [G] [W] [L] [R] [H]
		[R] [H] [M] [C] [P] [C] [V] [N] [W]
		[W] [T] [P] [J] [C] [G] [W] [P] [J]
		 1   2   3   4   5   6   7   8   9
	*/
	aux = list.New()
	aux.PushFront("W")
	aux.PushFront("V")
	aux.PushFront("L")
	aux.PushFront("Q")
	aux.PushFront("Z")
	aux.PushFront("J")
	aux.PushFront("G")
	aux.PushFront("C")
	config = append(config, aux)

	aux = list.New()
	aux.PushFront("P")
	aux.PushFront("N")
	aux.PushFront("R")
	aux.PushFront("F")
	aux.PushFront("W")
	aux.PushFront("T")
	aux.PushFront("V")
	aux.PushFront("C")
	config = append(config, aux)

	/*
		    [P]                 [C] [C]
		    [W]         [B]     [G] [V] [V]
		    [V]         [T] [Z] [J] [T] [S]
		    [D] [L]     [Q] [F] [Z] [W] [R]
		    [C] [N] [R] [H] [L] [Q] [F] [G]
		[F] [M] [Z] [H] [G] [W] [L] [R] [H]
		[R] [H] [M] [C] [P] [C] [V] [N] [W]
		[W] [T] [P] [J] [C] [G] [W] [P] [J]
		 1   2   3   4   5   6   7   8   9
	*/
	aux = list.New()
	aux.PushFront("J")
	aux.PushFront("W")
	aux.PushFront("H")
	aux.PushFront("G")
	aux.PushFront("R")
	aux.PushFront("S")
	aux.PushFront("V")
	config = append(config, aux)

	return config
}

func (c Configuration) print() {
	for i := 0; i < len(c); i++ {
		fmt.Printf("Printing config %v\n", i)
		c.printColumn(i)
	}
}

func (c Configuration) printColumn(i int) {
	queue := c[i]
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}
}

func (c Configuration) processStep(qty, from, to int) Configuration {
	fmt.Printf("processStep(qty %v, from %v, to %v int)\n", qty, from, to)
	columnToExtract := c[from-1]
	columnToAdd := c[to-1]

	aux := list.New()

	for i := 0; i < qty; i++ {
		e := columnToExtract.Front()
		aux.PushBack(e.Value)
		columnToExtract.Remove(e)
	}

	columnToAdd.PushFrontList(aux)

	return c
}

func (c Configuration) printFinalWord() {
	fmt.Println("Final Word")
	for i := 0; i < len(c); i++ {
		row := c[i]
		e := row.Front()
		fmt.Printf("%v", e.Value)
	}
	fmt.Println()
}

func strToInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}

func main() {

	//configuration := getInitialConfigurationTest()
	configuration := getInitialConfiguration()
	configuration.print()

	f, err := os.Open("./2022/05/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	regex := regexp.MustCompile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)
	for scanner.Scan() {
		stepStr := scanner.Text()
		fmt.Printf("Step: %v\n", stepStr)
		findStr := regex.FindStringSubmatch(stepStr)
		qty, from, to := strToInt(findStr[1]), strToInt(findStr[2]), strToInt(findStr[3])
		fmt.Printf("qty %v, from %v , to %v\n", qty, from, to)

		configuration = configuration.processStep(qty, from, to)
	}

	configuration.print()
	configuration.printFinalWord()
}
