package day13

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/13

func FoldInstructions(coords string) (string, error) {
	var idx int
	var points = map[[2]int]bool{}

	// Loading points
	for idx = 0; idx < len(coords); idx++ {
		var cur [2]int
		if coords[idx] == '\n' {
			break
		}
		for ; coords[idx] != ','; idx++ {
			cur[0] = cur[0]*10 + int(coords[idx]-'0')
		}
		idx++
		for ; coords[idx] != '\n'; idx++ {
			cur[1] = cur[1]*10 + int(coords[idx]-'0')
		}
		points[cur] = true
	}
	idx++
	// Looping over folds
	for ; idx < len(coords); idx++ {
		var newPoints = map[[2]int]bool{}
		var dir byte
		var val int
		idx += 11
		dir = coords[idx]
		idx += 2
		for val = 0; idx < len(coords) && coords[idx] != '\n'; idx++ {
			val = val*10 + int(coords[idx]-'0')
		}
		if dir == 'y' {
			// Horizontal fold
			for k := range points {
				// Only below the fold
				if k[1] > val {
					newPoints[[2]int{k[0], val - (k[1] - val)}] = true
				} else {
					newPoints[k] = true
				}
			}
		} else {
			// Vertical fold
			for k := range points {
				// Only right of the fold
				if k[0] > val {
					newPoints[[2]int{val - (k[0] - val), k[1]}] = true
				} else {
					newPoints[k] = true
				}
			}
		}
		points = newPoints
		break
	}
	return strconv.Itoa(len(points)), nil
}

func FoldAll(coords string) (string, error) {
	var idx int
	var points = map[[2]int]bool{}
	var board [][]byte
	var res strings.Builder

	// Loading points
	for idx = 0; idx < len(coords); idx++ {
		var cur [2]int
		if coords[idx] == '\n' {
			break
		}
		for ; coords[idx] != ','; idx++ {
			cur[0] = cur[0]*10 + int(coords[idx]-'0')
		}
		idx++
		for ; coords[idx] != '\n'; idx++ {
			cur[1] = cur[1]*10 + int(coords[idx]-'0')
		}
		points[cur] = true
	}
	idx++
	// Looping over folds
	for ; idx < len(coords); idx++ {
		var newPoints = map[[2]int]bool{}
		var dir byte
		var val int
		idx += 11
		dir = coords[idx]
		idx += 2
		for val = 0; idx < len(coords) && coords[idx] != '\n'; idx++ {
			val = val*10 + int(coords[idx]-'0')
		}
		if dir == 'y' {
			// Horizontal fold
			for k := range points {
				// Only below the fold
				if k[1] > val {
					newPoints[[2]int{k[0], val - (k[1] - val)}] = true
				} else {
					newPoints[k] = true
				}
			}
		} else {
			// Vertical fold
			for k := range points {
				// Only right of the fold
				if k[0] > val {
					newPoints[[2]int{val - (k[0] - val), k[1]}] = true
				} else {
					newPoints[k] = true
				}
			}
		}
		points = newPoints
	}
	// Making board
	var maxX, maxY int
	for k := range points {
		if k[0] > maxX {
			maxX = k[0]
		}
		if k[1] > maxY {
			maxY = k[1]
		}
	}
	board = make([][]byte, maxY+1)
	for i := range board {
		board[i] = make([]byte, maxX+1)
	}
	// Writing to board
	for k := range points {
		board[k[1]][k[0]] = '#'
	}
	// Writing to answer
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				res.WriteByte(' ')
			} else {
				res.WriteByte('#')
			}
		}
		res.WriteByte('\n')
	}
	return res.String(), nil
}
