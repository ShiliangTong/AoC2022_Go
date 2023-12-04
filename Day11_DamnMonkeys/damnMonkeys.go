package main

import (
	"fmt"
	"strings"
)

type DamnMonkey struct {
	name        int
	items       []int
	operation   func(val int) int
	divbleBy    int
	trueTarget  int
	falseTarget int
	//inspectedItems []int
	inspectionCount int
}

func (mon *DamnMonkey) PassItemTo(val int) (int, int) {
	after_op := mon.operation(val) % 9699690
	if after_op%mon.divbleBy == 0 {
		return after_op, mon.trueTarget
	} else {
		return after_op, mon.falseTarget
	}
}
func (mon *DamnMonkey) ShowBackpack() string {
	// itemsSt := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(mon.items)),
	// 	","), "[]")
	// inspectedItemsSt := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(mon.inspectedItems)),
	// 	","), "[]")
	//countInspected := len(mon.inspectedItems)
	// return fmt.Sprintf("Monkey %d, has items:%s. Inspected %d items. (%s)",
	// 	mon.name, itemsSt, countInspected, inspectedItemsSt)
	return fmt.Sprintf("Monkey %d inspected items %d times. \n", mon.name, mon.inspectionCount)

}

var monks [8]DamnMonkey

func init() {
	monks[0] = DamnMonkey{0, []int{83, 88, 96, 79, 86, 88, 70},
		func(val int) int { return val * 5 }, 11, 2, 3, 0}
	monks[1] = DamnMonkey{1, []int{59, 63, 98, 85, 68, 72},
		func(val int) int { return val * 11 }, 5, 4, 0, 0}
	monks[2] = DamnMonkey{2, []int{90, 79, 97, 52, 90, 94, 71, 70},
		func(val int) int { return val + 2 }, 19, 5, 6, 0}
	monks[3] = DamnMonkey{3, []int{97, 55, 62},
		func(val int) int { return val + 5 }, 13, 2, 6, 0}
	monks[4] = DamnMonkey{4, []int{74, 54, 94, 76},
		func(val int) int { return val * val }, 7, 0, 3, 0}
	monks[5] = DamnMonkey{5, []int{58},
		func(val int) int { return val + 4 }, 17, 7, 1, 0}
	monks[6] = DamnMonkey{6, []int{66, 63},
		func(val int) int { return val + 6 }, 2, 7, 5, 0}
	monks[7] = DamnMonkey{7, []int{56, 56, 90, 96, 68},
		func(val int) int { return val + 7 }, 3, 4, 1, 0}
}

func main() {
	for i := 0; i < 10000; i++ {
		//fmt.Printf("------------Round %d-------------\n", i)
		for j := 0; j < len(monks); j++ {
			for _, v := range monks[j].items {
				passVal, pass2 := monks[j].PassItemTo(v)
				monks[pass2].items = append(monks[pass2].items, passVal)
			}
			monks[j].inspectionCount += len(monks[j].items)
			monks[j].items = []int{}
		}
		if i == 8000 {
			for j := 0; j < len(monks); j++ {
				fmt.Printf("Mon %d items: %s \n", monks[j].name,
					strings.Trim(strings.Join(strings.Fields(fmt.Sprint(monks[j].items)),
						","), "[]"))
			}
		}

		//fmt.Println()
	}
	for j := 0; j < len(monks); j++ {
		fmt.Print(monks[j].ShowBackpack())
	}

}
