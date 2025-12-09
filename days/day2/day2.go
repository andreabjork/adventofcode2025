package day2

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
	"strconv"
)

func Day2(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Invalid sum: %d\n", sumInvalid(inputFile, 0))
	} else {
		fmt.Printf("Invalid sum: %d\n", sumInvalid(inputFile, 0))
	}
}

func sumInvalid(inputFile string, threshold int) int {
	ls := util.LineScanner(inputFile)
	input, _ := util.Read(ls)
	var ranges []string

	ranges = strings.Split(input, ",")
		
	var parts []string 
	var from, to int
	sum := 0
	for _, r := range ranges {
		parts = strings.Split(r, "-")
		from, _ = strconv.Atoi(parts[0])
		to, _  = strconv.Atoi(parts[1])
		sum += invalids(from, to)
	}

	return sum
} 

func invalids(from int, to int) int {
	fmt.Printf("Processing %d-%d\n", from, to)
	m := 0
	for x := from; x <= to; x++ {
		s := strconv.Itoa(x)	
		
		for splits := 2; splits <= len(s); splits++ {
			ss := ""
			if len(s)%splits == 0 {
				for k := 0; k < splits; k++ {
					ss += s[:len(s)/splits]
				}

				if s == ss {
					fmt.Printf("%s sequence, %d splits, adding %d to sum\n", s[:len(s)/splits], splits, x)
					m += x
					break
				}
			}
		}
	}

	return m
}
