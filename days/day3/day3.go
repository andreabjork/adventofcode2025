package day3

import (
	"fmt"
	"adventofcode/m/v2/util"
	"strconv"
)

func Day3(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Voltage sum: %d\n", voltageSum(inputFile, 2))
	} else {
		fmt.Printf("Voltage sum: %d\n", voltageSum(inputFile, 12))
	}
}

func voltageSum(inputFile string, size int) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	batteries := [][]int{}
	for ok {
		v := []int{}
		for _, c := range line { 
			x, _ := strconv.Atoi(string(c))
			v = append(v, x)
		}		

		batteries = append(batteries, v)
		line, ok = util.Read(ls)
	}
	
	return findGreatest(&batteries, size)
 }

 func findGreatest(batteries *[][]int, size int) int {
	sum := 0
	for _, battery := range *batteries {
		var x int
		idx := -1
		n := size
		XY := ""
		for n > 0 {
			idx, x = max(battery, idx+1, len(battery)-n)
			n--
			XY += strconv.Itoa(x)
		}

		sum += val(XY)
	}

	return sum
 }

 // Returns the (index, value) of the maximum value
 // within battery[i:j] (j included)
 func max(battery []int, i, j int) (int, int) {
	idx := i
	m := battery[i]
	for k := i+1; k <= j; k++ {
		if battery[k] > m {
			idx = k
			m = battery[k]
		}
	}

	return idx, m
 }

 func val(XY string) int {
	 v, _ := strconv.Atoi(XY)
	 return v
 }