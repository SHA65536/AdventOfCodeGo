package day07

import (
	"adventofcode/helper"
	"strconv"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var grid, _ = input.ReadByteLines()

	for i := range grid {
		for j := range grid[i] {
			// Propagate downward
			if grid[i][j] == '.' && helper.InsideC(grid, i-1, j) && (grid[i-1][j] == 'S' || grid[i-1][j] == '|') {
				grid[i][j] = '|'
			}
			// Split
			if grid[i][j] == '^' && helper.InsideC(grid, i-1, j) && grid[i-1][j] == '|' {
				res++
				grid[i+1][j-1] = '|'
				grid[i+1][j+1] = '|'
			}
		}
	}

	return strconv.Itoa(res), nil
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var rawGrid, _ = input.ReadByteLines()
	var grid = make([][]int, len(rawGrid))
	for i := range rawGrid {
		grid[i] = make([]int, len(rawGrid[i]))
		for j := range rawGrid[i] {
			if rawGrid[i][j] == 'S' {
				grid[i][j] = 1
			}
			if rawGrid[i][j] == '^' {
				grid[i][j] = -1
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] < 1 || !helper.InsideC(grid, i+1, j) {
				continue
			}
			if grid[i+1][j] != -1 {
				// Propagate downward
				grid[i+1][j] += grid[i][j]
			} else {
				// Split
				grid[i+1][j-1] += grid[i][j]
				grid[i+1][j+1] += grid[i][j]
			}
		}
	}

	for i := range len(grid[len(grid)-1]) {
		res += grid[len(grid)-1][i]
	}

	return strconv.Itoa(res), nil
}
