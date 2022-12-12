package day12

import (
	"strconv"
	"strings"
)

var dirs = [][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

// https://adventofcode.com/2022/day/12
func LeastSteps(input string) (string, error) {
	var seen = map[[2]int]bool{}
	start, end, board := LoadBoard(input)

	// Queue for positions we want to evaluate
	queue := make(chan Coord, len(board)*len(board[0]))

	// Starting position
	queue <- Coord{start, 0}

	// While we have positions to evaluate
	for len(queue) > 0 {
		cur := <-queue
		// Check if we reached the end
		if cur.c == end {
			return strconv.Itoa(cur.v), nil
		}
		// Check if we already evaluated this earlier
		if seen[cur.c] {
			continue
		}
		seen[cur.c] = true
		// Check all directions
		for _, dir := range dirs {
			nx, ny := cur.c[0]+dir[0], cur.c[1]+dir[1]
			// Make sure we're inside the board
			if nx < 0 || ny < 0 || nx >= len(board) || ny >= len(board[0]) {
				continue
			}
			// Make sure we are making legal moves
			if !canGo(board[cur.c[0]][cur.c[1]], board[nx][ny]) {
				continue
			}
			// Add to evaluation list
			queue <- Coord{[2]int{nx, ny}, cur.v + 1}
		}
	}

	return "-1", nil
}

func LeastStepsAny(input string) (string, error) {
	var seen = map[[2]int]bool{}
	// Here we want to go backwards, from the end point
	// towards the first point with elevation 0
	_, start, board := LoadBoard(input)

	// Queue for positions we want to evaluate
	queue := make(chan Coord, len(board)*len(board[0]))

	// Starting position
	queue <- Coord{start, 0}

	// While we have positions to evaluate
	for len(queue) > 0 {
		cur := <-queue
		// Check if we reached a tile with elevation 0
		if board[cur.c[0]][cur.c[1]] == 0 {
			return strconv.Itoa(cur.v), nil
		}
		// Check if we already evaluated this earlier
		if seen[cur.c] {
			continue
		}
		seen[cur.c] = true
		// Check all directions
		for _, dir := range dirs {
			nx, ny := cur.c[0]+dir[0], cur.c[1]+dir[1]
			// Make sure we're inside the board
			if nx < 0 || ny < 0 || nx >= len(board) || ny >= len(board[0]) {
				continue
			}
			// Make sure we are making legal moves (note: this is reversed from part1)
			if !canGo(board[nx][ny], board[cur.c[0]][cur.c[1]]) {
				continue
			}
			// Add to evaluation list
			queue <- Coord{[2]int{nx, ny}, cur.v + 1}
		}
	}

	return "-1", nil
}

func LoadBoard(input string) (start, end [2]int, board [][]int) {
	// Loading board and finding start and end
	for i, line := range strings.Split(input, "\n") {
		var cur = make([]int, len(line))
		for j := range line {
			if line[j] == 'S' { // Start is 0 elevation
				start = [2]int{i, j}
				cur[j] = 0
			} else if line[j] == 'E' { // End is 25 elevation
				end = [2]int{i, j}
				cur[j] = 25
			} else {
				cur[j] = int(line[j] - 'a')
			}
		}
		board = append(board, cur)
	}
	return start, end, board
}

// Can go from a to b
func canGo(a, b int) bool {
	if a > b {
		return true
	}
	return b-a <= 1
}

type Coord struct {
	c [2]int
	v int
}
