package day12

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
)

type Shape [4][3][3]byte

func Star1(input *helper.InputReader) (string, error) {
	var res int
	var line string
	var err error

	var shapes []Shape

	for line, err = input.ReadLine(); err == nil; line, err = input.ReadLine() {
		if line[1] != ':' {
			break
		}
		var cur Shape
		line, _ = input.ReadLine()
		cur[0][0] = [3]byte{line[0], line[1], line[2]}
		line, _ = input.ReadLine()
		cur[0][1] = [3]byte{line[0], line[1], line[2]}
		line, _ = input.ReadLine()
		cur[0][2] = [3]byte{line[0], line[1], line[2]}

		cur[1] = Clockwise(cur[0])
		cur[2] = Clockwise(cur[1])
		cur[3] = Clockwise(cur[2])
		shapes = append(shapes, cur)
		input.ReadLine()
	}

	for ; err == nil; line, err = input.ReadLine() {
		var board [][]byte
		var needed = make([]int, 6)
		var width, height int
		fmt.Sscanf(line, "%dx%d: %d %d %d %d %d %d", &width, &height, &(needed[0]), &(needed[1]), &(needed[2]), &(needed[3]), &(needed[4]), &(needed[5]))
		board = make([][]byte, width)
		for i := range board {
			board[i] = make([]byte, height)
		}
		if CanSolveBoard(board, shapes, needed) {
			res++
		}
	}

	return strconv.Itoa(res), nil
}

func CanSolveBoard(board [][]byte, shapes []Shape, needed []int) bool {
	var sumSize int
	for i := range needed {
		sumSize += SizeOfShape(shapes[i][0]) * needed[i]
	}
	if sumSize > len(board)*len(board[0]) {
		return false
	}

	var solveHelper func() bool
	solveHelper = func() bool {
		var shapeIdx int
		for ; shapeIdx < len(needed); shapeIdx++ {
			if needed[shapeIdx] > 0 {
				break
			}
		}
		if shapeIdx == len(needed) {
			return true
		}
		needed[shapeIdx]--
		for i := 0; i < len(board)-2; i++ {
			for j := 0; j < len(board[0])-2; j++ {
				for rot := range 4 {
					if CanPlace(board, shapes[shapeIdx][rot], i, j) {
						Place(board, shapes[shapeIdx][rot], i, j)
						if solveHelper() {
							return true
						}
						Unplace(board, shapes[shapeIdx][rot], i, j)
					}
				}
			}
		}
		needed[shapeIdx]++
		return false
	}
	return solveHelper()
}

func CanPlace(board [][]byte, shape [3][3]byte, i, j int) bool {
	if i+2 >= len(board) || j+2 >= len(board[0]) {
		return false
	}

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[i+r][j+c] == 1 && shape[r][c] == '#' {
				return false
			}
		}
	}

	return true
}

func Place(board [][]byte, shape [3][3]byte, i, j int) {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if shape[r][c] == '#' {
				board[i+r][j+c] = 1
			}
		}
	}
}

func Unplace(board [][]byte, shape [3][3]byte, i, j int) {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if shape[r][c] == '#' {
				board[i+r][j+c] = 0
			}
		}
	}
}

func Clockwise(in [3][3]byte) [3][3]byte {
	var out [3][3]byte
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			out[c][2-r] = in[r][c]
		}
	}
	return out
}

func SizeOfShape(in [3][3]byte) int {
	var res int
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if in[r][c] == '#' {
				res++
			}
		}
	}
	return res
}
