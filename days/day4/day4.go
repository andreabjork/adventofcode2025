package day4

import (
	"adventofcode/m/v2/util"
	"fmt"
)

func Day4(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Accessible: %d\n", checkAccessible(inputFile, true))
	} else {
		fmt.Printf("Accessible: %d\n", checkAccessible(inputFile, false))
	}
}

type ToiletPaper struct {
	i int 
	j int
}

func checkAccessible(inputFile string, singlePassThrough bool) int {
	ls := util.LineScanner(inputFile)
	i := 0
	line, ok := util.Read(ls)

	rolls := []*ToiletPaper{}
	grid := [][]rune{}
	for ok {
		grid = append(grid, []rune{})

		for j, r := range line {
			grid[i] = append(grid[i], r)
			if r == '@' {
				rolls = append(rolls, &ToiletPaper{i, j})
			}
		}
		i++
		line, ok = util.Read(ls)
	}

	remaining := []*ToiletPaper{}
	// Continue until we've reached equilibrium
	for true {
		N := len(remaining)
		remaining = []*ToiletPaper{}
	  remove_rolls := []*ToiletPaper{}
		for _, roll := range rolls {
			if accessible(grid, roll.i, roll.j) {
				remove_rolls = append(remove_rolls, roll)
			} else {
				remaining = append(remaining, roll)
			}
		}

		// Remove the accessible rolls
		for _, rem := range remove_rolls {
			 grid[rem.i][rem.j] = 'x' 
		}

		if singlePassThrough || len(remaining) == N {
			break
		}
	}
		
	return len(rolls)-len(remaining)
}

func accessible(grid [][]rune, x, y int) bool {
	xMin := util.Max(x-1, 0) 
  xMax := util.Min(x+1, len(grid)-1) 
	yMin := util.Max(y-1, 0)  
	yMax := util.Min(y+1, len(grid[0])-1) 
	
	tps := 0
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if i == x && j == y {
				continue
			}
			if grid[i][j] == '@' {
				tps++
			}
		}
	}

	return tps < 4
}