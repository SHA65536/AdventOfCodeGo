package day10

import (
	"adventofcode/helper"
	"strconv"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var board = makeBoard(input)
	res = getTrailHeads(board)

	return strconv.Itoa(res), nil
}

func makeBoard(input *helper.InputReader) [][]byte {
	var res [][]byte
	for line := range input.IterateLines {
		res = append(res, []byte(line))
	}
	return res
}

func getTrailHeads(board [][]byte) int {
	var res int
	var visited = make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}

	var traverseHike func(i, j int)
	traverseHike = func(i, j int) {
		visited[i][j] = true
		if board[i][j] == '9' {
			res++
			return
		}

		if inside(board, i, j+1) && !visited[i][j+1] && board[i][j+1] == board[i][j]+1 {
			traverseHike(i, j+1)
		}
		if inside(board, i+1, j) && !visited[i+1][j] && board[i+1][j] == board[i][j]+1 {
			traverseHike(i+1, j)
		}
		if inside(board, i, j-1) && !visited[i][j-1] && board[i][j-1] == board[i][j]+1 {
			traverseHike(i, j-1)
		}
		if inside(board, i-1, j) && !visited[i-1][j] && board[i-1][j] == board[i][j]+1 {
			traverseHike(i-1, j)
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '0' {
				visited = make([][]bool, len(board))
				for i := range visited {
					visited[i] = make([]bool, len(board[i]))
				}
				traverseHike(i, j)
			}
		}
	}
	return res
}

func inside[T any](board [][]T, i, j int) bool {
	return i >= 0 && i < len(board) && j >= 0 && j < len(board[0])
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var board = makeBoard(input)
	res = getTrailHeadsRating(board)

	return strconv.Itoa(res), nil
}

func getTrailHeadsRating(board [][]byte) int {
	var res int
	var visited = make([][]int, len(board))
	for i := range visited {
		visited[i] = make([]int, len(board[i]))
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}

	var traverseHike func(i, j int) int
	traverseHike = func(i, j int) int {
		if board[i][j] == '9' {
			visited[i][j] = 1
			return 1
		}

		var cur int

		if inside(board, i, j+1) && board[i][j+1] == board[i][j]+1 {
			if visited[i][j+1] == -1 {
				cur += traverseHike(i, j+1)
			} else {
				cur += visited[i][j+1]
			}
		}
		if inside(board, i+1, j) && board[i+1][j] == board[i][j]+1 {
			if visited[i+1][j] == -1 {
				cur += traverseHike(i+1, j)
			} else {
				cur += visited[i+1][j]
			}
		}
		if inside(board, i, j-1) && board[i][j-1] == board[i][j]+1 {
			if visited[i][j-1] == -1 {
				cur += traverseHike(i, j-1)
			} else {
				cur += visited[i][j-1]
			}
		}
		if inside(board, i-1, j) && board[i-1][j] == board[i][j]+1 {
			if visited[i-1][j] == -1 {
				cur += traverseHike(i-1, j)
			} else {
				cur += visited[i-1][j]
			}
		}
		return cur
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '0' {
				res += traverseHike(i, j)
			}
		}
	}
	return res
}
