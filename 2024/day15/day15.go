package day15

import (
	"adventofcode/helper"
	"strconv"
)

var up = [2]int{-1, 0}
var down = [2]int{1, 0}
var left = [2]int{0, -1}
var right = [2]int{0, 1}

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var board, moves = parseBoardInputs(input)
	var bot = findBot(board)

	for _, move := range moves {
		makeMove(board, move, &bot)
	}

	res = calScore(board)

	return strconv.Itoa(res), nil
}

func calScore(board [][]byte) int {
	var res int
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' || board[i][j] == '[' {
				res += (i * 100) + j
			}
		}
	}
	return res
}

func makeMove(board [][]byte, move [2]int, bot *[2]int) bool {
	var newPos = addDiff(*bot, move)
	if !helper.InsideP(board, newPos) || board[newPos[0]][newPos[1]] == '#' {
		return false
	}

	if board[newPos[0]][newPos[1]] == '.' {
		board[bot[0]][bot[1]] = '.'
		board[newPos[0]][newPos[1]] = '@'
		*bot = newPos
		return true
	}

	var curPos = newPos
	for helper.InsideP(board, curPos) && board[curPos[0]][curPos[1]] != '#' {
		if board[curPos[0]][curPos[1]] == '.' {
			board[bot[0]][bot[1]] = '.'
			board[newPos[0]][newPos[1]] = '@'
			board[curPos[0]][curPos[1]] = 'O'
			*bot = newPos
			return true
		}
		curPos = addDiff(curPos, move)
	}

	return false
}

func addDiff(a, b [2]int) [2]int {
	return [2]int{a[0] + b[0], a[1] + b[1]}
}

func parseBoardInputs(input *helper.InputReader) (board [][]byte, moves [][2]int) {
	board = make([][]byte, 0)
	moves = make([][2]int, 0)

	var doneMap bool
	for line := range input.IterateLines {
		if line == "" {
			doneMap = true
			continue
		}
		if !doneMap {
			board = append(board, []byte(line))
		} else {
			for _, char := range line {
				switch char {
				case '^':
					moves = append(moves, up)
				case '<':
					moves = append(moves, left)
				case '>':
					moves = append(moves, right)
				case 'v':
					moves = append(moves, down)
				}
			}
		}
	}

	return board, moves
}

func findBot(in [][]byte) [2]int {
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[0]); j++ {
			if in[i][j] == '@' {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var board, moves = parseBoardInputsTwo(input)
	var bot = findBot(board)

	for _, move := range moves {
		makeMoveTwo(board, move, &bot)
	}

	res = calScore(board)

	return strconv.Itoa(res), nil
}

func makeMoveTwo(board [][]byte, move [2]int, bot *[2]int) bool {
	var newPos = addDiff(*bot, move)
	if !helper.InsideP(board, newPos) || board[newPos[0]][newPos[1]] == '#' {
		return false
	}

	if board[newPos[0]][newPos[1]] == '.' {
		board[bot[0]][bot[1]] = '.'
		board[newPos[0]][newPos[1]] = '@'
		*bot = newPos
		return true
	}

	if move == left || move == right {
		var found bool
		var curPos = newPos
		for helper.InsideP(board, curPos) && board[curPos[0]][curPos[1]] != '#' {
			if board[curPos[0]][curPos[1]] == '.' {
				found = true
				break
			}
			curPos = addDiff(curPos, move)
		}
		if !found {
			return false
		}
		curPos = newPos
		var curMark = board[bot[0]][bot[1]]
		for board[curPos[0]][curPos[1]] != '.' {
			board[curPos[0]][curPos[1]], curMark = curMark, board[curPos[0]][curPos[1]]
			curPos = addDiff(curPos, move)
		}
		board[curPos[0]][curPos[1]] = curMark
		board[bot[0]][bot[1]] = '.'
		*bot = newPos
		return true
	}

	if canMoveVertical(board, move, newPos) {
		var extra = addDiff(newPos, left)
		if board[newPos[0]][newPos[1]] == '[' {
			extra = addDiff(newPos, right)
		}

		pushVertical(board, move, newPos)
		board[bot[0]][bot[1]] = '.'
		board[extra[0]][extra[1]] = '.'
		board[newPos[0]][newPos[1]] = '@'
		*bot = newPos
		return true
	}

	return false
}

func canMoveVertical(board [][]byte, move [2]int, pos [2]int) bool {
	var extra = addDiff(pos, left)
	if board[pos[0]][pos[1]] == '[' {
		extra = addDiff(pos, right)
	}

	var nPos, nExtra = addDiff(pos, move), addDiff(extra, move)
	if board[nPos[0]][nPos[1]] == '.' && board[nExtra[0]][nExtra[1]] == '.' {
		return true
	}

	if board[nPos[0]][nPos[1]] == '#' || board[nExtra[0]][nExtra[1]] == '#' {
		return false
	}

	if board[nPos[0]][nPos[1]] != '.' {
		if !canMoveVertical(board, move, nPos) {
			return false
		}
	}

	if board[nExtra[0]][nExtra[1]] != '.' {
		if !canMoveVertical(board, move, nExtra) {
			return false
		}
	}

	return true
}

func pushVertical(board [][]byte, move [2]int, pos [2]int) {
	var extra = addDiff(pos, left)
	if board[pos[0]][pos[1]] == '[' {
		extra = addDiff(pos, right)
	}

	var nPos, nExtra = addDiff(pos, move), addDiff(extra, move)
	if board[nPos[0]][nPos[1]] != '.' {
		pushVertical(board, move, nPos)
	}

	if board[nExtra[0]][nExtra[1]] != '.' {
		pushVertical(board, move, nExtra)
	}
	board[nPos[0]][nPos[1]] = board[pos[0]][pos[1]]
	board[nExtra[0]][nExtra[1]] = board[extra[0]][extra[1]]
	board[pos[0]][pos[1]] = '.'
	board[extra[0]][extra[1]] = '.'
}

func parseBoardInputsTwo(input *helper.InputReader) (board [][]byte, moves [][2]int) {
	board = make([][]byte, 0)
	moves = make([][2]int, 0)

	var doneMap bool
	for line := range input.IterateLines {
		if line == "" {
			doneMap = true
			continue
		}
		if !doneMap {
			var row = make([]byte, 0, len(line)*2)
			for _, char := range line {
				if char == 'O' {
					row = append(row, '[', ']')
				} else if char == '@' {
					row = append(row, '@', '.')
				} else {
					row = append(row, byte(char), byte(char))
				}
			}
			board = append(board, row)
		} else {
			for _, char := range line {
				switch char {
				case '^':
					moves = append(moves, up)
				case '<':
					moves = append(moves, left)
				case '>':
					moves = append(moves, right)
				case 'v':
					moves = append(moves, down)
				}
			}
		}
	}

	return board, moves
}
