package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	countContain := 0
	countOverlap := 0
	countline := 0
	for fileScanner.Scan() {
		sections := strings.Split(fileScanner.Text(), ",")
		group := strings.Split(sections[0], "-")
		group = append(group, strings.Split(sections[1], "-")...)
		var gp_int []int
		for _, v := range group {
			v_int, _ := strconv.Atoi(v)
			gp_int = append(gp_int, v_int)
		}
		if (gp_int[0] <= gp_int[2] && gp_int[1] >= gp_int[3]) || (gp_int[0] >= gp_int[2] && gp_int[1] <= gp_int[3]) {
			countContain++
		}
		if (gp_int[0] >= gp_int[2] && gp_int[0] <= gp_int[3]) || (gp_int[1] >= gp_int[2] && gp_int[1] <= gp_int[3]) ||
			(gp_int[2] >= gp_int[0] && gp_int[2] <= gp_int[1]) || (gp_int[3] >= gp_int[0] && gp_int[3] <= gp_int[1]) {
			countOverlap++
		}

		//fmt.Printf("Line %d, count %d", countline, count)
		countline++
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	fmt.Println(countOverlap)

	// sort.Ints(eachElf)
	// sum := 0
	// for _, cal := range eachElf[len(eachElf)-3:] {
	// 	fmt.Println(cal)
	// 	sum += cal
	// }
	// fmt.Println(sum)

}
