package day04

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/4

func Bingo(bingo string) (string, error) {
	var boards [][5][5]int
	var drawn []int
	var marked = map[int]bool{}
	var cur, winNum, res int
	split := strings.Split(bingo, "\n\n")
	boards = make([][5][5]int, len(split)-1)
	// Loading drawn numbers
	for i := range split[0] {
		if split[0][i] >= '0' && split[0][i] <= '9' {
			cur = cur*10 + int(split[0][i]-'0')
		} else {
			drawn = append(drawn, cur)
			cur = 0
		}
	}
	// Loading boards
	for b, board := range split[1:] {
		lines := strings.Split(board, "\n")
		for i, line := range lines {
			for n, num := range strings.Fields(line) {
				boards[b][i][n], _ = strconv.Atoi(num)
			}
		}
	}
	// Finding the first board to win
findWin:
	for i := range drawn {
		// We mark the numbers from first to last
		marked[drawn[i]] = true
		for b := range boards {
			for r := range boards[b] {
				var numRow, numCol int
				for c := range boards[b][r] {
					if marked[boards[b][r][c]] {
						numRow++
					}
					if marked[boards[b][c][r]] {
						numCol++
					}
				}
				// If a row or col is all marked, we found the board
				if numRow == 5 || numCol == 5 {
					cur = b
					winNum = drawn[i]
					break findWin
				}
			}
		}
	}
	// Summing up the score
	for r := range boards[cur] {
		for c := range boards[cur][r] {
			if !marked[boards[cur][r][c]] {
				res += boards[cur][r][c]
			}
		}
	}
	return strconv.Itoa(res * winNum), nil
}

func BingoLast(bingo string) (string, error) {
	var boards [][5][5]int
	var drawn []int
	var marked = map[int]bool{}
	var cur, winNum, res int
	split := strings.Split(bingo, "\n\n")
	boards = make([][5][5]int, len(split)-1)
	// Loading drawn numbers
	for i := range split[0] {
		if split[0][i] >= '0' && split[0][i] <= '9' {
			cur = cur*10 + int(split[0][i]-'0')
		} else {
			// Marking all drawn numbers
			drawn = append(drawn, cur)
			marked[cur] = true
			cur = 0
		}
	}
	// Loading boards
	for b, board := range split[1:] {
		lines := strings.Split(board, "\n")
		for i, line := range lines {
			for n, num := range strings.Fields(line) {
				boards[b][i][n], _ = strconv.Atoi(num)
			}
		}
	}
	// We find the last losing board by having all numbers marked
	// and starting to unmark numbers until we found a board that isn't winning
findLosing:
	for i := len(drawn) - 1; i >= 0; i-- {
		// Unmarking numbers from last to first
		marked[drawn[i]] = false
	checkWin:
		for b := range boards {
			// Checking each board
			for r := range boards[b] {
				var numRow, numCol int
				for c := range boards[b][r] {
					if marked[boards[b][r][c]] {
						numRow++
					}
					if marked[boards[b][c][r]] {
						numCol++
					}
				}
				// If the board is winning, continue
				if numRow == 5 || numCol == 5 {
					continue checkWin
				}
			}
			// If the board is not winning, this is the last board that doesn't win
			cur = b
			winNum = drawn[i]
			break findLosing
		}
	}
	// Summing the score
	for r := range boards[cur] {
		for c := range boards[cur][r] {
			if !marked[boards[cur][r][c]] {
				res += boards[cur][r][c]
			}
		}
	}
	// Need to subtract last number since it's actually marked
	return strconv.Itoa((res - winNum) * winNum), nil
}
