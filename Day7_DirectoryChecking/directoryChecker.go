package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	name   string
	size   int
	items  map[string]item
	parent *item
	isDir  bool
}

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	root := item{name: "/", isDir: true, items: map[string]item{}}
	currentDir := root

	countDir := map[string]int{}
	for i, line := range fileLines {
		sr := strings.Split(line, " ")
		if sr[1] == "cd" && sr[2] != ".." {
			countDir[fmt.Sprintf("%s-%d", sr[2], i)] = 0
		}
	}
	fmt.Printf("Amount of dirs : %d \n", len(countDir))

	for i, line := range fileLines {
		sts := strings.Split(line, " ")
		if sts[0] == "$" {
			if sts[1] == "cd" {
				if sts[2] == "/" {
					currentDir = root
				} else if sts[2] == ".." {
					//input data does have null ref parent problem, but here should have some check
					//fmt.Println("before back to parent: " + currentDir.name)
					currentDir = *currentDir.parent
					//fmt.Println("back to parent: " + currentDir.name)
				} else {
					currentDir = currentDir.items[sts[2]]
				}
			} else {
				// $ ls
				continue
			}
		} else {
			p := currentDir
			child := item{name: fmt.Sprintf("%s-%d", sts[1], i), size: 0, parent: &p}
			if sts[0] == "dir" {
				child.isDir = true
				child.items = map[string]item{}
			} else {
				size, _ := strconv.Atoi(sts[0])
				child.size = size
				child.isDir = false
			}
			currentDir.items[sts[1]] = child
		}
		//fmt.Println(currentDir.name)
	}

	allDir := map[string]int{}
	wholeDirsize := CalcDirSize(root, allDir)
	fmt.Println(wholeDirsize)
	//smallDir := make(map[string]item)
	totalSize_below1m := 0
	//SumSize_SmallDir(root, smallDir)

	for _, v := range allDir {
		//count++
		//fmt.Printf("Count: %d, Dir: %s, size: %d \n", count, k, v)
		if v <= 100000 {
			//count_s++
			totalSize_below1m += v
			//fmt.Printf("-----Count: %d, Dir: %s, size: %d \n", count_s, k, v)
		}
	}

	minSize := 30000000 - (70000000 - 48044502)
	bestV := minSize
	bestK := ""
	for k, v := range allDir {
		//count++
		//fmt.Printf("Count: %d, Dir: %s, size: %d \n", count, k, v)
		if v >= minSize && (v-minSize) < bestV {
			//count_s++
			bestV = v - minSize
			bestK = k
			//fmt.Printf("-----Count: %d, Dir: %s, size: %d \n", count_s, k, v)
		}
	}
	fmt.Printf("Best dir to delete: %s, size: %d, min size: %d \n", bestK, allDir[bestK], minSize)

	// for _, it := range smallDir {
	// 	//totalSize_below1m += it.size
	// 	fmt.Printf("Current item: %s, size: %d \n", it.name, it.size)
	// }
	fmt.Println(totalSize_below1m)

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
}

func CalcDirSize(dir item, allDirs map[string]int) int {
	p := dir
	for _, child := range p.items {
		if !child.isDir {
			p.size += child.size
		} else {
			p.size += CalcDirSize(child, allDirs)
		}
	}
	allDirs[p.name] = p.size
	return p.size
}

func SumSize_SmallDir(dir item, smallDirs map[string]item) {
	//sumTemp := sum
	if dir.size <= 100000 {
		smallDirs[dir.name] = dir
	}
	for _, child := range dir.items {
		if child.isDir {
			SumSize_SmallDir(child, smallDirs)
		}
	}
	//fmt.Printf("%s - Current list: %d", dir.name, len(smallDirs))
}
