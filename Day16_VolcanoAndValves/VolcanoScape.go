package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Valve struct {
	Name   string
	Rate   int
	LeadTo []string
}
type Node struct {
	minute           int
	Score            int
	ShouldOpenValve  bool
	currentNodeValve string
}

func (n *Node) toString() string {
	return fmt.Sprintf("Node %s Minute %d Score %d ShouldOpen %s", n.currentNodeValve, n.minute, n.minute, n.ShouldOpenValve)
}

func (v *Valve) toString() string {
	return fmt.Sprintf("Valve %s Value %d LeadsTo %s", v.Name, v.Rate, strings.Join(v.LeadTo, ","))
}

func str2Int(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func ParseInput(filePath string) map[string]Valve {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	valves := map[string]Valve{}
	replacer := strings.NewReplacer(";", "", ",", "", "=", " ")
	for fileScanner.Scan() {
		splited := strings.Split(replacer.Replace(fileScanner.Text()), " ")
		valves[splited[1]] = Valve{
			splited[1],
			str2Int(splited[5]),
			splited[10:],
		}
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
	return valves
}

func Max(nodes []Node) int {
	maxGeo := nodes[0].Score
	for i := 1; i < len(nodes); i++ {
		if maxGeo < nodes[i].Score {
			maxGeo = nodes[i].Score
		}
	}
	return maxGeo
}

func (node *Node) NextMinute() ([]Node, int, int) {
	minute := node.minute - 1
	newNodes := make([]Node, 0, 5)
	max := node.Score
	if node.ShouldOpenValve {
		newNodes = append(newNodes, Node{minute: minute, ShouldOpenValve: false, currentNodeValve: node.currentNodeValve, Score: node.Score}) // only open, the score comes a minute after
	} else {
		va := valves[node.currentNodeValve]
		max += va.Rate * minute
		for _, v := range va.LeadTo {
			newNodes = append(newNodes, Node{minute: minute, ShouldOpenValve: valves[v].Rate != 0, currentNodeValve: v, Score: max})
		}
	}
	return newNodes, minute, max
}

var valves map[string]Valve

func main() {

	MAX_MINUTE := 30
	startTime := time.Now()
	valves = ParseInput("input.txt")
	maxRateValve := 0
	for _, v := range valves {
		if maxRateValve < v.Rate {
			maxRateValve = v.Rate
		}
	}

	startValve := "AA"
	queue := make([]Node, 0)
	firstNode := Node{MAX_MINUTE, 0, valves[startValve].Rate != 0, startValve}

	queue = append(queue, firstNode)
	highestScore := 0
	for {
		//fifo queue with slice
		if len(queue) == 0 {
			break
		}
		nodes, min, maxOfBranch := queue[0].NextMinute()

		if highestScore < maxOfBranch {
			highestScore = maxOfBranch
		}
		//nodesStr := []string{}
		// for _, n := range nodes {
		// 	nodesStr = append(nodesStr, n.toString())
		// }
		//fmt.Printf("Current Max: %d, Parent: %s \n\t Child: %s. \n", highestScore, queue[0].toString(), strings.Join(nodesStr, "\n"))

		if min == 0 {
			queue = queue[1:]
			continue
		} else {
			//place for optimization, to filter out the nodes that are not promising.
			//if every second steps opens the max rate valve still can not reach max score, it will be filtered.
			// no lack, out of memeory all the time.
			//queue is not a good idea. i believe garbage collection is not so smart.
			//queue = append(queue[1:], nodes...)
			for i, _ := range nodes {
				potentialMax := nodes[i].Score
				if nodes[i].ShouldOpenValve {
					potentialMax += valves[nodes[i].currentNodeValve].Rate * (min - 1)
					for m := min - 2; m >= 0; m -= 2 {
						potentialMax += m * maxRateValve
					}
				} else {
					for m := min - 1; m >= 0; m -= 2 {
						potentialMax += m * maxRateValve
					}
				}

				if highestScore < potentialMax {
					queue = append(queue, nodes[i])
				}
			}
		}
	}

	// for _, v := range valves {
	// 	fmt.Println(v.toString())
	// }
	diff := time.Now().Sub(startTime)
	minutes := diff.String()

	fmt.Printf("The optimization path score is %d, calculation Time: %s. \n", highestScore, minutes)

}
