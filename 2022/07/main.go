package main

import (
	"bufio"
	"encoding/json"
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

func PrintJSON(obj interface{}) {
	bytes, err := json.MarshalIndent(obj, "  ", "  ")
	panic(err)
	fmt.Println(string(bytes))
}

func strToInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}

func main() {

	f, err := os.Open("./2022/07/input_test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	root := Dir{Name: "/"}

	pointer := &root
	//root.Parent = pointer

	for scanner.Scan() {
		cmd := scanner.Text()
		fmt.Printf(cmd)
		// is a command, not an output
		if strings.HasPrefix(cmd, "$") {
			//parentPointer = pointer
			if strings.HasPrefix(cmd, "$ cd ") {
				moveToDir := strings.TrimPrefix(cmd, "$ cd ")
				//fmt.Printf(" # moveToDir: %v parent %v", moveToDir, pointer.Parent.Name)
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
	fmt.Printf("%+v", root)
	//PrintJSON(root)

	//root := Dir{
	//	Name: "/",
	//	Files: []File{
	//		{
	//			Name: "b.txt",
	//			Size: 14848514,
	//		},
	//		{
	//			Name: "c.dat",
	//			Size: 8504156,
	//		},
	//	},
	//	Dirs: []Dir{
	//		{
	//			Name:  "d",
	//			Files: nil,
	//			Dirs:  nil,
	//		},
	//	},
	//}
	//
	//PrintJSON(root)
	//fmt.Printf("%+v", PrintJSON)

}
