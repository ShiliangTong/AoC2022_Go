package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Question: how to achieve the LinQ lambda expression to get the max of an attribute of the object in the list.
// golang lib works almost like LinQ.

func Max(nodes []NodeOfMinute) (int, int) {
	maxGeo := nodes[0].geode
	geoBot := nodes[0].robot[3]
	for i := 1; i < len(nodes); i++ {
		if maxGeo < nodes[i].geode {
			maxGeo = nodes[i].geode
			geoBot = nodes[i].robot[3]
		}
	}
	return maxGeo, geoBot
}

func MaxOfIntArr(intAr []int) int {
	max := intAr[0]
	for i := 1; i < len(intAr); i++ {
		if max < intAr[i] {
			max = intAr[i]
		}
	}
	return max
}

func str2Int(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

// ore robot ingredient [4, 0, 0] 4 ores
// clay robot ingredient [2, 0, 0] 2 ores
// obsidian robot ingredient [3, 14, 0], 3 ores and 14 clay
// geode robot ingredient [2, 0, 7], 2 ores and 7 obsidian
type blueprint struct {
	index           int
	robotIngredient [4][3]int
	maxResourceRq   [3]int
}
type NodeOfMinute struct {
	minute, geode int
	resource      [3]int
	robot         [4]int
}

func (node *NodeOfMinute) toString() string {
	return fmt.Sprintf("M_%d_Geo_%d***Ore_%d_Clay_%d_Obi_%d***OreB_%d_ClayB_%d_ObiB_%d_GB_%d***",
		node.minute, node.geode, node.resource[0], node.resource[1], node.resource[2], node.robot[0], node.robot[1], node.robot[2], node.robot[3])
}

// e.g. Blueprint 30: Each ore robot costs 4 ore.
// Each clay robot costs 4 ore.
// Each obsidian robot costs 4 ore and 12 clay.
// Each geode robot costs 3 ore and 8 obsidian.
func Str2Blueprint(str string) blueprint {
	spltStr := strings.Split(str, " ")
	bp := blueprint{}
	bp.index = str2Int(strings.Replace(spltStr[1], ":", "", 1))
	bp.robotIngredient = [4][3]int{}

	bp.robotIngredient[0] = [3]int{str2Int(spltStr[6]), 0, 0}
	bp.robotIngredient[1] = [3]int{str2Int(spltStr[12]), 0, 0}
	bp.robotIngredient[2] = [3]int{str2Int(spltStr[18]), str2Int(spltStr[21]), 0}
	bp.robotIngredient[3] = [3]int{str2Int(spltStr[27]), 0, str2Int(spltStr[30])}

	bp.maxResourceRq = [3]int{
		MaxOfIntArr([]int{bp.robotIngredient[0][0], bp.robotIngredient[1][0], bp.robotIngredient[2][0], bp.robotIngredient[3][0]}),
		MaxOfIntArr([]int{bp.robotIngredient[0][1], bp.robotIngredient[1][1], bp.robotIngredient[2][1], bp.robotIngredient[3][1]}),
		MaxOfIntArr([]int{bp.robotIngredient[0][2], bp.robotIngredient[1][2], bp.robotIngredient[2][2], bp.robotIngredient[3][2]}),
	}

	return bp
}

func (node *NodeOfMinute) NextMinute(bp blueprint) []NodeOfMinute {
	minute := node.minute + 1
	rs_org := [3]int{node.resource[0], node.resource[1], node.resource[2]}
	rb_org := [4]int{node.robot[0], node.robot[1], node.robot[2], node.robot[3]}
	newNodes := make([]NodeOfMinute, 0, 5)

	//Opt 1: no new robot, keep the resource
	//nodes are too much, and it takes a long time to loop thru. At minute 22, nodes are 122,891,409.
	//optimization action 1: if existing ore resource is bigger or equal to the max required, then never append the no new robot node (clearly not max)
	//Result has improvement, but still requires a significant amount of RAM to run.

	if rs_org[0] < bp.maxResourceRq[0] || rs_org[1] < bp.maxResourceRq[1] { //|| rs_org[2] < bp.maxResourceRq[2] {
		newNodes = append(newNodes, NodeOfMinute{
			minute:   minute,
			geode:    node.geode,
			resource: rs_org,
			robot:    rb_org,
		})
	}

	//Opt: loop thru all robots, add node if potential to produce robot
	//Question, complexity might increase if the robot factory can make more robots
	//optimization action 2: due to one night can only make 1 robot, max robot needed should be the max production requirement among all 4 robots,
	//e.g blueprint OreBot costs [4,0,0], ClayBot costs [2,0,0], ObsiBot costs [3,14,0], GeoBot costs [2,0,7]
	//max ore bot needed should be 4, max claybot needed should be 14, max obsi bot needed shall be 7, more than that are not optimal.
	//Result: blueprint example 2 minute 24 has 147328010 node, and after this optimization it decreased to 120144. That is a hug improvement.
	for i, _ := range node.robot {
		if (i == 3 || rb_org[i] < bp.maxResourceRq[i]) && rs_org[0] >= bp.robotIngredient[i][0] && rs_org[1] >= bp.robotIngredient[i][1] && rs_org[2] >= bp.robotIngredient[i][2] {
			newRobots := [4]int{rb_org[0], rb_org[1], rb_org[2], rb_org[3]}
			newRobots[i]++
			newNodes = append(newNodes, NodeOfMinute{
				minute:   minute,
				geode:    node.geode,
				resource: [3]int{rs_org[0] - bp.robotIngredient[i][0], rs_org[1] - bp.robotIngredient[i][1], rs_org[2] - bp.robotIngredient[i][2]},
				robot:    newRobots,
			})
		}
	}
	//after making robots, havest the existing robot resources
	//nodesAsStrings := []string{}
	for i, _ := range newNodes {
		newNodes[i].geode += node.robot[3]
		for j, n_rb := range node.robot[0:3] {
			newNodes[i].resource[j] += n_rb
		}
		//nodesAsStrings = append(nodesAsStrings, newNodes[i].toString())
	}

	//fmt.Printf("Minute %d, ParentNode: %s, ChildNodes are as below: \n %s \n\t", minute, node.toString(), strings.Join(nodesAsStrings, "\n\t"))
	return newNodes
}

func main() {
	startTime := time.Now()
	//input parsing
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var blueprints []blueprint

	for fileScanner.Scan() {
		bp := Str2Blueprint(fileScanner.Text())
		blueprints = append(blueprints, bp)
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}

	sumQualityLevel := 1
	maxMinutes := 32
	for _, bp := range blueprints {
		fmt.Printf("BluePrint loaded, Blueprint %d: OreBot costs [%d,%d,%d], ClayBot costs [%d,%d,%d], ObsiBot costs [%d,%d,%d], GeoBot costs [%d,%d,%d]. \n",
			bp.index,
			bp.robotIngredient[0][0], bp.robotIngredient[0][1], bp.robotIngredient[0][2],
			bp.robotIngredient[1][0], bp.robotIngredient[1][1], bp.robotIngredient[1][2],
			bp.robotIngredient[2][0], bp.robotIngredient[2][1], bp.robotIngredient[2][2],
			bp.robotIngredient[3][0], bp.robotIngredient[3][1], bp.robotIngredient[3][2])
		//RAM consumption is super high, if I keep track of all the history nodes
		//Optimization 3: only need the 2 generations of nodes.
		var singleGenNodes = []NodeOfMinute{{minute: 1, geode: 0, resource: [3]int{1, 0, 0}, robot: [4]int{1, 0, 0, 0}}}
		maxGeode, maxGeoRobit := 0, 0
		for mi := 2; mi <= maxMinutes; mi++ {
			newGenNodes := []NodeOfMinute{}
			//fmt.Printf("Minute %d: Parent Nodes amount: %d \n", mi, len(singleGenNodes))

			//optimization 4: speed is still rather slow, add node comparsion, to not add it to the next generation if its a bad node
			//Especially at the last a couple of minutes, ignore the node that is not possible to get to the current max.
			//If current node can generate a geode robot every coming minute, and can still not achieve the current max, there is no need to expand that branch.
			// max potential will be building an geode robot every night, then it will be 1+2+3...+ (24-mi-1) = (24-mi)*(23-mi)/2
			//Stack memory saving to only use index, not the item of the slice
			for n := range singleGenNodes {
				if maxGeode+maxGeoRobit*(maxMinutes-mi) <= singleGenNodes[n].geode+singleGenNodes[n].robot[3]*(maxMinutes-mi)+(maxMinutes-mi)*(maxMinutes-mi-1)/2 {
					newGenNodes = append(newGenNodes, singleGenNodes[n].NextMinute(bp)...)
				}
			}

			if m, mg := Max(newGenNodes); m > maxGeode {
				maxGeode = m
				maxGeoRobit = mg
			}

			fmt.Printf("Minute %d: Parent Nodes amount: %d, new child nodes amout: %d, max Geode: %d \n", mi, len(singleGenNodes), len(newGenNodes), maxGeode)
			singleGenNodes = newGenNodes
		}
		diff := time.Now().Sub(startTime)
		minutes := diff.String()
		sumQualityLevel *= maxGeode
		fmt.Printf("The max of the blue print is %d, quality level: %d, calculation Time: %s. \n", maxGeode, bp.index*maxGeode, minutes)
	}

	fmt.Printf("The sum of quality level is %d. \n", sumQualityLevel)
}
