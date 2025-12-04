package day04

import (
	"adventofcode/helper"
	"strconv"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var grid, _ = input.ReadByteLines()
	res, _ = CountAccessible(grid)

	return strconv.Itoa(res), nil
}

func CountAccessible(grid [][]byte) (int, [][2]int) {
	var accessible [][2]int
	var res int
	for i := range grid {
		for j := range grid[i] {
			var total int
			if grid[i][j] != '@' {
				continue
			}
			for ni, nj := range helper.IterateAroundC(i, j) {
				if helper.InsideC(grid, ni, nj) && grid[ni][nj] == '@' {
					total++
				}
			}
			if total < 4 {
				accessible = append(accessible, [2]int{i, j})
				res++
			} else {
			}
		}
	}
	return res, accessible
}

func Star2(input *helper.InputReader) (string, error) {
	var res int
	var prev int = 1

	var grid, _ = input.ReadByteLines()

	for prev > 0 {
		total, accesible := CountAccessible(grid)
		res += total

		for _, p := range accesible {
			grid[p[0]][p[1]] = '.'
		}

		prev = total
	}

	return strconv.Itoa(res), nil
}

func Star2SinglePass(input *helper.InputReader) (string, error) {
	var res int

	var grid, _ = input.ReadByteLines()

	var findAndRemove func(int, int)
	findAndRemove = func(i, j int) {
		var total int
		if !helper.InsideC(grid, i, j) || grid[i][j] == '.' {
			return
		}
		for ni, nj := range helper.IterateAroundC(i, j) {
			if helper.InsideC(grid, ni, nj) && grid[ni][nj] == '@' {
				total++
			}
		}
		if total < 4 {
			res++
			grid[i][j] = '.'
			for ni, nj := range helper.IterateAroundC(i, j) {
				findAndRemove(ni, nj)
			}
		}
	}

	for i := range grid {
		for j := range grid[0] {
			findAndRemove(i, j)
		}
	}

	return strconv.Itoa(res), nil
}
