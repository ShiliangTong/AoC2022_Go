package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func move1(orgStack map[int]string, orgSt, newSt int) {
	orgstr := orgStack[orgSt]
	orgDestStr := orgStack[newSt]
	toMove := orgstr[len(orgstr)-1:]
	orgStack[orgSt] = orgstr[:len(orgstr)-1]
	orgStack[newSt] = orgDestStr + toMove
	//fmt.Printf("Org stack %d: %s, after shift %s \n", orgSt, orgstr, orgStack[orgSt])
	//fmt.Printf("Org dest stack %d: %s, after shift %s \n", newSt, orgDestStr, orgStack[newSt])
}

func move(orgStack map[int]string, quantity, orgSt, newSt int) {
	for i := 0; i < quantity; i++ {
		move1(orgStack, orgSt, newSt)
	}
}

func Stoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
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

	//read the start matrix
	// contents are in 1, 5, 9, 13...
	stacks := make(map[int]string)

	for i := 7; i >= 0; i-- {
		for j := 0; j < 9; j++ {
			if string(fileLines[i][1+4*j]) != "" {
				stacks[j+1] += strings.Trim(string(fileLines[i][1+4*j]), " ")
			}
		}
	}

	for id, v := range stacks {
		fmt.Printf("%d : %s \n", id, v)
	}

	for _, line := range fileLines[10:] {
		splited := strings.Split(line, " ")

		move(stacks, Stoi(splited[1]), Stoi(splited[3]), Stoi(splited[5]))
	}

	Head := ""

	for i := 1; i <= 9; i++ {
		Head += string(stacks[i][len(stacks[i])-1:])
	}

	// for i, line := range fileLines {
	// 	sum += substr(line[:(len(line)/2)], line[(len(line)/2):])
	// 	if (i+1)%3 == 0 {
	// 		fmt.Println("1: " + fileLines[i-2] + ", 2: " + fileLines[i-2] + ", 3: " + line)
	// 		sum_group += substr3(fileLines[i-2], fileLines[i-1], line)
	// 	}
	// 	//sum += subst([]byte(line[:(len(line)/2)]), []byte(line[(len(line)/2):]))
	// }

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	fmt.Println(Head)

	// fmt.Println(sum_group)
}
