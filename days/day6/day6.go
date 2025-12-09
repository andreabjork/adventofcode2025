package day6

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
	"strconv"
	"regexp"
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
	line, ok := util.Read(ls)
	
	row := 0
	var columns [][]int
	ops := []*OP{}
	var ns []string
	for ok {
		// White space processing is annoying, remove leading and trailing first
		// so we can split properly on other whitespace
		leading := regexp.MustCompile(`^\s+`)
		trailing := regexp.MustCompile(`\s+$`)
		lline := trailing.ReplaceAllString(leading.ReplaceAllString(line, ""), "")

		re := regexp.MustCompile(`\s+`)
		ns = re.Split(lline, -1)
		if columns == nil { 
			columns = make([][]int, len(ns))
		}
		for col, n := range ns {
			x, err := strconv.Atoi(n)
			if err == nil {
				columns[col] = append(columns[col], x)
			}
		}
		row++

		line, ok = util.Read(ls)
	}
	
	// Create ops vector if we're at the last line
	PLUS := &OP{func(col []int) int { t := 0; for _, c := range col { t += c }; return t }}
	MUL  := &OP{func(col []int) int { t := 1; for _, c := range col { t *= c }; return t }}
	if cephalopod {
		columns = mapCols(columns)
	}

	// Create the op vector
	for _, n := range ns {
		if strings.Compare(n, "+") == 0 {
			ops = append(ops, PLUS)
		} else {
			ops = append(ops, MUL)
		}
	}

	return compute(columns, ops, row-1)
}

func mapCols(columns[][]int) [][]int {
	mappedColumn := []int{}
	for _, col := range columns {
		for i := 0; i < len(col); i++ {
			
		}
	}
}

func compute(columns [][]int, ops []*OP, MAX_ROWS int) int {
	total := 0
	for j := 0; j < len(columns); j++ {
		total += ops[j].op(columns[j])
	}

	return total
}
