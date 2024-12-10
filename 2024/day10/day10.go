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
	var visited [][]int

	var traverseHike func(i, j int)
	traverseHike = func(i, j int) {
		if board[i][j] == '9' {
			visited[i][j] = 1
			res++
			return
		}

		for ni, nj := range helper.IterateAdjacentC(i, j) {
			if helper.InsideC(board, ni, nj) && visited[ni][nj] == -1 && board[ni][nj] == board[i][j]+1 {
				traverseHike(ni, nj)
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '0' {
				visited = make([][]int, len(board))
				for ti := range visited {
					visited[ti] = make([]int, len(board[ti]))
					for tj := range visited[ti] {
						visited[ti][tj] = -1
					}
				}
				traverseHike(i, j)
			}
		}
	}
	return res
}

func traverseHike(board [][]byte, visited [][]int, i, j int) int {
	if board[i][j] == '9' {
		visited[i][j] = 1
		return 1
	}

	var cur int

	for ni, nj := range helper.IterateAdjacentC(i, j) {
		if !helper.InsideC(board, ni, nj) || board[ni][nj] != board[i][j]+1 {
			continue
		}
		if visited[ni][nj] == -1 {
			cur += traverseHike(board, visited, ni, nj)
		} else {
			cur += visited[ni][nj]
		}
	}
	return cur
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

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '0' {
				res += traverseHike(board, visited, i, j)
			}
		}
	}
	return res
}
