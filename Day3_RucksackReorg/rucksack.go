package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

const abcABC = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func char2Score(c string) int {
	return strings.Index(abcABC, c) + 1
}

// charSet is a limited bitset to contain all lowercase latin characters (a-z).
type charSet struct {
	bits uint64
}

// Set switches a bit to 1 for any given character i that is in the range 'a'
// to 'z'. Characters outside this range have undefined behavior.
func (c *charSet) Set(i uint8) {
	c.bits |= 1 << (i - 'a')
}

// Intersects returns whether two charSets intersect (i.e., share one or more
// on bits).
// func (c charSet) Intersects(o charSet) bool {
// 	return c.bits&o.bits != 0
// }

func (c charSet) SharedChar(o charSet) string {

	return string(abcABC[bits.TrailingZeros64(c.bits&o.bits)])
}
func (c charSet) SharedCharScore(o charSet) int {
	fmt.Println(string(abcABC[bits.TrailingZeros64(c.bits&o.bits)]))
	return bits.TrailingZeros64(c.bits & o.bits)
}

// set returns a charSet for all bytes in s. Bytes in s must be in the range of
// 'a' to 'z'. Anything outside that range is regarded as undefined behavior.
func set(s []byte) charSet {
	var c charSet
	for i := 0; i < len(s); i++ {
		c.Set(s[i])
	}
	return c
}

// subst returns whether two strings share any characters. l and r are assumed
// to only contain characters in the range of 'a' to 'z'.
// func subst(l, r []byte) int {
// 	return set(l).SharedCharScore(set(r))
// }

func substr3(s1, s2, s3 string) int {
	m := map[uint8]bool{}
	m2 := map[uint8]bool{}

	for i := 0; i < len(s1); i++ {
		m[s1[i]] = true
	}

	for i := 0; i < len(s2); i++ {
		if m[s2[i]] {
			m2[s2[i]] = true
		}
	}

	for i := 0; i < len(s3); i++ {
		if m2[s3[i]] {
			fmt.Println(string(s3[i]))
			return char2Score(string(s3[i]))
		}
	}
	return 0
}

func substr(s1 string, s2 string) int {
	m := map[uint8]bool{}

	for i := 0; i < len(s1); i++ {
		m[s1[i]] = true
	}

	for i := 0; i < len(s2); i++ {
		if m[s2[i]] {
			//fmt.Println(string(s2[i]))
			return char2Score(string(s2[i]))
		}
	}
	return 0
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

	sum := 0
	sum_group := 0
	for i, line := range fileLines {
		sum += substr(line[:(len(line)/2)], line[(len(line)/2):])
		if (i+1)%3 == 0 {
			fmt.Println("1: " + fileLines[i-2] + ", 2: " + fileLines[i-2] + ", 3: " + line)
			sum_group += substr3(fileLines[i-2], fileLines[i-1], line)
		}
		//sum += subst([]byte(line[:(len(line)/2)]), []byte(line[(len(line)/2):]))
	}

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	fmt.Println(sum)

	fmt.Println(sum_group)
}
