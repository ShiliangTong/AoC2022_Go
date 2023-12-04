package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}
type rope struct {
	head, tail coordinate
}

func (c *coordinate) AsStr() string {
	return fmt.Sprintf("%d-%d", c.x, c.y)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (r *rope) TailPos_AfterMove() coordinate {
	temp := coordinate{r.tail.x, r.tail.y}
	if Abs(r.head.x-r.tail.x) <= 1 && Abs(r.head.y-r.tail.y) <= 1 {
		return temp
	} else {
		if r.head.x > r.tail.x {
			temp.x++
		} else if r.head.x < r.tail.x {
			temp.x--
		}

		if r.head.y > r.tail.y {
			temp.y++
		} else if r.head.y < r.tail.y {
			temp.y--
		}
	}
	//fmt.Printf("Tail Moved from %s to %s", r.tail.AsStr(), temp.AsStr())
	return temp
}

func RecordMove(dirc, steps string, start rope, record map[string]bool) rope {
	stps, _ := strconv.Atoi(steps)
	temp := rope{coordinate{start.head.x, start.head.y}, coordinate{start.tail.x, start.tail.y}}
	if dirc == "L" {
		for i := 0; i < stps; i++ {
			temp.head.x--
			tail := temp.TailPos_AfterMove()
			temp.tail = tail
			record[tail.AsStr()] = true
		}
	} else if dirc == "R" {
		for i := 0; i < stps; i++ {
			temp.head.x++
			tail := temp.TailPos_AfterMove()
			temp.tail = tail
			record[tail.AsStr()] = true
		}
	} else if dirc == "U" {
		for i := 0; i < stps; i++ {
			temp.head.y++
			tail := temp.TailPos_AfterMove()
			temp.tail = tail
			record[tail.AsStr()] = true
		}
	} else if dirc == "D" {
		for i := 0; i < stps; i++ {
			temp.head.y--
			tail := temp.TailPos_AfterMove()
			temp.tail = tail
			record[tail.AsStr()] = true
		}
	}
	// fmt.Printf("Command: %s %s, start pos: head(%s),tail(%s), new pos: head(%s),tail(%s).\n",
	// 	dirc, steps, start.head.AsStr(), start.tail.AsStr(), temp.head.AsStr(), temp.tail.AsStr())
	return temp
}

type rope_shit struct {
	knots [10]coordinate
}

func (rs *rope_shit) DeepCopyRope() rope_shit {
	var newRS rope_shit
	for i := 0; i < 10; i++ {
		newRS.knots[i] = coordinate{rs.knots[i].x, rs.knots[i].y}
	}
	return newRS
}

func RecordMove_2(dirc, steps string, start rope_shit, record map[string]bool) rope_shit {
	stps, _ := strconv.Atoi(steps)
	temp := start.DeepCopyRope()
	if dirc == "L" {
		for i := 0; i < stps; i++ {
			temp.knots[0].x--
			for i := 0; i < 9; i++ {
				tempRope := rope{head: temp.knots[i], tail: temp.knots[i+1]}
				tail := tempRope.TailPos_AfterMove()
				temp.knots[i+1] = tail
			}
			record[temp.knots[9].AsStr()] = true
		}

	} else if dirc == "R" {
		for i := 0; i < stps; i++ {
			temp.knots[0].x++
			for i := 0; i < 9; i++ {
				tempRope := rope{head: temp.knots[i], tail: temp.knots[i+1]}
				tail := tempRope.TailPos_AfterMove()
				temp.knots[i+1] = tail
			}
			record[temp.knots[9].AsStr()] = true
		}
	} else if dirc == "U" {
		for i := 0; i < stps; i++ {
			temp.knots[0].y++
			for i := 0; i < 9; i++ {
				tempRope := rope{head: temp.knots[i], tail: temp.knots[i+1]}
				tail := tempRope.TailPos_AfterMove()
				temp.knots[i+1] = tail
			}
			record[temp.knots[9].AsStr()] = true
		}
	} else if dirc == "D" {
		for i := 0; i < stps; i++ {
			temp.knots[0].y--
			for i := 0; i < 9; i++ {
				tempRope := rope{head: temp.knots[i], tail: temp.knots[i+1]}
				tail := tempRope.TailPos_AfterMove()
				temp.knots[i+1] = tail
			}
			record[temp.knots[9].AsStr()] = true
		}
	}
	fmt.Printf("Command: %s %s, start pos: head(%s),tail(%s), new pos: head(%s),tail(%s).\n",
		dirc, steps, start.knots[0].AsStr(), start.knots[9].AsStr(), temp.knots[0].AsStr(), temp.knots[9].AsStr())
	return temp
}

func main() {
	filepath := "input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("file %s can not be open, err: %s \n", filepath, err)
	}
	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)

	tailPoses_1 := map[string]bool{"0-0": true}
	tailPoses_2 := map[string]bool{"0-0": true}

	Rope2Start := rope_shit{knots: [10]coordinate{}}

	currentPos := rope{head: coordinate{0, 0}, tail: coordinate{0, 0}}
	curKnots := Rope2Start
	for filescanner.Scan() {
		line := filescanner.Text()
		move := strings.Split(line, " ")
		currentPos = RecordMove(move[0], move[1], currentPos, tailPoses_1)
		curKnots = RecordMove_2(move[0], move[1], curKnots, tailPoses_2)
	}
	fmt.Printf("Tail Positions: %d", len(tailPoses_1))

	fmt.Printf("Tail Positions 2: %d", len(tailPoses_2))

}
