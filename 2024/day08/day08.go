package day08

import (
	"adventofcode/helper"
	"log"
	"strconv"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var lines, _ = input.ReadLines()
	var freqs = map[byte][][2]int{}
	var antinodes = make([][]byte, len(lines))

	for i := range lines {
		antinodes[i] = []byte(lines[i])
		for j := range lines[i] {
			if lines[i][j] != '.' {
				freqs[lines[i][j]] = append(freqs[lines[i][j]], [2]int{i, j})
			}
		}
	}

	for _, antenas := range freqs {
		for i := range antenas {
			for j := range antenas {
				if i == j {
					continue
				}
				anti := findAfterDouble(antenas[i], antenas[j])
				if inside(antinodes, anti) && antinodes[anti[0]][anti[1]] != '#' {
					antinodes[anti[0]][anti[1]] = '#'
					res++
				}
			}
			//PrintBoard(antinodes)
		}
	}

	return strconv.Itoa(res), nil
}

func findAfterDouble(a, b [2]int) [2]int {
	var res [2]int
	res[0] = a[0] + (a[0] - b[0])
	res[1] = a[1] + (a[1] - b[1])
	return res
}

func inside[T any](board [][]T, pos [2]int) bool {
	return pos[0] >= 0 && pos[0] < len(board) && pos[1] >= 0 && pos[1] < len(board[0])
}

func PrintBoard(in [][]byte) {
	for _, line := range in {
		log.Println(string(line))
	}
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var lines, _ = input.ReadLines()
	var freqs = map[byte][][2]int{}
	var antinodes = make([][]byte, len(lines))

	for i := range lines {
		antinodes[i] = []byte(lines[i])
		for j := range lines[i] {
			if lines[i][j] != '.' {
				freqs[lines[i][j]] = append(freqs[lines[i][j]], [2]int{i, j})
				antinodes[i][j] = '#'
				res++
			}
		}
	}

	for _, antenas := range freqs {
		for i := range antenas {
			for j := range antenas {
				if i == j {
					continue
				}
				for _, anti := range findInLine(antinodes, antenas[i], antenas[j]) {
					if inside(antinodes, anti) && antinodes[anti[0]][anti[1]] != '#' {
						antinodes[anti[0]][anti[1]] = '#'
						res++
					}
				}
			}
			//PrintBoard(antinodes)
		}
	}

	return strconv.Itoa(res), nil
}

func findInLine(board [][]byte, a, b [2]int) [][2]int {
	var res [][2]int
	var dx, dy = (a[0] - b[0]), (a[1] - b[1])
	var cur = [2]int{a[0] + dx, a[1] + dy}
	for inside(board, cur) {
		res = append(res, cur)
		cur[0] += dx
		cur[1] += dy
	}
	return res
}
