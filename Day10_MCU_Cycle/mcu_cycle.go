package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	countCycle := 0
	cycleScores := map[int]int{0: 1}
	X := 1
	for fileScanner.Scan() {
		sa := strings.Split(fileScanner.Text(), " ")
		countCycle++
		cycleScores[countCycle] = X
		if sa[0] == "noop" {
			continue
		} else {
			countCycle++
			cycleScores[countCycle] = X
			v, _ := strconv.Atoi(sa[1])
			X += v
		}
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}

	sum := 0
	for i := 20; i <= 220; i += 40 {
		sum += i * cycleScores[i]
		//fmt.Printf("Cycle: %d, X: %d sum: %d \n", i, cycleScores[i], sum)
	}
	fmt.Println(sum)
	// for k, v := range cycleScores {
	// 	fmt.Printf("Cycle: %d, X: %d \n", k, v)
	// }
	CTU := [6][40]bool{}
	for i := 0; i < 240; i++ {
		//fmt.Printf("C: %d, X: %d \n", i, cycleScores[i])
		CTU[i/40][i%40] = Abs(i%40-cycleScores[i+1]) <= 1
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			if CTU[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}
}
