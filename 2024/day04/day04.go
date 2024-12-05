package day04

import (
	"adventofcode/helper"
	"strconv"
)

var find = "XMAS"
var dirs = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func Star1(input *helper.InputReader) (string, error) {
	var res int
	var chars, _ = input.ReadLines()

	for r := range chars {
		for c := range chars[r] {
			if chars[r][c] != 'X' {
				continue
			}
			var nr, nc int
		eachdir:
			for _, dir := range dirs {
				nr, nc = r, c
				for ch := range find {
					if nr < 0 || nr >= len(chars) || nc < 0 || nc >= len(chars[0]) {
						continue eachdir
					}
					if chars[nr][nc] != find[ch] {
						continue eachdir
					}
					nr, nc = nr+dir[0], nc+dir[1]
				}
				res++
			}
		}
	}

	return strconv.Itoa(res), nil
}

var dirs2 = [4][2][2]int{
	{{-1, -1}, {1, 1}},
	{{1, 1}, {-1, -1}},
	{{-1, 1}, {1, -1}},
	{{1, -1}, {-1, 1}},
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var chars, _ = input.ReadLines()

	for r := 1; r < len(chars)-1; r++ {
		for c := 1; c < len(chars[0])-1; c++ {
			var fnd bool
			if chars[r][c] != 'A' {
				continue
			}
			for _, dir := range dirs2 {
				if chars[r+dir[0][0]][c+dir[0][1]] == 'M' &&
					chars[r+dir[1][0]][c+dir[1][1]] == 'S' {
					if fnd {
						res++
						break
					}
					fnd = true
				}
			}
		}
	}

	return strconv.Itoa(res), nil
}
