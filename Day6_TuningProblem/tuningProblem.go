package main

import (
	"bufio"
	"fmt"
	"os"
)

func CheckDistinct(str string) bool {
	m := make(map[rune]bool)
	for _, b := range str {
		m[b] = true
	}
	return len(str) == len(m)
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

	index := 0
	//sum_group := 0
	for i := 14; i < len(fileLines[0]); i++ {
		if CheckDistinct(fileLines[0][i-14 : i]) {
			index = i
			break
		}
	}

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	fmt.Println(index)

	//fmt.Println(sum_group)
}
