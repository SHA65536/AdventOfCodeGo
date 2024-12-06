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
		// Mark unique visited spots and add to result
		if board[pos[0]][pos[1]] != 'X' {
			board[pos[0]][pos[1]] = 'X'
			res++
		}

		// Next position
		npos := [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		// If next position is blocked
		for inside(board, npos) && board[npos[0]][npos[1]] == '#' {
			// Rotate until you are not blocked
			dir = (dir + 1) % 4
			npos = [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		}

		// Move to the next position
		pos = npos
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

// Find guard position and direction
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

func inside[T any](board [][]T, pos [2]int) bool {
	return pos[0] >= 0 && pos[0] < len(board) && pos[1] >= 0 && pos[1] < len(board[0])
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var board = makeBoard(input)
	var initialPos, initialDir = findGuard(board)
	var pos, dir = initialPos, initialDir

	// Mark all spots where we an obstacle will change the path
	for inside(board, pos) {
		board[pos[0]][pos[1]] = 'X'

		npos := [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		for inside(board, npos) && board[npos[0]][npos[1]] == '#' {
			dir = (dir + 1) % 4
			npos = [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		}

		pos = npos
	}

	// Check all possible places to place an obstacle
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' && !(i == initialPos[0] && j == initialPos[1]) {
				board[i][j] = '#'
				if FindLoop(board, initialPos, initialDir) {
					res++
				}
				board[i][j] = 'X'
			}
		}
	}

	return strconv.Itoa(res), nil
}

func FindLoop(board [][]byte, pos [2]int, dir int) bool {
	// dirboard[x][y] stores which directions we were when we walked in this spot
	var dirBoard = make([][][4]bool, len(board))
	for i := range dirBoard {
		dirBoard[i] = make([][4]bool, len(board[0]))
	}

	// Walk until you find a loop
	for inside(board, pos) {
		// If we were in the same position in the same direction, loop
		if dirBoard[pos[0]][pos[1]][dir] {
			return true
		}
		// Record current direction
		dirBoard[pos[0]][pos[1]][dir] = true

		npos := [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		for inside(board, npos) && board[npos[0]][npos[1]] == '#' {
			dir = (dir + 1) % 4
			npos = [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		}

		pos = npos
	}

	return false
}

type Positon struct {
	B byte
	D [4]bool
}

func Star2Clone(input *helper.InputReader) (string, error) {
	var res int
	var board = makeBoardDirs(input)
	var pos, dir = findGuardDirs(board)

	for inside(board, pos) {
		board[pos[0]][pos[1]].B = 'X'
		npos := [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}

		// Attempt to place an obstacle
		if inside(board, npos) && board[npos[0]][npos[1]].B != 'X' {
			var copyBoard = copyBoardState(board)
			copyBoard[npos[0]][npos[1]].B = '#'
			if findLoopState(copyBoard, pos, dir) {
				res++
			}
		}

		board[pos[0]][pos[1]].D[dir] = true

		for inside(board, npos) && board[npos[0]][npos[1]].B == '#' {
			dir = (dir + 1) % 4
			npos = [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}

			// Attempt to place an obstacle after a spin
			if inside(board, npos) && board[npos[0]][npos[1]].B != 'X' {
				var copyBoard = copyBoardState(board)

				copyBoard[npos[0]][npos[1]].B = '#'

				if findLoopState(copyBoard, pos, dir) {
					res++
				}
			}
		}

		pos = npos
	}

	return strconv.Itoa(res), nil
}

func copyBoardState(board [][]Positon) [][]Positon {
	var copyBoard = make([][]Positon, len(board))
	for i := range board {
		copyBoard[i] = make([]Positon, len(board[i]))
		copy(copyBoard[i], board[i])
	}
	return copyBoard
}

func makeBoardDirs(input *helper.InputReader) [][]Positon {
	var board [][]byte
	for line := range input.IterateLines {
		board = append(board, []byte(line))
	}
	var res = make([][]Positon, len(board))
	for i := range board {
		res[i] = make([]Positon, len(board[i]))
		for j := range board[i] {
			res[i][j].B = board[i][j]
		}
	}
	return res
}

func findLoopState(board [][]Positon, pos [2]int, dir int) bool {
	// Walk until you find a loop
	for inside(board, pos) {
		// If we were in the same position in the same direction, loop
		if board[pos[0]][pos[1]].D[dir] {
			return true
		}
		// Record current direction
		board[pos[0]][pos[1]].D[dir] = true

		npos := [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		for inside(board, npos) && board[npos[0]][npos[1]].B == '#' {
			dir = (dir + 1) % 4
			npos = [2]int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
		}

		pos = npos
	}

	return false
}

func findGuardDirs(in [][]Positon) ([2]int, int) {
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[0]); j++ {
			switch in[i][j].B {
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
