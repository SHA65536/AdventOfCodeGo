package day06

import (
	"adventofcode/helper"
	"strconv"
)

var dirs = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var board = makeBoard(input)
	var pos, dir = findGuard(board)

	for inside(board, pos) {
		if board[pos[0]][pos[1]] != 'X' {
			board[pos[0]][pos[1]] = 'X'
			res++
		}

		npos := [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		for inside(board, npos) && board[npos[0]][npos[1]] == '#' {
			dir = (dir + 1) % 4
			npos = [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		}

		pos = [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
	}

	return strconv.Itoa(res), nil
}

func makeBoard(input *helper.InputReader) [][]byte {
	var res [][]byte
	for line := range input.IterateLines {
		res = append(res, []byte(line))
	}
	return res
}

func findGuard(in [][]byte) ([2]int, int) {
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[0]); j++ {
			switch in[i][j] {
			case '^':
				return [2]int{i, j}, 0
			case '>':
				return [2]int{i, j}, 1
			case 'v':
				return [2]int{i, j}, 2
			case '<':
				return [2]int{i, j}, 3
			}

		}
	}
	return [2]int{}, 0
}

func inside(board [][]byte, pos [2]int) bool {
	return pos[0] >= 0 && pos[0] < len(board) && pos[1] >= 0 && pos[1] < len(board[0])
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var board = makeBoard(input)

	for i := range board {
		for j := range board[i] {
			if board[i][j] != '#' && board[i][j] != '^' {
				board[i][j] = '#'
				if FindLoop(board) {
					res++
				}
				board[i][j] = '.'
			}
		}
	}

	return strconv.Itoa(res), nil
}

func FindLoop(board [][]byte) bool {
	var dirBoard = make([][][4]bool, len(board))
	for i := range dirBoard {
		dirBoard[i] = make([][4]bool, len(board[0]))
	}
	var pos, dir = findGuard(board)

	for inside(board, pos) {
		if dirBoard[pos[0]][pos[1]][dir] {
			return true
		}
		dirBoard[pos[0]][pos[1]][dir] = true

		npos := [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		for inside(board, npos) && board[npos[0]][npos[1]] == '#' {
			dir = (dir + 1) % 4
			npos = [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		}

		pos = [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
	}

	return false
}
