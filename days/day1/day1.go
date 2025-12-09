package day1

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day1(inputFile string, part int) {
	zeros, zeroTouches := calc(inputFile)
	if part == 0 {
		fmt.Printf("Password: %d\n", zeros)
	} else {
		fmt.Printf("Password: %d\n", zeroTouches)
	}
}

func calc(inputFile string) (int, int) {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	pos := 50
	zeros := 0
	zero_touches := 0
	for ok {
		d, _ := strconv.Atoi(line[1:]) 
		// Complete rounds always count for extra hits to start with
		zero_touches += (d - d % 100)/100
		d = d % 100

		if line[0] == 'L' {
			pos = pos-d

			if pos < 0 {
				if pos+d != 0 { // Counts only if we didn't start at 0
					zero_touches++
				}
				pos = pos + 100
				
			} else if pos == 0 { 
				zero_touches++
			}
		} else {
			pos = pos+d
			if pos > 99 {
				pos = pos - 100
				zero_touches++
			}
		}

		if pos == 0 {
			zeros++
		}

		line, ok = util.Read(ls)
	}
	return zeros, zero_touches
}


