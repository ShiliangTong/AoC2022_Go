package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	max := 0
	temp := 0
	var eachElf []int
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {

			if temp > max {
				max = temp
			}
			eachElf = append(eachElf, temp)
			temp = 0
			continue
		}
		val, _ := strconv.Atoi(fileScanner.Text())
		temp += val
		//fmt.Println(temp)
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	fmt.Println(max)

	sort.Ints(eachElf)
	sum := 0
	for _, cal := range eachElf[len(eachElf)-3:] {
		fmt.Println(cal)
		sum += cal
	}
	fmt.Println(sum)

}
