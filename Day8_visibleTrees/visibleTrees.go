package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filepath := "input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("file %s can not be open, err: %s \n", filepath, err)
	}
	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)
	row := 0
	treeMatx := [99][99]int{}
	for filescanner.Scan() {
		line := filescanner.Text()
		for col := 0; col < len(line); col++ {
			treeMatx[row][col], _ = strconv.Atoi(string(line[col]))
		}
		row++
	}
	// Visible trees
	treeMatx_vis := [99][99]byte{}
	for r := 1; r < 98; r++ {
		prevMax := treeMatx[r][0]
		for co := 1; co < 98; co++ {
			if treeMatx[r][co] > prevMax {
				treeMatx_vis[r][co] |= 1
				prevMax = treeMatx[r][co]
			}
		}
		prevMax = treeMatx[r][98]
		for co := 97; co > 0; co-- {
			if treeMatx[r][co] > prevMax {
				treeMatx_vis[r][co] |= 1 << 1
				prevMax = treeMatx[r][co]
			}
		}
	}

	for c := 1; c < 98; c++ {
		prevMax := treeMatx[0][c]
		for ro := 1; ro < 98; ro++ {
			if treeMatx[ro][c] > prevMax {
				treeMatx_vis[ro][c] |= 1 << 2
				prevMax = treeMatx[ro][c]
			}
		}
		prevMax = treeMatx[98][c]
		for ro := 97; ro > 0; ro-- {
			if treeMatx[ro][c] > prevMax {
				treeMatx_vis[ro][c] |= 1 << 3
				prevMax = treeMatx[ro][c]
			}
		}
	}

	count := 0
	for r := 1; r < 98; r++ {
		for c := 1; c < 98; c++ {
			if treeMatx_vis[r][c] != 0 {
				count++
			}
		}
	}
	fmt.Println(count + 99*2 + 97*2)

	// Highest Scenic score
	scoreMatx := [99][99]int{}
	maxScore := 0
	coor := [2]int{}
	for r := 1; r < 98; r++ {
		for c := 1; c < 98; c++ {
			score := getScenicScore_ByCoodinate(treeMatx, r, c)
			scoreMatx[r][c] = score
			if score > maxScore {
				maxScore = score
				coor = [2]int{r, c}
			}
		}
	}

	//fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(scoreMatx[48])), ";"), "[]"))

	fmt.Printf("Max Score: %d, at X:%d, Y:%d \n", maxScore, coor[0], coor[1])
}

func getScenicScore_ByCoodinate(matx [99][99]int, r int, c int) int {
	score_x1 := 1
	for i := c - 1; i >= 0; i-- {
		if matx[r][i] >= matx[r][c] || i == 0 {
			score_x1 = c - i
			break
		}
	}
	score_x2 := 1
	for i := c + 1; i < 99; i++ {
		if matx[r][i] >= matx[r][c] || i == 98 {
			score_x2 = i - c
			break
		}
	}

	score_y1 := 1
	for i := r - 1; i >= 0; i-- {
		if matx[i][c] >= matx[r][c] || i == 0 {
			score_y1 = r - i
			break
		}
	}
	score_y2 := 1
	for i := r + 1; i < 99; i++ {
		if matx[i][c] >= matx[r][c] || i == 98 {
			score_y2 = i - r
			break
		}
	}

	return score_x1 * score_x2 * score_y1 * score_y2
}
