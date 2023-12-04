package main

import (
	"bufio"
	"fmt"
	"os"
)

var MatchToScore1 = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}

var MatchToScore2 = map[string]int{
	"A X": 3, //rock - sci
	"A Y": 4, //rock - rock
	"A Z": 8, //rock - paper
	"B X": 1, //paper - rock
	"B Y": 5, //paper - paper
	"B Z": 9, //paper - sci
	"C X": 2, // sci - paper
	"C Y": 6, // sci - sci
	"C Z": 7, // sci - rock
}

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalScore1 := 0
	totalScore2 := 0
	var scores1 []int
	var scores2 []int
	for fileScanner.Scan() {
		totalScore1 += MatchToScore1[fileScanner.Text()]
		totalScore2 += MatchToScore2[fileScanner.Text()]
		scores1 = append(scores1, MatchToScore1[fileScanner.Text()])
		scores2 = append(scores2, MatchToScore1[fileScanner.Text()])
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	fmt.Println(totalScore1)
	fmt.Println(totalScore2)
}
