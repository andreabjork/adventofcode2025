package day5

import (
	"adventofcode/m/v2/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Day5(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Fresh: %d\n", countFresh(inputFile, true))
	} else {
		fmt.Printf("Fresh: %d\n", countFresh(inputFile, false))
	}
}

type Interval struct {
	from, to int
	merged bool
}

func countFresh(inputFile string, withIngredients bool) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	intervals := []*Interval{}
	for ok {
		// Continue parsing ingredients
		if line == "" {
			break
		}
		// Parse ranges
		parts := strings.Split(line, "-")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])

		intervals = append(intervals, &Interval{from, to, false})

		line, ok = util.Read(ls)
	}

	// Extend every interval as much as possible
	for _, iv := range intervals {
		extendInterval(iv, intervals)
	}
	
	if withIngredients {
		return checkIngredients(intervals, ls)
	} else {
		return sumIntervals(intervals)
	}
}

func extendInterval(i *Interval, intervals []*Interval) {
	if i.merged {
		return
	}
	for true {
		from, to := i.from, i.to
		for _, j := range intervals {
			if i == j || j.merged {
				continue
			}

			if i.from >= j.from && i.from <= j.to { // i.from is contained in interval j
				i.from = j.from
				j.merged = true 
			} 

			if i.to <= j.to && i.to >= j.from { // i.to is contained in interval j
				i.to = j.to
				j.merged = true
			}
		}

		if from == i.from && to == i.to {
			// If we made a whole iteration without any change, we're done
			break 
		}
	}
	
}

func checkIngredients(intervals []*Interval, ls *bufio.Scanner) int {
	line, ok := util.Read(ls)
	fresh := 0
	for ok {
		x, _ := strconv.Atoi(line)
		for _, iv := range intervals {
			if !iv.merged && x >= iv.from && x <= iv.to {
				fresh++
				break	
			}
		}

		line, ok = util.Read(ls)
	}
	return fresh
}


func sumIntervals(intervals []*Interval) int {
	sum := 0
	for _, iv := range intervals {
		if !iv.merged {
			fmt.Printf("%d %d\n", iv.from, iv.to)
			sum += iv.to-iv.from+1
		}
	}

	return sum
}