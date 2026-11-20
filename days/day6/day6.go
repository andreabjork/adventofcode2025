package day6

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day6(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Grand total: %d\n", grandTotal(inputFile, false))
	} else {
		fmt.Printf("Grand total: %d\n", grandTotal(inputFile, true))
	}
}

type OP struct {
	op func(col []int) int
}

func grandTotal(inputFile string, cephalopod bool) int {
	ls := util.LineScanner(inputFile)
	line, _ := util.Read(ls)
	
	// Store all lines but the last one, before processing into columns
	lines := []string{}
	row := 0
	var columns [][]int
	ops := []*OP{}
	for true {
		next, ok := util.Read(ls)
		if !ok {
			break // before adding the last line
		} else {
			lines = append(lines, line)
			line = next
		}
	}

	// Operations
	PLUS := &OP{func(col []int) int { t := 0; for _, c := range col { t += c }; return t }}
	MUL  := &OP{func(col []int) int { t := 1; for _, c := range col { t *= c }; return t }}

	// Keep track of column indices and operations, as well as the width of the last column
	cidx := []int{}
	LAST_WIDTH := 1
	for idx, c := range line { // Reading the last line
		if c == '+' {
			LAST_WIDTH = 1
			ops = append(ops, PLUS)
			cidx = append(cidx, idx)
		} else if c == '*' {
			LAST_WIDTH = 1
			ops = append(ops, MUL)
			cidx = append(cidx, idx)
		}	else {
			LAST_WIDTH++
		}
	}

	if cephalopod {
		columns = cephalopodCols(lines, cidx, LAST_WIDTH)
	} else {
		columns = basicCols(lines, cidx)
	}

	return compute(columns, ops, row-1)
}

func cephalopodCols(lines []string, cidx []int, lastWidth int) [][]int {
	columns := [][]int{}
	for j := 0; j < len(cidx); j++ {
		column := []int{}
		width := 0
		if j == len(cidx)-1 {
			width = lastWidth
		} else {
			width = cidx[j+1]-cidx[j]-1
		}
		for i := 0; i < width; i++ {
			num := ""
			for _, line := range lines {
				num += line[cidx[j]+i:cidx[j]+i+1]
			}
			x, _ := strconv.Atoi(strings.TrimSpace(num))
			column = append(column, x)
		}
		columns = append(columns, column)
	}
	
	return columns
}

func basicCols(lines []string, cidx []int) [][]int {
	columns := [][]int{}
	for j := 0; j < len(cidx); j++ {
		column := []int{}
		for _, line := range lines {
			var num string
			if j == len(cidx)-1 {
				num = line[cidx[j]:]
			} else {
				num = line[cidx[j]:cidx[j+1]-1]
			}
			x, _ := strconv.Atoi(strings.TrimSpace(num))
			column = append(column, x)
		}
		columns = append(columns, column)
	}
	
	return columns
}

func compute(columns [][]int, ops []*OP, MAX_ROWS int) int {
	total := 0
	for j := 0; j < len(columns); j++ {
		total += ops[j].op(columns[j])
	}

	return total
}
