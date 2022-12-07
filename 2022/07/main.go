package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	Name   string
	Files  []File
	Dirs   []Dir
	Parent *Dir
}
type File struct {
	Name string
	Size int
}

func (d Dir) Size() int {
	acc := 0
	for _, dir := range d.Dirs {
		// here is the recursion
		acc += dir.Size()
	}
	for _, file := range d.Files {
		acc += file.Size
	}
	return acc
}

func (d Dir) Part1() int {

	acc := 0
	size := d.Size()
	if size <= 100000 {
		acc += size
	}

	for _, dir := range d.Dirs {
		acc += dir.Part1()
	}

	return acc
}

func strToInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}

func main() {

	f, err := os.Open("./2022/07/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	root := Dir{Name: "/"}

	pointer := &root

	for scanner.Scan() {
		cmd := scanner.Text()
		fmt.Printf(cmd)
		// is a command, not an output
		if strings.HasPrefix(cmd, "$") {
			//parentPointer = pointer
			if strings.HasPrefix(cmd, "$ cd ") {
				moveToDir := strings.TrimPrefix(cmd, "$ cd ")
				if moveToDir == ".." {
					pointer = pointer.Parent
				} else {
					for i, dir := range pointer.Dirs {
						if dir.Name == moveToDir {
							pointer = &pointer.Dirs[i]
							break
						}
					}
				}
			} else {
				fmt.Printf(" # LS is ignored\n")
				continue
			}
		} else { // is output
			line := strings.Split(cmd, " ")
			size, filename := line[0], line[1]
			fmt.Printf(" # size: %v, filename: %v", size, filename)
			if size == "dir" { // is a directory
				pointer.Dirs = append(pointer.Dirs, Dir{
					Name:   filename,
					Parent: pointer,
				})
			} else { // is a file
				pointer.Files = append(pointer.Files, File{
					Name: filename,
					Size: strToInt(size),
				})
			}
		}

		fmt.Println()

	}

	fmt.Println("Printing struct")
	fmt.Printf("%+v\n", root)
	fmt.Printf("Full size (size of root): %+v\n", root.Size())
	fmt.Printf("Part 1: %+v\n", root.Part1())
}
