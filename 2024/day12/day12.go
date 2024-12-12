package day12

import (
	"adventofcode/helper"
	"strconv"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var visited = map[[2]int]bool{}
	var board = makeBoard(input)

	for i := range board {
		for j := range board[i] {
			if !visited[[2]int{i, j}] {
				area, perim := getPerimeter(board, visited, i, j)
				res += area * perim
			}
		}
	}

	return strconv.Itoa(res), nil
}

func getPerimeter(board [][]byte, visited map[[2]int]bool, i, j int) (area, perim int) {
	area = 1
	visited[[2]int{i, j}] = true
	for ni, nj := range helper.IterateAdjacentC(i, j) {
		if !helper.InsideC(board, ni, nj) {
			perim++
			continue
		}
		if board[i][j] != board[ni][nj] {
			perim++
			continue
		}
		if !visited[[2]int{ni, nj}] {
			_a, _p := getPerimeter(board, visited, ni, nj)
			area += _a
			perim += _p
		}
	}
	return area, perim
}

func makeBoard(input *helper.InputReader) [][]byte {
	var res [][]byte
	for line := range input.IterateLines {
		res = append(res, []byte(line))
	}
	return res
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var board = makeBoard(input)
	var plots = make([][]int, len(board))
	var plotMap = make([]int, 0)
	for i := range board {
		plots[i] = make([]int, len(board[i]))
	}

	var pos int = 1
	for i := range board {
		for j := range board[i] {
			if plots[i][j] == 0 {
				markPlots(board, plots, pos, i, j)
				plotMap = append(plotMap, pos)
				pos++
			}
		}
	}

	// calc each plot independently for easier edge cases
	for _, plt := range plotMap {
		var area, perim int
		// calc up fences
		for i := range plots {
			var bulk = false
			for j := range plots[i] {
				if plots[i][j] != plt {
					bulk = false
					continue
				}
				area++
				// border up
				if !helper.InsideC(plots, i-1, j) || plots[i-1][j] != plt {
					if !bulk {
						perim++
					}
					bulk = true
				} else {
					bulk = false
				}
			}
		}

		// calc down fences
		for i := range plots {
			var bulk = false
			for j := range plots[i] {
				if plots[i][j] != plt {
					bulk = false
					continue
				}
				// border down
				if !helper.InsideC(plots, i+1, j) || plots[i+1][j] != plt {
					if !bulk {
						perim++
					}
					bulk = true
				} else {
					bulk = false
				}
			}
		}

		// calc left fences
		for i := range plots[0] {
			var bulk = false
			for j := range plots {
				if plots[j][i] != plt {
					bulk = false
					continue
				}
				// border left
				if !helper.InsideC(plots, j, i-1) || plots[j][i-1] != plt {
					if !bulk {
						perim++
					}
					bulk = true
				} else {
					bulk = false
				}
			}
		}

		// calc right fences
		for i := range plots[0] {
			var bulk = false
			for j := range plots {
				if plots[j][i] != plt {
					bulk = false
					continue
				}
				// border right
				if !helper.InsideC(plots, j, i+1) || plots[j][i+1] != plt {
					if !bulk {
						perim++
					}
					bulk = true
				} else {
					bulk = false
				}
			}
		}

		res += area * perim
	}

	return strconv.Itoa(res), nil
}

func markPlots(board [][]byte, plots [][]int, plot int, i, j int) {
	plots[i][j] = plot
	for ni, nj := range helper.IterateAdjacentC(i, j) {
		if !helper.InsideC(board, ni, nj) {
			continue
		}
		if board[i][j] == board[ni][nj] && plots[ni][nj] == 0 {
			plots[ni][nj] = plot
			markPlots(board, plots, plot, ni, nj)
		}
	}
}
